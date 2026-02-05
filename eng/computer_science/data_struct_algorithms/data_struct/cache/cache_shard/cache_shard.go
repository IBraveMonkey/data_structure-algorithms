/*
Cache Sharding

What is it?
Cache sharding is a method of distributing cache data across multiple independent segments (shards). Each shard is a separate block of memory with its own synchronization, which reduces locking and increases performance during multi-threaded access.

Why is it needed?
- Increase concurrency when working with the cache.
- Reduce contention (wait for locks) during parallel access.
- Distribute the load among several segments.

What's the point?
- Splitting the cache into several independent parts.
- Using a key hash to determine the correct shard.
- Each shard manages access to its own data independently.

When to use?
- When the cache is used in a multi-threaded environment.
- Under high load with frequent read/write operations.
- When lock contention becomes a bottleneck.
*/

package main

import (
	"fmt"

	"hash/fnv"
	"sync"
)

// ICache - interface for the cache
type ICache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

// Shard - a cache segment with its own synchronization
type Shard struct {
	data map[string]string
	mu   *sync.RWMutex
}

// Cache - a cache structure consisting of several shards
type Cache struct {
	shards []*Shard
}

// New - creates a new sharded cache with the specified number of shards
func New(shardCount int64) *Cache {
	shards := make([]*Shard, shardCount)

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
			mu:   &sync.RWMutex{}, // Adding a mutex for each shard
		}
	}

	return &Cache{
		shards: shards,
	}
}

// Set - adds a key-value pair to the cache
func (c *Cache) Set(k string, v string) {
	shard := c.getShard(k)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[k] = v
}

// Get - returns the value for a key from the cache
func (c *Cache) Get(k string) (string, bool) {
	shard := c.getShard(k)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	if val, ok := shard.data[k]; ok {
		return val, true
	}

	return "", false
}

// getShard - returns the shard for the specified key using hashing
func (c *Cache) getShard(key string) *Shard {
	hasher := fnv.New32()
	_, _ = hasher.Write([]byte(key))
	hash := hasher.Sum32()

	return c.shards[hash%uint32(len(c.shards))]
}

func main() {
	cache := New(10)
	cache.Set("salary", "500000")

	value, ok := cache.Get("salary")
	if !ok {
		fmt.Println("Value not found")
		return
	}

	fmt.Println("Value:", value)
}
