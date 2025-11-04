package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("key not found")

type elem struct {
	value    string
	exp_date time.Time
}

type Cache struct {
	storage map[string]elem
	mu      *sync.RWMutex
	TTL     time.Duration
	done    chan struct{}
}

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

func (c *Cache) Set(_ context.Context, key, value string) error {

	c.mu.Lock()
	el := elem{
		value:    value,
		exp_date: time.Now().Add(c.TTL),
	}

	c.storage[key] = el
	c.mu.Unlock()

	return nil
}

func (c *Cache) Stop() {
	close(c.done)
}

func (c *Cache) Get(_ context.Context, key string) (string, error) {

	c.mu.RLock()
	el, ok := c.storage[key]
	c.mu.RUnlock()

	if !ok {
		return "", ErrNotFound
	}

	if el.exp_date.Before(time.Now()) {
		c.delete(key)
		return "", ErrNotFound
	}

	return el.value, nil
}

func (c *Cache) Del(_ context.Context, key string) error {
	c.delete(key)
	return nil
}

func (c *Cache) clearByTTL() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.clear()
			case <-c.done:
				// c.Stop()
				return
			}
		}
	}()
}

func (c *Cache) delete(key string) {
	c.mu.Lock()
	delete(c.storage, key)
	c.mu.Unlock()
}

func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, el := range c.storage {
		if el.exp_date.Before(time.Now()) {
			delete(c.storage, key)
		}
	}
}

func main() {

}
