package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	data map[string]cacheEntry
	mux *sync.Mutex
}

func NewCache(interval time.Duration) *Cache{
	cache := &Cache{
		data: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
} 

func (cache *Cache) Add(key string, val []byte){
	cache.mux.Lock()
	defer cache.mux.Unlock()

	cache.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool){
	cache.mux.Lock()
	defer cache.mux.Unlock()

	val, ok := cache.data[key]
	return val.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for now := range ticker.C {
		cache.reap(now, interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}