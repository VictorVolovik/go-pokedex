package cache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheData: make(map[string]cacheEntry),
		interval:  interval,
		mux:       &sync.Mutex{},
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheData[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, exists := c.cacheData[key]
	c.mux.Unlock()
	if !exists {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		now := time.Now()
		for key, entry := range c.cacheData {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.cacheData, key)
			}
		}
		c.mux.Unlock()
	}
}
