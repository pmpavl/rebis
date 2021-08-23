# Rebis [![build](https://github.com/pmpavl/rebis/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/pmpavl/rebis/actions/workflows/go.yaml) [![codecov](https://codecov.io/gh/pmpavl/rebis/branch/master/graph/badge.svg?token=MLE06MIFZD)](https://codecov.io/gh/pmpavl/rebis) [![Go Report Card](https://goreportcard.com/badge/github.com/pmpavl/rebis)](https://goreportcard.com/report/github.com/pmpavl/rebis) [![Go Reference](https://pkg.go.dev/badge/github.com/pmpavl/rebis.svg)](https://pkg.go.dev/github.com/pmpavl/rebis)

Key-value in-memory concurrent cache storage with: logger, backup save, auto collect expired element, limit on the maximum number of elements. All cache settings are set through the yaml config file. The elements are implemented as `map[string]Item` where Item is structure with fields Value: `interface{}` and Expiration time: `int64`.

Requires Go 1.14 or newer.

## Usage
### Installation
``` 
go get github.com/pmpavl/rebis
```
### Simple initialization
You can create a standard config file. The cache created by it will be size 1 GB, without a logger, without backups, without a onExpired function, cleanup interval 1 minute.
``` golang
import "github.com/pmpavl/rebis"

rebis.ConfigCreateDefault("defaultConfig.yaml")

rebisConfig, _ := rebis.ConfigFrom("defaultConfig.yaml")
rebisCache, _ := rebis.NewCache(rebisConfig)

rebisCache.SetDefault("my-key", "my-value")
v, _ := rebisCache.Get("my-key")
fmt.Println(v)
```
### Yaml custom initialization
You can write a config file yourself, like the one below.
``` yaml
size: 8196              # size of cache in MB
backup:
    path: "./backup"    # path to save backup 
    interval: 1m        # interval backup
    inUse: true         # do backup or not
defaultExpiration: -1ns # element standard lifetime
cleanupInterval: 1m     # cache standart interval cleanup
logAll: true            # do standart log in stdout or not
evicted: true           # do standart func to expired element or not
```
### Custom initialization
You can initialize config variable.
``` golang
import "github.com/pmpavl/rebis"

config := &rebis.Config{
		Size: 1024,
		Backup: rebis.Backup{
			InUse: false,
		},
		DefaultExpiration: time.Duration(-1),
		CleanupInterval:   time.Duration(time.Minute * 5),
		LogAll:            false,
		Evicted:           false,
	}
rebisCache, _ := rebis.NewCache(config)

rebisCache.SetDefault("my-key", "my-value")
v, _ := rebisCache.Get("my-key")
fmt.Println(v)
```
> Important, you need to remember that the backup.Path is specified as a folder, and then a file with a timestamp is created in the folder. Therefore, you need to specify exactly the path of the existing folder where will put the backups.

## Client example
In the github repository there is a folder `cmd/client` in it there is an example of concurrent writing to the cache for 20 milliseconds and concurrent reading of 2000 records.

There is also an example of a default config file `rebisDefaultConfig.yaml` and custom config file `rebisConfig.yaml` you can use it.

## Realised function
- `NewCache` `NewCacheFrom` - create an instance of rebis cache.
- `OnEvicted` - you can write a function yourself that will be applied to the evicted elements.
- `DeleteExpired` - delete all expired items from the cache.
- `Delete` - delete from cache by key.
- `Flush` - completely clears the cache.
- `BackupSave` `BackupSaveFile` `BackupRecovery` `BackupRecoveryFile` - functions responsible for saving cache backups to a default or custom path.
- `ItemCount` - count items (including evicted).
- `Items` - return map items withot evicted.
- `Set` `SetDefault` `Add` `Get` `GetWithExpiration` `Replace` - ordinary functions for accessing cache elements.
- Increment and Decrement function with all possible variations of integers and float.
- `ConfigCreateDefault` - create default config in yaml filename.
- `ConfigFrom` - create an instance of rebis cache config.

## Improvements
- Add replication support (use more than one cache instance for the cache wrapper).
- Add HTTP implementation of rebis cache.
- It seems that the implementation `map[string]Item` is not the best. We need to look at the implementation of [bigcache](https://github.com/allegro/bigcache) where is the hash function used for key and value.
- Backup is saved in json format, it's bad, because there are costs for serialization and json takes up a lot of space. Need to use a binary protocol like [protobuf](https://github.com/protocolbuffers/protobuf).
- Add it is possible to transfer logs and backups over the network.
- Add tag support for Item structure, for a faster search on them.
- Currently there is no support for deserialization of custom data types, this needs to be fixed. To do this, need to change the logic of saving backups.

## Benchmark
Three caches were compared: [rebis](https://github.com/pmpavl/rebis), [bigcache](https://github.com/allegro/bigcache), [freecache](https://github.com/coocood/freecache) and map. Benchmark tests were made using an Ryzen 7 3700X CPU @ 3.60GHz with 32GB of RAM on Windows 21H1 (19043.1165).
```
go version go1.16.4 windows/amd64

goos: windows
goarch: amd64
pkg: github.com/pmpavl/rebis
cpu: AMD Ryzen 7 3700X 8-Core Processor @ 3.60GHz

BenchmarkCacheGetExpiring-16                            60389815                79.17 ns/op           48 B/op          2 allocs/op
BenchmarkCacheGetNotExpiring-16                         78374000                68.50 ns/op           48 B/op          2 allocs/op
BenchmarkFreeCacheGet-16                                 8444374               622.6 ns/op           136 B/op          3 allocs/op
BenchmarkBigCacheGet-16                                 10803576               544.4 ns/op           152 B/op          4 allocs/op
BenchmarkRWMutexMapGet-16                               467264154               10.21 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexInterfaceMapGetStruct-16                146306853               32.83 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexInterfaceMapGetString-16                176881310               26.67 ns/op            0 B/op          0 allocs/op
BenchmarkCacheGetConcurrentExpiring-16                  105458304               45.49 ns/op           48 B/op          2 allocs/op
BenchmarkCacheGetConcurrentNotExpiring-16               100000000               52.13 ns/op           48 B/op          2 allocs/op
BenchmarkBigCacheGetConcurrent-16                       95998464               164.6 ns/op           152 B/op          4 allocs/op
BenchmarkFreeCacheGetConcurrent-16                      49889566               220.8 ns/op           136 B/op          3 allocs/op
BenchmarkRWMutexMapGetConcurrent-16                     280039142               17.10 ns/op            0 B/op          0 allocs/op
BenchmarkCacheGetManyConcurrentExpiring-16              94455712                50.17 ns/op           47 B/op          1 allocs/op
BenchmarkCacheGetManyConcurrentNotExpiring-16           100000000               44.32 ns/op           46 B/op          1 allocs/op
BenchmarkCacheSetExpiring-16                            32226573               153.0 ns/op            72 B/op          3 allocs/op
BenchmarkCacheSetNotExpiring-16                         42172146               118.5 ns/op            72 B/op          3 allocs/op
BenchmarkFreeCacheSet-16                                 7916127               922.5 ns/op           347 B/op          2 allocs/op
BenchmarkBigCacheSet-16                                  7725384               559.7 ns/op           345 B/op          2 allocs/op
BenchmarkRWMutexMapSet-16                               229314998               20.94 ns/op            0 B/op          0 allocs/op
BenchmarkCacheSetDelete-16                              20428626               245.4 ns/op           120 B/op          5 allocs/op
BenchmarkRWMutexMapSetDelete-16                         100000000               47.57 ns/op            0 B/op          0 allocs/op
BenchmarkCacheSetDeleteSingleLock-16                    100000000               46.57 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexMapSetDeleteSingleLock-16               131097927               36.26 ns/op            0 B/op          0 allocs/op
BenchmarkIncrementInt-16                                96997718                51.32 ns/op            8 B/op          1 allocs/op
BenchmarkDeleteExpiredLoop-16                               4623           1048725 ns/op              67 B/op          0 allocs/op
PASS
ok      github.com/pmpavl/rebis 337.580s
```
What conclusions can be drawn from this benchmark data:
- Get: rebis implementation is faster and makes fewer memory allocations than bigcache, freecache.
- Get Concurrent: rebis implementation is faster and makes fewer memory allocations than bigcache, freecache.
- Set: rebis many times faster than freecache and bigcache, but have more memory allocations.
>But you need to remember that freecache and bigcache libraries with more functionality than rebis at this stage.

### References
- [go-cache](https://github.com/patrickmn/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
- [bigcache](https://github.com/allegro/bigcache) - Efficient cache for gigabytes of data written in Go.
- [freecache](https://github.com/coocood/freecache) - A cache library for Go with zero GC overhead.
- [Simple in-memory cache manager in Go](https://habr.com/ru/post/359078/)
- [Under the hood of Redis: Part 0](https://habr.com/ru/post/271487/)
- [Under the hood of Redis: Part 1](https://habr.com/ru/post/271205/)
- [Under the hood of Redis: Part 2](https://habr.com/ru/post/272089/)
- [Optimizing a microservice in Go with a live example](https://habr.com/ru/company/avito/blog/539024/)
