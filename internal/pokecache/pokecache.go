package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]CacheEntry
	Mu      sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	entry := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if _, ok := c.Entries[key]; !ok {
		return nil, false
	}
	return c.Entries[key].Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.CreatedAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		Entries: map[string]CacheEntry{},
	}
	go c.reapLoop(interval)
	return &c
}
