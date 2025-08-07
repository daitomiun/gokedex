package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	interval time.Duration
	entries  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		interval: time.Duration(10 * time.Second),
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.entries[key]
	c.mu.Unlock()
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			now := time.Now()
			c.mu.Lock()
			for key, entry := range c.entries {
				println(entry)
				if entry.createdAt.Before(now.Add(-c.interval)) {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
