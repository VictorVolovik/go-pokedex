package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	interval  time.Duration
	mux       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
