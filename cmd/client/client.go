package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pmpavl/rebis"
)

const (
	durationIntSet = 100 * time.Millisecond
	durationStrSet = 100 * time.Millisecond
	confDir        = "rebisConfig.yaml"
	confDef        = "rebisDefaultConfig.yaml"
)

type keyAndValue struct {
	k string
	v interface{}
}

func workerSet(ctx context.Context, wg *sync.WaitGroup, workerNum int, rebisCache *rebis.Cache, set <-chan keyAndValue) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			// fmt.Printf("worker done: %d\n", workerNum)
			break LOOP
		case kv := <-set:
			rebisCache.Set(kv.k, kv.v, rebis.DefaultExpiration)
		}
	}
}

func workerGet(ctx context.Context, wg *sync.WaitGroup, workerNum int, rebisCache *rebis.Cache, get <-chan string, cancelCh chan<- interface{}) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			// fmt.Printf("worker done: %d\n", workerNum)
			break LOOP
		case k := <-get:
			_, find := rebisCache.Get(k)
			if !find {
				cancelCh <- true
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(0)

	rebisConfig, err := rebis.ConfigFrom(confDir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rebisCache, err := rebis.NewCache(rebisConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	setInt(rebisCache)
	setStr(rebisCache)

	//! parallel reading from cache
	var wg sync.WaitGroup
	getCh := make(chan string, 1)
	cancelCh := make(chan interface{}, 1)
	ctx, finish := context.WithCancel(context.Background())
	defer finish()
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go workerGet(ctx, &wg, i, rebisCache, getCh, cancelCh)
	}

	var i int64 = 0
LOOPGET:
	for {
		select {
		case <-ctx.Done():
			break LOOPGET
		case <-cancelCh:
			finish()
			break LOOPGET
		default:
			if i >= 1000 {
				finish()
				break LOOPGET
			}
			getCh <- fmt.Sprintf("int%d", i)
			getCh <- fmt.Sprintf("str%d", i)
			atomic.AddInt64(&i, 1)
		}
	}
	fmt.Printf("total found: %d\n", i*2)

	close(getCh)
	close(cancelCh)
}

func setInt(rebisCache *rebis.Cache) {
	//! parallel recording in cache
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), durationIntSet)
	defer cancel()
	setCh := make(chan keyAndValue)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go workerSet(ctx, &wg, i, rebisCache, setCh)
	}

	var i int64 = 0
LOOPINT:
	for {
		select {
		case <-ctx.Done():
			break LOOPINT
		default:
			// fmt.Println(i)
			setCh <- keyAndValue{
				k: fmt.Sprintf("int%d", i),
				v: i,
			}
			atomic.AddInt64(&i, 1)
		}
	}

	wg.Wait()
	fmt.Printf("done set int all count: %d\n", i-1)
	close(setCh)
}

func setStr(rebisCache *rebis.Cache) {
	//! parallel recording in cache
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), durationStrSet)
	defer cancel()
	setCh := make(chan keyAndValue)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go workerSet(ctx, &wg, i, rebisCache, setCh)
	}

	var i int64 = 0
LOOPINT:
	for {
		select {
		case <-ctx.Done():
			break LOOPINT
		default:
			// fmt.Println(i)
			setCh <- keyAndValue{
				k: fmt.Sprintf("str%d", i),
				v: fmt.Sprintf("%d", i),
			}
			atomic.AddInt64(&i, 1)
		}
	}

	wg.Wait()
	fmt.Printf("done set str all count: %d\n", i-1)
	close(setCh)
}
