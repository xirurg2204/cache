package cache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	value interface{}
	ttl   time.Time
}
type Cache struct {
	data map[string]cacheEntry
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]cacheEntry),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		value: value,
		ttl:   time.Now().Add(ttl),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, error) {
	if entry, ok := c.data[key]; ok {
		if time.Now().After(entry.ttl) {
			c.Delete(key)
			return nil, errors.New("ttl expired")
		}
		c.mu.RLock()
		defer c.mu.RUnlock()
		return entry.value, nil
	}
	return nil, errors.New("key not found")
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}
