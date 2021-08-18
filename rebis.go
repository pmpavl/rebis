package rebis

import (
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
	logger            *logger
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

func NewCache(config *Config) (*Cache, error) {
	return newCache(config, make(map[string]Item))
}

func newCache(config *Config, items map[string]Item) (*Cache, error) {
	c := &cache{
		defaultExpiration: config.DefaultExpiration,
		items:             items,
	}
	C := &Cache{c}
	if ci := config.CleanupInterval; ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}

	if config.LoggerPath == "-1" {
		startLogger(DefaultLoggerPath, config.LoggerLevel)
	} else {
		startLogger(config.LoggerPath, config.LoggerLevel)
	}
	return C, nil
}

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
}

func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, false
		}
	}
	c.mu.RUnlock()
	return item.Value, true
}
