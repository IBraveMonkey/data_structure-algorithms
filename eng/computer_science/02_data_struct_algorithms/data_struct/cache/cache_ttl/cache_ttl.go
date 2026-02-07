/*
Cache with TTL (Time To Live)

What is it?
A cache with TTL is a data structure where each element has a limited lifetime. After the lifetime expires, the element is automatically removed from the cache, freeing up memory and ensuring data relevance.

Why is it needed?
- Automatically delete obsolete data.
- Limit the amount of memory used.
- Ensure the freshness of data in the cache.

What's the point?
- Each element is assigned a lifetime when added.
- Elements are deleted after their lifetime expires.
- Data relevance is maintained without manual cleanup.

When to use?
- When data in the cache can become obsolete.
- To limit memory usage.
- When it's important to have up-to-date data.
*/

package cache_ttl

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("key not found")

// elem - a cache element with data and a lifetime
type elem struct {
	value    string
	exp_date time.Time
}

// Cache - a cache structure with TTL
type Cache struct {
	storage map[string]elem
	mu      *sync.RWMutex
	TTL     time.Duration // Elements' lifetime
	done    chan struct{} // Channel to stop the cleanup goroutine
}

// New - creates a new cache with the specified TTL
func New(ttl time.Duration) *Cache {
	cache := &Cache{
		storage: make(map[string]elem),
		mu:      &sync.RWMutex{},
		TTL:     ttl,
		done:    make(chan struct{}),
	}

	cache.clearByTTL()

	return cache
}

// Set - adds a key-value pair to the cache with TTL
func (c *Cache) Set(_ context.Context, key, value string) error {

	c.mu.Lock()
	el := elem{
		value:    value,
		exp_date: time.Now().Add(c.TTL), // Setting the lifetime
	}

	c.storage[key] = el
	c.mu.Unlock()

	return nil
}

// Stop - stops the background cache cleanup
func (c *Cache) Stop() {
	close(c.done)
}

// Get - returns the value for a key from the cache, checking the TTL
func (c *Cache) Get(_ context.Context, key string) (string, error) {

	c.mu.RLock()
	el, ok := c.storage[key]
	c.mu.RUnlock()

	if !ok {
		return "", ErrNotFound
	}

	// Check if the lifetime has expired
	if el.exp_date.Before(time.Now()) {
		c.delete(key) // Remove the expired element
		return "", ErrNotFound
	}

	return el.value, nil
}

// Del - removes an element from the cache
func (c *Cache) Del(_ context.Context, key string) error {
	c.delete(key)
	return nil
}

// clearByTTL - starts the background cleanup of expired elements
func (c *Cache) clearByTTL() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.clear() // Cleanup expired elements
			case <-c.done:
				return
			}
		}
	}()
}

// delete - removes an element from the cache
func (c *Cache) delete(key string) {
	c.mu.Lock()
	delete(c.storage, key)
	c.mu.Unlock()
}

// clear - removes all expired elements from the cache
func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, el := range c.storage {
		if el.exp_date.Before(time.Now()) {
			delete(c.storage, key)
		}
	}
}
