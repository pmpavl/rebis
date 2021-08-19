package rebis

import (
	"encoding/gob"
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
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
	DefaultLoggerPath string        = "stderr"
)

/*
	Create new rebis cache from config struct.
*/
func NewCache(config *Config) (*Cache, error) {
	return newCache(config, make(map[string]Item))
}

/*
	Init cache struct with default logger (stdout).

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
	c.logger.Printf("start delete expired")
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

func (c *cache) BackupSave() {
	enc := gob.NewEncoder(c.backup.File)
	defer func() {
		if x := recover(); x != nil {

		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		fmt.Printf("%T\n", v.Value)
		gob.Register(v.Value)
	}
	err := enc.Encode(&c.items)
	if err != nil {
		c.logger.Printf(err.Error())
	}
	return
}

func (c *cache) BackupLoad(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := map[string]Item{}
	err := dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range items {
			ov, found := c.items[k]
			if !found || ov.Expired() {
				c.items[k] = v
			}
		}
	}
	fmt.Println(c)
	return err
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
		return fmt.Errorf("Item %s already exists", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
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
	return nil
}
