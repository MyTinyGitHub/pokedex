package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

type Cache struct {
	Entries map[string]cacheEntry
	m       *sync.Mutex
}

func NewCache(c *Cache, interval time.Duration) Cache {
	if c != nil {
		c.reap(interval)
		return *c
	}

	return Cache{
		Entries: make(map[string]cacheEntry),
		m:       &sync.Mutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()

	entry, ok := c.Entries[key]

	if ok {
		entry.Val = val
	} else {
		entry = cacheEntry{
			createdAt: time.Now(),
			Val:       val,
		}

	}

	c.Entries[key] = entry

	c.m.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.Entries[key]
	if ok {
		return entry.Val, ok
	}
	return nil, ok
}

func (c *Cache) reap(interval time.Duration) {
	c.m.Lock()
	for key, value := range c.Entries {
		if time.Until(value.createdAt) > interval {
			delete(c.Entries, key)
		}
	}
	c.m.Unlock()
}
