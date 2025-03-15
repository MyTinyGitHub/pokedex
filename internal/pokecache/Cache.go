package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	m       *sync.Mutex
}

func NewCache(c *Cache, interval time.Duration) Cache {
	if c != nil {
		c.reap(interval)
		return *c
	}

	return Cache{
		entries: make(map[string]cacheEntry),
		m:       &sync.Mutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()

	entry, ok := c.entries[key]

	if ok {
		entry.val = val
	} else {
		entry = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}

	}

	c.entries[key] = entry

	c.m.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if ok {
		return entry.val, ok
	}
	return nil, ok
}

func (c *Cache) reap(interval time.Duration) {
	c.m.Lock()
	for key, value := range c.entries {
		if time.Until(value.createdAt) > interval {
			delete(c.entries, key)
		}
	}
	c.m.Unlock()
}
