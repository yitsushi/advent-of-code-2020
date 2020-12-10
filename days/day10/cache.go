package day10

type cache struct {
	storage map[int64]int64
}

func newCache() cache {
	return cache{
		storage: map[int64]int64{},
	}
}

func (c *cache) Get(idx int64) int64 {
	return c.storage[idx]
}

func (c *cache) Has(idx int64) bool {
	_, k := c.storage[idx]

	return k
}

func (c *cache) Store(idx int64, value int64) {
	c.storage[idx] = value
}
