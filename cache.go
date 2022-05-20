package cache

import "time"

var (
	MinTime = time.Unix(-2208988800, 0) // Jan 1, 1900
	MaxTime = MinTime.Add(1<<63 - 1)
)

type Cache struct {
	cache map[string]cache
}

type cache struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{cache: make(map[string]cache)}
}

// func Get returns the value associated with the key
// and the boolean ok (true if exists, false if not),
// if the deadline of the key/value pair has not been exceeded yet
func (c Cache) Get(key string) (string, bool) {
	if v, ok := c.cache[key]; ok {
		if time.Now().Before(v.deadline) {
			return v.value, true
		}
		return v.value, false
	}
	return "", false
}

// func Put places a value with an associated key into cache.
// Value put with this method never expired(have infinite deadline).
// Putting into the existing key should overwrite the value
func (c *Cache) Put(key, value string) {
	c.cache[key] = cache{
		value:    value,
		deadline: MaxTime,
	}
}

// func Keys should return the slice of existing (non-expired keys)
func (c Cache) Keys() []string {
	slice := []string{}
	for k, v := range c.cache {
		if time.Now().Before(v.deadline) {
			slice = append(slice, k)
		}
	}
	return slice
}

// func PutTill should do the same as Put, but the key/value pair should expire by given deadline
func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = cache{
		value:    value,
		deadline: deadline,
	}
}
