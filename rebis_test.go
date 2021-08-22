package rebis

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/allegro/bigcache/v2"
	"github.com/coocood/freecache"
)

var (
	config = &Config{
		Size: 8192,
		Backup: Backup{
			InUse: false,
		},
		DefaultExpiration: time.Duration(-1),
		CleanupInterval:   time.Duration(time.Minute * 10),
		LogAll:            false,
		Evicted:           false,
	}
)

const maxEntrySize = 256

type TestStruct struct {
	Num      int
	Children []*TestStruct
}

func key(i int) string {
	return fmt.Sprintf("key-%010d", i)
}

func value() []byte {
	return make([]byte, 100)
}

func initBigCache(entriesInWindow int) *bigcache.BigCache {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             256,
		LifeWindow:         10 * time.Minute,
		MaxEntriesInWindow: entriesInWindow,
		MaxEntrySize:       maxEntrySize,
		Verbose:            true,
	})

	return cache
}

func TestCache(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}

	a, found := tc.Get("a")
	if found || a != nil {
		t.Error("Getting A found value that shouldn't exist:", a)
	}

	b, found := tc.Get("b")
	if found || b != nil {
		t.Error("Getting B found value that shouldn't exist:", b)
	}

	c, found := tc.Get("c")
	if found || c != nil {
		t.Error("Getting C found value that shouldn't exist:", c)
	}

	tc.Set("a", 1, DefaultExpiration)
	tc.Set("b", "b", DefaultExpiration)
	tc.Set("c", 3.5, DefaultExpiration)

	x, found := tc.Get("a")
	if !found {
		t.Error("a was not found while getting a2")
	}
	if x == nil {
		t.Error("x for a is nil")
	} else if a2 := x.(int); a2+2 != 3 {
		t.Error("a2 (which should be 1) plus 2 does not equal 3; value:", a2)
	}

	x, found = tc.Get("b")
	if !found {
		t.Error("b was not found while getting b2")
	}
	if x == nil {
		t.Error("x for b is nil")
	} else if b2 := x.(string); b2+"B" != "bB" {
		t.Error("b2 (which should be b) plus B does not equal bB; value:", b2)
	}

	x, found = tc.Get("c")
	if !found {
		t.Error("c was not found while getting c2")
	}
	if x == nil {
		t.Error("x for c is nil")
	} else if c2 := x.(float64); c2+1.2 != 4.7 {
		t.Error("c2 (which should be 3.5) plus 1.2 does not equal 4.7; value:", c2)
	}
}

func TestCacheTimes(t *testing.T) {
	var found bool
	conf := configDefault()
	conf.DefaultExpiration = 50 * time.Millisecond
	conf.CleanupInterval = 10 * time.Millisecond
	tc, err := NewCache(conf)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("a", 1, DefaultExpiration)
	tc.Set("b", 2, NoExpiration)
	tc.Set("c", 3, 20*time.Millisecond)
	tc.set("c", 3, 20*time.Millisecond)
	tc.Set("d", 4, 80*time.Millisecond)

	<-time.After(25 * time.Millisecond)
	_, _, found = tc.GetWithExpiration("c")
	if found {
		t.Error("Found c when it should have been automatically deleted")
	}

	<-time.After(30 * time.Millisecond)
	_, found = tc.Get("a")
	if found {
		t.Error("Found a when it should have been automatically deleted")
	}

	_, _, found = tc.GetWithExpiration("b")
	if !found {
		t.Error("Did not find b even though it was set to never expire")
	}

	_, found = tc.Get("d")
	if !found {
		t.Error("Did not find d even though it was set to expire later than the default")
	}

	<-time.After(50 * time.Millisecond)
	tc.Get("d")
	_, _, found = tc.GetWithExpiration("d")
	if found {
		t.Error("Found d when it should have been automatically deleted (later than the default)")
	}
}

func TestNewFrom(t *testing.T) {
	m := map[string]Item{
		"a": {
			Value:      1,
			Expiration: 0,
		},
		"b": {
			Value:      2,
			Expiration: 0,
		},
	}
	tc, err := NewCacheFrom(config, m)
	if err != nil {
		t.Error("err with default config")
	}
	a, found := tc.Get("a")
	if !found {
		t.Fatal("Did not find a")
	}
	if a.(int) != 1 {
		t.Fatal("a is not 1")
	}
	b, found := tc.Get("b")
	if !found {
		t.Fatal("Did not find b")
	}
	if b.(int) != 2 {
		t.Fatal("b is not 2")
	}
}

func TestStorePointerToStruct(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", &TestStruct{Num: 1}, DefaultExpiration)
	x, found := tc.Get("foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo")
	}
	foo := x.(*TestStruct)
	foo.Num++

	y, found := tc.Get("foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo (second time)")
	}
	bar := y.(*TestStruct)
	if bar.Num != 2 {
		t.Fatal("TestStruct.Num is not 2")
	}
}

func TestAdd(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	err = tc.Add("foo", "bar", DefaultExpiration)
	if err != nil {
		t.Error("Couldn't add foo even though it shouldn't exist")
	}
	err = tc.Add("foo", "baz", DefaultExpiration)
	if err == nil {
		t.Error("Successfully added another foo when it should have returned an error")
	}
}

func TestReplace(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	err = tc.Replace("foo", "bar", DefaultExpiration)
	if err == nil {
		t.Error("Replaced foo when it shouldn't exist")
	}
	tc.Set("foo", "bar", DefaultExpiration)
	err = tc.Replace("foo", "bar", DefaultExpiration)
	if err != nil {
		t.Error("Couldn't replace existing key foo")
	}
}

func TestDelete(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", "bar", DefaultExpiration)
	tc.Delete("foo")
	x, found := tc.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
}

func TestSetEmpty(t *testing.T) {
	conf := configDefault()
	conf.Size = 1
	tc, err := NewCache(conf)
	if err != nil {
		t.Error("err with default config")
	}
	for i := 0; i < 100; i++ {
		tc.SetDefault(strconv.Itoa(i), i)
	}
	for i := 0; i < 100; i++ {
		tc.Add(strconv.Itoa(i), i, DefaultExpiration)
	}
}

func TestItemCount(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", "1", DefaultExpiration)
	tc.Set("bar", "2", DefaultExpiration)
	tc.Set("baz", "3", DefaultExpiration)
	if n := tc.ItemCount(); n != 3 {
		t.Errorf("Item count is not 3: %d", n)
	}
}

func TestItems(t *testing.T) {
	conf := configDefault()
	conf.CleanupInterval = 1000 * time.Millisecond
	tc, err := NewCache(conf)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", "1", -1)
	tc.Set("bar", "2", -1)
	tc.Set("baz", "3", -1)
	tc.Set("not", 0, 10*time.Millisecond)
	<-time.After(100 * time.Millisecond)
	items := tc.Items()
	if items["foo"].Value != "1" || items["bar"].Value != "2" || items["baz"].Value != "3" {
		t.Errorf("Items get wrong, items %+v", items)
	}
	if _, ok := items["not"]; ok {
		t.Errorf("Items not expired")
	}

	stopJanitor(tc)
}

func TestLogger(t *testing.T) {
	conf := configDefault()
	conf.LogAll = true
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.ChangeLogger(log.New(os.Stdout, "", log.LstdFlags))
	tc.logIf("test logger check")
}

func TestFlush(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", "bar", DefaultExpiration)
	tc.Set("baz", "yes", DefaultExpiration)
	tc.Flush()
	x, found := tc.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
	x, found = tc.Get("baz")
	if found {
		t.Error("baz was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
}

func TestRunEvicted(t *testing.T) {
	conf := configDefault()
	conf.Evicted = true
	conf.CleanupInterval = 5 * time.Millisecond
	tc, err := NewCache(conf)
	if err != nil {
		t.Error("err with evicted config")
	}
	tc.Set("foo", 3, 1*time.Millisecond)
	<-time.After(10 * time.Millisecond)
}

func TestOnEvicted(t *testing.T) {
	tc, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	tc.Set("foo", 3, DefaultExpiration)
	works := false
	tc.OnEvicted(func(k string, v interface{}) {
		if k == "foo" && v.(int) == 3 {
			works = true
		}
		tc.Set("bar", 4, DefaultExpiration)
	})
	tc.Delete("foo")
	x, _ := tc.Get("bar")
	if !works {
		t.Error("works bool not true")
	}
	if x.(int) != 4 {
		t.Error("bar was not 4")
	}
}

func TestRunBackup(t *testing.T) {
	conf := configDefault()
	conf.Backup = Backup{
		InUse:    true,
		Interval: time.Millisecond * 2,
		Path:     "./qwe",
	}
	_, err := NewCache(conf)
	if err != nil {
		t.Errorf("wrong open json")
	}

	conf.Backup.Path = "./backup"
	tc, err := NewCache(conf)
	if err != nil {
		t.Errorf("err with backup config, %s", err.Error())
	}

	err = tc.BackupSave()
	if err != nil {
		t.Errorf("err with backup save, %s", err.Error())
	}

	err = tc.BackupSaveFile("./qwe/wqr")
	if err == nil {
		t.Errorf("not check file open error")
	}

	err = tc.BackupRecovery()
	if err != nil {
		t.Errorf("could not recoverly")
	}

	<-time.After(time.Millisecond * 20)
	stopBackup(tc)
}

func TestConfig(t *testing.T) {
	err := ConfigCreateDefault("qwe")
	if err.Error() != "config file should be in yaml format" {
		t.Errorf("not check format config")
	}
	err = ConfigCreateDefault("./qwer/aqw.yaml")
	if err == nil {
		t.Errorf("not check correct path")
	}

	err = ConfigCreateDefault("test.yaml")
	if err != nil {
		t.Errorf("not create config")
	}

	_, err = ConfigFrom("./qwer/aqw.yaml")
	if err == nil {
		t.Errorf("not check readeble file")
	}

	config, _ = ConfigFrom("test.yaml")
	if config.CleanupInterval != configDefault().CleanupInterval {
		t.Errorf("not create default")
	}
}

func TestBackupEmptySlot(t *testing.T) {
	conf := configDefault()
	conf.Backup = Backup{
		InUse:    true,
		Interval: time.Second * 5,
		Path:     "./backup",
	}
	conf.Size = 10
	tc, _ := NewCache(conf)

	for i := 0; i < 100; i++ {
		tc.Set("foo"+strconv.Itoa(i), i, 0)
	}
	tc.BackupSaveFile("test.json")

	conf.Size = 1
	tc, _ = NewCache(conf)
	err := tc.BackupRecoveryFile("test.json")
	if err.Error() != "no empty slot, for next items" {
		t.Errorf("no recover empty memmory")
	}
}

func TestCacheBackup(t *testing.T) {
	ts, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	testBackupSave(t, ts)

	tr, err := NewCache(config)
	if err != nil {
		t.Error("err with default config")
	}
	testBackupRecovery(t, tr)
}

func testBackupSave(t *testing.T, tc *Cache) {
	tc.Set("a", "a", DefaultExpiration)
	tc.Set("b", "b", DefaultExpiration)
	tc.Set("c", "c", DefaultExpiration)
	tc.Set("expired", "foo", 1*time.Millisecond)

	err := tc.BackupSaveFile("test.json")
	if err != nil {
		t.Fatal("Couldn't save cache to test.json:", err)
	}

}

func testBackupRecovery(t *testing.T, tr *Cache) {
	err := tr.BackupRecoveryFile("test.json")
	if err != nil {
		t.Fatal("Couldn't load cache from test.json:", err)
	}
	a, found := tr.Get("a")
	if !found {
		t.Error("a was not found")
	}
	if a.(string) != "a" {
		t.Error("a is not a")
	}
	b, found := tr.Get("b")
	if !found {
		t.Error("b was not found")
	}
	if b.(string) != "b" {
		t.Error("b is not b")
	}

	c, found := tr.Get("c")
	if !found {
		t.Error("c was not found")
	}
	if c.(string) != "c" {
		t.Error("c is not c")
	}

	<-time.After(5 * time.Millisecond)
	_, found = tr.Get("expired")
	if found {
		t.Error("expired was found")
	}
}

func BenchmarkCacheGetExpiring(b *testing.B) {
	benchmarkCacheGet(b, 5*time.Minute)
}

func BenchmarkCacheGetNotExpiring(b *testing.B) {
	benchmarkCacheGet(b, NoExpiration)
}

func benchmarkCacheGet(b *testing.B, exp time.Duration) {
	conf := configDefault()
	b.StopTimer()
	conf.DefaultExpiration = exp
	tc, err := NewCache(conf)
	if err != nil {
		b.Errorf("err with default config")
	}
	tc.Set("foo", "bar", DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Get("foo")
	}
}

func BenchmarkFreeCacheGet(b *testing.B) {
	b.StopTimer()
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		cache.Set([]byte(key(i)), value(), 0)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache.Get([]byte(key(i)))
	}
}

func BenchmarkBigCacheGet(b *testing.B) {
	b.StopTimer()
	cache := initBigCache(b.N)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value())
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(key(i))
	}
}

func BenchmarkRWMutexMapGet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetStruct(b *testing.B) {
	b.StopTimer()
	s := struct{ name string }{name: "foo"}
	m := map[interface{}]string{
		s: "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m[s]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetString(b *testing.B) {
	b.StopTimer()
	m := map[interface{}]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkCacheGetConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, NoExpiration)
}

func benchmarkCacheGetConcurrent(b *testing.B, exp time.Duration) {
	conf := configDefault()
	b.StopTimer()
	conf.DefaultExpiration = exp
	tc, err := NewCache(conf)
	if err != nil {
		b.Errorf("err with default config")
	}
	tc.Set("foo", "bar", DefaultExpiration)
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				tc.Get("foo")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkBigCacheGetConcurrent(b *testing.B) {
	b.StopTimer()
	cache := initBigCache(b.N)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value())
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			cache.Get(key(counter))
			counter = counter + 1
		}
	})
}

func BenchmarkFreeCacheGetConcurrent(b *testing.B) {
	b.StopTimer()
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		cache.Set([]byte(key(i)), value(), 0)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			cache.Get([]byte(key(counter)))
			counter = counter + 1
		}
	})
}

func BenchmarkRWMutexMapGetConcurrent(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				mu.RLock()
				_, _ = m["foo"]
				mu.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkCacheGetManyConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetManyConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, NoExpiration)
}

func benchmarkCacheGetManyConcurrent(b *testing.B, exp time.Duration) {
	// This is the same as BenchmarkCacheGetConcurrent, but its result
	// can be compared against BenchmarkShardedCacheGetManyConcurrent
	// in sharded_test.go.
	b.StopTimer()
	n := 10000
	conf := configDefault()
	conf.DefaultExpiration = exp
	tc, err := NewCache(conf)
	if err != nil {
		b.Errorf("err with default config")
	}
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		k := "foo" + strconv.Itoa(i)
		keys[i] = k
		tc.Set(k, "bar", DefaultExpiration)
	}
	each := b.N / n
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, v := range keys {
		go func(k string) {
			for j := 0; j < each; j++ {
				tc.Get(k)
			}
			wg.Done()
		}(v)
	}
	b.StartTimer()
	wg.Wait()
}

func BenchmarkCacheSetExpiring(b *testing.B) {
	benchmarkCacheSet(b, 5*time.Minute)
}

func BenchmarkCacheSetNotExpiring(b *testing.B) {
	benchmarkCacheSet(b, NoExpiration)
}

func benchmarkCacheSet(b *testing.B, exp time.Duration) {
	b.StopTimer()
	conf := configDefault()
	conf.DefaultExpiration = exp
	tc, err := NewCache(conf)
	if err != nil {
		b.Errorf("err with default config")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpiration)
	}
}
func BenchmarkFreeCacheSet(b *testing.B) {
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		cache.Set([]byte(key(i)), value(), 0)
	}
}

func BenchmarkBigCacheSet(b *testing.B) {
	cache := initBigCache(b.N)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value())
	}
}

func BenchmarkRWMutexMapSet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
	}
}

func BenchmarkCacheSetDelete(b *testing.B) {
	b.StopTimer()
	tc, err := NewCache(config)
	if err != nil {
		b.Errorf("err with default config")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpiration)
		tc.Delete("foo")
	}
}

func BenchmarkRWMutexMapSetDelete(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
		mu.Lock()
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkCacheSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	tc, err := NewCache(config)
	if err != nil {
		b.Errorf("err with default config")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.mu.Lock()
		tc.set("foo", "bar", DefaultExpiration)
		tc.delete("foo")
		tc.mu.Unlock()
	}
}

func BenchmarkRWMutexMapSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkIncrementInt(b *testing.B) {
	b.StopTimer()
	tc, err := NewCache(config)
	if err != nil {
		b.Errorf("err with default config")
	}
	tc.Set("foo", 0, DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.IncrementInt("foo", 1)
	}
}
func BenchmarkDeleteExpiredLoop(b *testing.B) {
	b.StopTimer()
	conf := configDefault()
	conf.DefaultExpiration = 5 * time.Minute
	tc, err := NewCache(conf)
	if err != nil {
		b.Errorf("err with default config")
	}
	tc.mu.Lock()
	for i := 0; i < 100000; i++ {
		tc.set(strconv.Itoa(i), "bar", DefaultExpiration)
	}
	tc.mu.Unlock()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.DeleteExpired()
	}
}
