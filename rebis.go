package rebis

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"go.uber.org/zap"
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
	logger            *zap.Logger
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
	Init cache struct.
	Run zap logger, if -1 log write in stderr.
	Run janitor, if <= 0 janitor does not start.
*/
func newCache(config *Config, items map[string]Item) (*Cache, error) {
	c := &cache{
		defaultExpiration: config.DefaultExpiration,
		items:             items,
	}
	C := &Cache{c}

	logPath := config.LoggerPath
	if logPath == "-1" {
		logPath = DefaultLoggerPath
	}
	err := runLogger(c, logPath, config.LoggerLevel)
	if err != nil {
		return nil, err
	}

	if ci := config.CleanupInterval; ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}

	c.logger.Info(
		"CREATE NEW CACHE",
		zap.String("default expiration", c.defaultExpiration.String()),
		zap.String("cleanup interval", c.janitor.Interval.String()),
		zap.String("log path", logPath),
		zap.Int8("log level", config.LoggerLevel),
	)
	return C, nil
}

/*
	Returns the number of items in the cache. This may include items that have
	expired, but have not yet been cleaned up.
*/
func (c *cache) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	c.logger.Info(
		"COUNT",
		zap.Int("items count", n),
	)
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

func (c *cache) delete(k string) (interface{}, bool) {
	if c.onEvicted != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v.Value, true
		}
	}
	delete(c.items, k)
	c.logger.Debug(
		"DELETE",
		zap.String("key", k),
	)
	return nil, false
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
	c.logger.Debug(
		"SET",
		zap.String("key", k),
		zap.String("value", fmt.Sprint(x)),
		zap.String("expiration", d.String()),
	)
}

/*
	Add an item to the cache, replacing any existing item, using the default expiration.
*/
func (c *cache) SetDefault(k string, x interface{}) {
	c.Set(k, x, DefaultExpiration)
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
		c.logger.Debug(
			"GET",
			zap.String("not found", k),
		)
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.logger.Debug(
				"GET",
				zap.String("expired", k),
			)
			return nil, false
		}
	}
	c.logger.Debug(
		"GET",
		zap.String("key", k),
		zap.String("value", fmt.Sprint(item.Value)),
	)
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
		c.logger.Debug(
			"GET",
			zap.String("not found", k),
		)
		return nil, time.Time{}, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.logger.Debug(
				"GET",
				zap.String("expired", k),
			)
			return nil, time.Time{}, false
		}
	}
	c.logger.Debug(
		"GET",
		zap.String("key", k),
		zap.String("value", fmt.Sprint(item.Value)),
		zap.Time("expiration", time.Unix(0, item.Expiration)),
	)
	return item.Value, time.Unix(0, item.Expiration), true
}
