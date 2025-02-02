package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	reapInterval time.Duration
	Store        map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	newCache := Cache{
		Store:        make(map[string]cacheEntry),
		reapInterval: duration,
	}
	return &newCache
}

func (c *Cache) Add(name string, data []byte) error {
	c.Lock()
	defer c.Unlock()
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	c.Store[name] = ce
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	entry, exists := c.Store[key]
	if !exists || time.Since(entry.createdAt) > c.reapInterval {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	for {
		<-ticker.C
		c.Lock()
		for k, v := range c.Store {
			if time.Since(v.createdAt) > c.reapInterval {
				delete(c.Store, k)
			}
		}
		c.Unlock()
	}
}
