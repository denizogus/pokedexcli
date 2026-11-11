package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data  map[string]cacheEntry
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Data:  map[string]cacheEntry{},
		mutex: sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if data, ok := c.Data[key]; ok {
		return data.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		c.mutex.Lock()

		for key, value := range c.Data {
			if time.Since(value.createdAt) > interval {
				delete(c.Data, key)
			}
		}

		c.mutex.Unlock()
	}
}
