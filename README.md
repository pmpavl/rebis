# Rebis [![build](https://github.com/pmpavl/rebis/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/pmpavl/rebis/actions/workflows/go.yaml) [![codecov](https://codecov.io/gh/pmpavl/rebis/branch/master/graph/badge.svg?token=MLE06MIFZD)](https://codecov.io/gh/pmpavl/rebis) [![Go Report Card](https://goreportcard.com/badge/github.com/pmpavl/rebis)](https://goreportcard.com/report/github.com/pmpavl/rebis)

Key-value in-memory concurrent cache storage with: logger, backup save, auto collect expired element, limit on the maximum number of elements. All cache settings are set through the yaml config file.

Requires Go 1.14 or newer.



## Usage
### Simple Usage



### References
- [go-cache](https://github.com/patrickmn/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
- [bigcache](https://github.com/allegro/bigcache) - Efficient cache for gigabytes of data written in Go.
- [freecache](https://github.com/coocood/freecache) - A cache library for Go with zero GC overhead.
- [Simple in-memory cache manager in Go](https://habr.com/ru/post/359078/)
- [Under the hood of Redis: Part 0](https://habr.com/ru/post/271487/)
- [Under the hood of Redis: Part 1](https://habr.com/ru/post/271205/)
- [Under the hood of Redis: Part 2](https://habr.com/ru/post/272089/)
- [Optimizing a microservice in Go with a live example](https://habr.com/ru/company/avito/blog/539024/)
