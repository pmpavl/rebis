package rebis

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/pquerna/ffjson/ffjson"
)

// Item is element of cache.
type Item struct {
	Value      interface{}
	Expiration int64
}

// Cache is wrapper over hidden cache
type Cache struct {
	*cache
}

type cache struct {
	maxSize           uintptr
	size              uintptr
	mu                sync.RWMutex
	items             map[string]Item
	defaultExpiration time.Duration
	janitor           *janitor
	backup            *backup
	logger            Logger
	logAll            bool
	onEvicted         func(string, interface{})
}

type keyAndValue struct {
	key   string
	value interface{}
}

const (
	NoExpiration      time.Duration = -1 // if expiration in Item = -1, then element never delete
	DefaultExpiration time.Duration = 0  // if expiration in Set func = 0, then element expiration = config.defaultExpiration
	sizeItem          uintptr       = unsafe.Sizeof(Item{})
)

/*
	NewCache create new rebis cache from config struct.
*/
func NewCache(config *Config) (*Cache, error) {
	return newCache(config, make(map[string]Item))
}

/*
	NewCacheFrom create new rebis cache from config struct and reuse map items.
*/
func NewCacheFrom(config *Config, items map[string]Item) (*Cache, error) {
	return newCache(config, items)
}

/*
	Init cache struct with default logger (stdout).

	Run backup if inUse = true in config file with backup interval

	Run janitor, if CleanupInterval <= 0 janitor does not start.
*/
func newCache(config *Config, items map[string]Item) (*Cache, error) {
	c := &cache{
		maxSize:           config.Size * 1024,
		size:              0,
		defaultExpiration: config.DefaultExpiration,
		items:             items,
		logger:            DefaultLogger(),
		logAll:            config.LogAll,
	}
	C := &Cache{c}
	c.logIf("initialize new cache with defaul expiration duration: %s, items count: %d, max count: %d",
		c.defaultExpiration,
		len(c.items),
		c.maxSize/sizeItem,
	)

	if config.Evicted {
		c.onEvicted = func(s string, i interface{}) {
			c.logIf("delete evicted: %s -> %v", s, i)
		}
	}

	if config.Backup.InUse {
		runBackup(c, config.Backup.Path, config.Backup.Interval)
	}

	if ci := config.CleanupInterval; ci > 0 {
		runJanitor(c, ci)
	}

	runtime.SetFinalizer(C, nil)
	return C, nil
}

/*
	Custom function onEvicted.
*/
func (c *cache) OnEvicted(f func(string, interface{})) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

/*
	Delete all expired items from the cache.
*/
func (c *cache) DeleteExpired() {
	c.logIf("delete expired")
	var evictedItems []keyAndValue
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			ov, evicted := c.delete(k)
			if evicted {
				evictedItems = append(evictedItems, keyAndValue{k, ov})
			}
		}
	}
	c.mu.Unlock()

	for _, v := range evictedItems {
		c.onEvicted(v.key, v.value)
	}
}

/*
	Delete an item from the cache. Does nothing if the key is not in the cache.
*/
func (c *cache) Delete(k string) {
	c.mu.Lock()
	v, evicted := c.delete(k)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(k, v)
	}
	c.logIf("delete %s -> %v", k, v)
}

/*
	Delete item by key, return his value and if have evicted function
	then true else false.
*/
func (c *cache) delete(k string) (interface{}, bool) {
	defer func() {
		c.size -= sizeItem
	}()
	if c.onEvicted != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v.Value, true
		}
	}
	delete(c.items, k)
	return nil, false
}

/*
	Delete all items from the cache.
*/
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = map[string]Item{}
	c.mu.Unlock()
}

/*
	Saving backup to the path defined in the structure.
*/
func (c *cache) BackupSave() error {
	return c.BackupSaveFile(c.backup.Path)
}

/*
	Saving backup by filename path.
*/
func (c *cache) BackupSaveFile(filename string) error {
	c.mu.Lock()
	buf, err := ffjson.Marshal(&c.items)
	c.mu.Unlock()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(buf)
	if err != nil {
		return err
	}
	c.logIf("backup save in file: %s", filename)
	ffjson.Pool(buf)
	return nil
}

/*
	Recovery backup by path defined in the structure.
*/
func (c *cache) BackupRecovery() error {
	return c.BackupRecoveryFile(c.backup.Path)
}

/*
	Recovery backup by filename path.
*/
func (c *cache) BackupRecoveryFile(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	items := map[string]Item{}
	err = ffjson.Unmarshal(buf, &items)
	ffjson.Pool(buf)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range items {
		ov, found := c.items[k]
		if !found || ov.Expired() {
			if !c.haveSlot() {
				return fmt.Errorf("no empty slot, for next items")
			}
			c.items[k] = v
			c.size += sizeItem
		}
	}
	c.logIf("backup load from file: %s", filename)
	return nil
}

/*
	Returns the number of items in the cache. This may include items that have
	expired, but have not yet been cleaned up.
*/
func (c *cache) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	c.logIf("all items count: %d", n)
	return n
}

/*
	Copies all unexpired items in the cache into a new map and returns it.
*/
func (c *cache) Items() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]Item, len(c.items))
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if v.Expiration > 0 {
			if now > v.Expiration {
				continue
			}
		}
		m[k] = v
	}
	c.logIf("unexpired items count: %d", len(m))
	return m
}

func (c *cache) haveSlot() bool {
	return c.maxSize-c.size > sizeItem
}

/*
	Add an item to the cache, replacing any existing item. If the duration is 0
	(DefaultExpiration), the cache's default expiration time is used. If it is -1
	(NoExpiration), the item never expires.
*/
func (c *cache) Set(k string, x interface{}, d time.Duration) error {
	if !c.haveSlot() {
		return fmt.Errorf("no empty slot, wait for janitor")
	}

	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	_, ok := c.get(k)
	if !ok {
		c.size += sizeItem
	}
	c.items[k] = Item{
		Value:      x,
		Expiration: e,
	}
	c.mu.Unlock()

	c.logIf("set %s -> %v <- %s", k, x, d)
	return nil
}

/*
	Add an item to the cache, replacing any existing item, using the default expiration.
*/
func (c *cache) SetDefault(k string, x interface{}) {
	c.Set(k, x, DefaultExpiration)
}

/*
	Add an item to the cache only if an item doesn't already exist for the given
	key, or if the existing item has expired. Returns an error otherwise.
*/
func (c *cache) Add(k string, x interface{}, d time.Duration) error {
	if !c.haveSlot() {
		return fmt.Errorf("no empty slot, wait for janitor")
	}
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("item %s already exists", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	c.logIf("add %s -> %v <- %s", k, x, d)
	return nil
}

func (c *cache) get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Value, true
}

func (c *cache) set(k string, x interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Value:      x,
		Expiration: e,
	}
	c.size += sizeItem
}

/*
	Get an item from the cache. Returns the item or nil, and a bool indicating
	whether the key was found.
*/
func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	c.logIf("get %s -> %v", k, item.Value)
	return item.Value, true
}

/*
	GetWithExpiration returns an item and its expiration time from the cache.
	It returns the item or nil, the expiration time if one is set (if the item
	never expires a zero value for time.Time is returned), and a bool indicating
	whether the key was found.
*/
func (c *cache) GetWithExpiration(k string) (interface{}, time.Time, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[k]
	if !found {
		return nil, time.Time{}, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, time.Time{}, false
		}
	}
	c.logIf("get with exp %s -> %v <- %s", k, item.Value, time.Unix(0, item.Expiration))
	return item.Value, time.Unix(0, item.Expiration), true
}

/*
	Set a new value for the cache key only if it already exists, and the existing
	item hasn't expired. Returns an error otherwise.
*/
func (c *cache) Replace(k string, x interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return fmt.Errorf("item %s doesn't exist", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	c.logIf("replace %s -> %v <- %s", k, x, d)
	return nil
}
