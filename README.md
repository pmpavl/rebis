# Rebis [![build](https://github.com/pmpavl/rebis/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/pmpavl/rebis/actions/workflows/go.yaml) [![codecov](https://codecov.io/gh/pmpavl/rebis/branch/master/graph/badge.svg?token=MLE06MIFZD)](https://codecov.io/gh/pmpavl/rebis) [![Go Report Card](https://goreportcard.com/badge/github.com/pmpavl/rebis)](https://goreportcard.com/report/github.com/pmpavl/rebis)

Key-value in-memory concurrent cache storage with: logger, backup save, auto collect expired element, limit on the maximum number of elements. All cache settings are set through the yaml config file.

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
logAll: true            # do standart log or not
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
## Benchmark
Three caches were compared: [rebus](https://github.com/pmpavl/rebis), [bigcache](https://github.com/allegro/bigcache), [freecache](https://github.com/coocood/freecache) and map. Benchmark tests were made using an i7-8550U CPU @ 1.80GHz with 16GB of RAM on Windows 21H1 (19043.1165).
```
go version go1.16.5 windows/amd64

goos: windows
goarch: amd64
pkg: github.com/pmpavl/rebis
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz

BenchmarkCacheGetExpiring-8                     33020690               174.1 ns/op            48 B/op          2 allocs/op
BenchmarkCacheGetNotExpiring-8                  42928779               122.4 ns/op            48 B/op          2 allocs/op
BenchmarkFreeCacheGet-8                          5835063               831.9 ns/op           136 B/op          3 allocs/op
BenchmarkBigCacheGet-8                           6230550               807.5 ns/op           152 B/op          4 allocs/op
BenchmarkRWMutexMapGet-8                        203686885               22.38 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexInterfaceMapGetStruct-8         79589846                54.67 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexInterfaceMapGetString-8         100000000               46.85 ns/op            0 B/op          0 allocs/op
BenchmarkCacheGetConcurrentExpiring-8           45726217               145.9 ns/op            48 B/op          2 allocs/op
BenchmarkCacheGetConcurrentNotExpiring-8        31850788               134.4 ns/op            48 B/op          2 allocs/op
BenchmarkBigCacheGetConcurrent-8                21937942               252.3 ns/op           152 B/op          3 allocs/op
BenchmarkFreeCacheGetConcurrent-8               24994635               513.5 ns/op           136 B/op          3 allocs/op
BenchmarkRWMutexMapGetConcurrent-8              71646922                66.31 ns/op            0 B/op          0 allocs/op
BenchmarkCacheGetManyConcurrentExpiring-8       23866846               188.2 ns/op            46 B/op          1 allocs/op
BenchmarkCacheGetManyConcurrentNotExpiring-8    19545223               243.5 ns/op            47 B/op          1 allocs/op
BenchmarkCacheSetExpiring-8                     15588424               276.8 ns/op            72 B/op          3 allocs/op
BenchmarkCacheSetNotExpiring-8                  23725929               201.7 ns/op            72 B/op          3 allocs/op
BenchmarkFreeCacheSet-8                          7141473              2336 ns/op             355 B/op          2 allocs/op
BenchmarkBigCacheSet-8                           5970564               749.2 ns/op           323 B/op          2 allocs/op
BenchmarkRWMutexMapSet-8                        97942992                48.42 ns/op            0 B/op          0 allocs/op
BenchmarkCacheSetDelete-8                        9910062               437.6 ns/op           120 B/op          5 allocs/op
BenchmarkRWMutexMapSetDelete-8                  44339547                95.26 ns/op            0 B/op          0 allocs/op
BenchmarkCacheSetDeleteSingleLock-8             63755433                84.74 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexMapSetDeleteSingleLock-8        65697086                81.57 ns/op            0 B/op          0 allocs/op
BenchmarkIncrementInt-8                         56427504               107.0 ns/op             8 B/op          1 allocs/op
BenchmarkDeleteExpiredLoop-8                        2696           1784172 ns/op             123 B/op          0 allocs/op
PASS
ok      github.com/pmpavl/rebis 237.023s
```
What conclusions can be drawn from this benchmark data:
- Get: rebis implementation is faster and makes fewer memory allocations than bigcache, freecache.
- Get Concurrent: rebis implementation is faster and makes fewer memory allocations than bigcache, freecache.
- Set: rebis many times faster than freecache and bigcache, but have more memory allocations.
But you need to remember that freecache and bigcache libraries with more functionality than rebis at this stage.

### References
- [go-cache](https://github.com/patrickmn/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
- [bigcache](https://github.com/allegro/bigcache) - Efficient cache for gigabytes of data written in Go.
- [freecache](https://github.com/coocood/freecache) - A cache library for Go with zero GC overhead.
- [Simple in-memory cache manager in Go](https://habr.com/ru/post/359078/)
- [Under the hood of Redis: Part 0](https://habr.com/ru/post/271487/)
- [Under the hood of Redis: Part 1](https://habr.com/ru/post/271205/)
- [Under the hood of Redis: Part 2](https://habr.com/ru/post/272089/)
- [Optimizing a microservice in Go with a live example](https://habr.com/ru/company/avito/blog/539024/)
