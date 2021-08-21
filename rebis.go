package rebis

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type Item struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	*cache
}

type cache struct {
	mu                sync.RWMutex
	items             map[string]Item
	defaultExpiration time.Duration
	janitor           *janitor
	backup            *backup
	logger            Logger
	onEvicted         func(string, interface{})
}

type keyAndValue struct {
	key   string
	value interface{}
}

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

/*
	Create new rebis cache from config struct.
*/
func NewCache(config *Config) (*Cache, error) {
	return newCache(config, make(map[string]Item))
}

/*
	Init cache struct with default logger (stdout).

	Run backup if inUse = true in config file with backup interval

	Run janitor, if CleanupInterval <= 0 janitor does not start.
*/
func newCache(config *Config, items map[string]Item) (*Cache, error) {
	c := &cache{
		defaultExpiration: config.DefaultExpiration,
		items:             items,
		logger:            DefaultLogger(),
	}
	C := &Cache{c}
	c.logger.Printf("initialize new cache with defaul expiration duration: %s and items count: %d",
		c.defaultExpiration,
		len(c.items),
	)

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
	Delete all expired items from the cache.
*/
func (c *cache) DeleteExpired() {
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
	Delete item by key, return his value and if have evicted function
	then true else false.
*/
func (c *cache) delete(k string) (interface{}, bool) {
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
	Saving backup to the path defined in the structure.
*/
func (c *cache) BackupSave() error {
	return c.BackupSaveFile(c.backup.Path)
}

/*
	Saving backup by filename path.
*/
func (c *cache) BackupSaveFile(filename string) error {
	buf, err := ffjson.Marshal(&c.items)
	if err != nil {
		c.logger.Printf(err.Error())
		return err
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		c.logger.Printf(err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(buf)
	if err != nil {
		c.logger.Printf(err.Error())
		return err
	}

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
		c.logger.Printf(err.Error())
		return err
	}

	items := map[string]Item{}
	err = ffjson.Unmarshal(buf, &items)
	if err != nil {
		c.logger.Printf(err.Error())
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range items {
		ov, found := c.items[k]
		if !found || ov.Expired() {
			c.items[k] = v
		}
	}
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
	return m
}

/*
	Add an item to the cache, replacing any existing item. If the duration is 0
	(DefaultExpiration), the cache's default expiration time is used. If it is -1
	(NoExpiration), the item never expires.
*/
func (c *cache) Set(k string, x interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[k] = Item{
		Value:      x,
		Expiration: e,
	}
	c.mu.Unlock()
	c.logger.Printf("set %s -> %v -> %s", k, x, d)
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
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("item %s already exists", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	c.logger.Printf("add %s -> %v -> %s", k, x, d)
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
	c.logger.Printf("get %s -> %v", k, item.Value)
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
	c.logger.Printf("get with %s -> %v", k, item.Value)
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
	c.logger.Printf("replace with %s -> %v -> %s", k, x, d)
	return nil
}
