# CacheShard

**Description**: 
Cache Sharding is a method of scaling a cache by dividing it into several independent fragments ("shards"). In multi-threaded programs, this helps eliminate "bottlenecks" when accessing data.

- **How it works internally**: If you have a million entries in a single hash table protected by one mutex (lock), thousands of threads will queue up just to read a single value. Sharding divides that million into 10-20 smaller tables, each with its own lock. A hash function determines which shard will hold the data. Now, threads can work in parallel as long as they access different shards.
- **Analogy**: Imagine a supermarket with a single checkout counter. That is a cache without sharding. Sharding is like opening 10 checkout counters. Customers distribute themselves among the different queues, and the overall service speed increases dramatically.


### Pros and Cons
✅ **Pros**:
1. **High Concurrency**: Allows many threads to read and write data simultaneously without waiting for one another.
2. **Reduced CPU Overhead**: The processor spends less time managing locks and switching tasks.

❌ **Cons**:
1. **Implementation Complexity**: Requires ensuring that the same key always lands in the same shard (requires an effective hash function).
2. **Data Imbalance**: If the hash function is poor, one shard may become overloaded while others remain idle.

---


## 💻 Implementation

```go
package cache_shard

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

// Shard represents a single segment of the cache with its own lock
type Shard struct {
	sync.RWMutex
	data map[string]interface{}
}

// Cache implements a sharded cache to reduce lock contention
type Cache struct {
	shardCount int
	shards     []*Shard
}

// New creates a new sharded cache
func New(shardCount int) *Cache {
	shards := make([]*Shard, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = &Shard{
			data: make(map[string]interface{}),
		}
	}
	return &Cache{
		shardCount: shardCount,
		shards:     shards,
	}
}

// getShardIndex calculates which shard holds the key
func (c *Cache) getShardIndex(key string) int {
	hash := sha1.Sum([]byte(key))
	return int(hash[0]) % c.shardCount
}

// Set adds or updates a value in the appropriate shard
func (c *Cache) Set(key string, value interface{}) {
	shard := c.shards[c.getShardIndex(key)]
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

// Get retrieves a value from the appropriate shard
func (c *Cache) Get(key string) (interface{}, bool) {
	shard := c.shards[c.getShardIndex(key)]
	shard.RLock()
	defer shard.RUnlock()
	val, exists := shard.data[key]
	return val, exists
}

// Delete removes a key from its shard
func (c *Cache) Delete(key string) {
	shard := c.shards[c.getShardIndex(key)]
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

func main() {
	cache := New(8)
	cache.Set("user:101", "Boris")
	
	if val, ok := cache.Get("user:101"); ok {
		fmt.Printf("Found in cache: %v\n", val)
	}
}
```

```javascript
/**
 * ShardedCache - logical partitioning for scalability.
 */
class ShardedCache {
  constructor(shardCount = 8) {
    this.shardCount = shardCount;
    this.shards = Array.from({ length: shardCount }, () => new Map());
  }

  /**
   * Simple hash function for keys.
   */
  _hash(key) {
    let hash = 0;
    for (let i = 0; i < key.length; i++) {
      hash = (hash << 5) - hash + key.charCodeAt(i);
      hash |= 0;
    }
    return Math.abs(hash);
  }

  _getShardIndex(key) {
    return this._hash(key) % this.shardCount;
  }

  /**
   * Set a key-value pair in the specific shard.
   */
  set(key, value) {
    const index = this._getShardIndex(key);
    this.shards[index].set(key, value);
  }

  /**
   * Get a value from the specific shard.
   */
  get(key) {
    const index = this._getShardIndex(key);
    return this.shards[index].get(key);
  }

  delete(key) {
    const index = this._getShardIndex(key);
    this.shards[index].delete(key);
  }

  stats() {
    return this.shards.map((s, i) => `Shard ${i}: ${s.size} items`);
  }
}

// Usage example
const cache = new ShardedCache(4);
cache.set("session_id", "ABC-123");
console.log(cache.get("session_id"));
console.log(cache.stats());
```


<!-- QUIZ_START 
[
    {
        "question": "What is the primary goal of Cache Sharding in multi-threaded programs?",
        "options": ["To encrypt data", "To reduce lock contention by dividing data into independently locked fragments", "To sort data automatically", "To use less memory"],
        "correctIndex": 1
    },
    {
        "question": "What determines which shard a specific key will be stored in?",
        "options": ["The size of the value", "A random number generator", "A hash function applied to the key", "The time of insertion"],
        "correctIndex": 2
    },
    {
        "question": "What is a potential downside of having too few shards?",
        "options": ["Data will be lost", "Threads may still experience high wait times for locks (bottleneck)", "The hash function will stop working", "Memory will be used inefficiently"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```


## 🚀 Practical Problems
