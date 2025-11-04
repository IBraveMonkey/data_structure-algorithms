package main

import (
	"fmt"

	"hash/fnv"
	"sync"
)

type ICache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type Shard struct {
	data map[string]string
	mu   *sync.RWMutex
}

type Cache struct {
	shards []*Shard
}

func New(shardCount int64) *Cache {

	shards := make([]*Shard, shardCount) // [shard, shard, shar{mu }]

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}

	return &Cache{
		shards: shards,
	}
}

func (c *Cache) Set(k string, v string) {
	shard := c.getShard(k)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[k] = v
}

func (c *Cache) Get(k string) (string, bool) {
	shard := c.getShard(k)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	if val, ok := shard.data[k]; ok {
		return val, true
	}

	return "", false
}

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
