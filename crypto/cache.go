package crypto

import (
	"fmt"
	"sync"
)

type Entry struct {
	data []Response
}

type Cache struct {
	cache map[string]*Entry
	sync.Mutex
}

func NewCache() *Cache {
	cache := make(map[string]*Entry)
	return &Cache{
		cache: cache,
	}
}

func (c *Cache) Read(symbol string) ([]Response, error) {
	entry, found := c.cache[symbol]
	if !found {
		return []Response{}, fmt.Errorf("no entries found for symbol %s", symbol)
	}
	return entry.data, nil
}

func (c *Cache) Add(symbol string, data Response) {
	c.Lock()

	if _, found := c.cache[symbol]; !found {
		empty := Entry{}
		c.cache[symbol] = &empty
	}

	c.cache[symbol].data = append(c.cache[symbol].data, data)

	c.Unlock()
}
