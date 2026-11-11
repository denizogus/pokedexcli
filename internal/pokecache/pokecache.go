package pokecache

type Cache struct {
	Data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		Data: map[string][]byte{},
	}
}

func (c *Cache) Add(key string, data []byte) {
	c.Data[key] = data
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if data, ok := c.Data[key]; ok {
		return data, true
	}
	return nil, false
}
