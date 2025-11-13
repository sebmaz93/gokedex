package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entires map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entires: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(k string, v []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entires[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	i, ok := c.entires[k]
	return i.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entires {
		if now.Sub(v.createdAt) > last {
			delete(c.entires, k)
		}
	}
}
