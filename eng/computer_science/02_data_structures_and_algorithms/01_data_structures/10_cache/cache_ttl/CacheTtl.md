# CacheTtl

**Description**: 
A Cache with TTL (Time To Live) is a data structure that knows how to "forget" outdated information. Every entry is assigned an expiration date at the time of insertion, after which it is considered invalid.

- **How it works internally**: When saving data, the algorithm records not only the value but also a "time of death" timestamp (`expiry_time = now + TTL`). 
  - *Lazy Deletion*: An element is checked for "freshness" only when it is accessed (Get). If its time is up, it is deleted.
  - *Active Deletion*: A background process (worker) periodically scans the memory and purges data "corpses" to prevent them from wasting space.
- **Analogy**: Imagine a refrigerator. You put milk inside with a receipt showing its expiration date. If you take out the milk and see that the date has passed, you throw it away. Additionally, once a week, you do a full audit and toss anything expired, even if you weren't planning to consume it.


### Pros and Cons
✅ **Pros**:
1. **Data Freshness**: You can be certain that a user won't see information from a hundred years ago.
2. **Auto-Cleanup**: Memory doesn't fill up indefinitely with junk; data naturally "leaves" over time.

❌ **Cons**:
1. **Additional Overhead**: Storing timestamps and running a background cleaner consumes CPU and RAM resources.
2. **Risk of Stale Hits vs. Cache Misses**: If the TTL is too short, you will frequently hit the primary database, negating the purpose of the cache.

---


## 💻 Implementation

```go
package cache_ttl

import (
	"fmt"
	"sync"
	"time"
)

// Item represents a single cached entry with an expiration time
type Item struct {
	Value      interface{}
	ExpiryTime time.Time
}

// TTLCache implements a cache with time-to-live expiration
type TTLCache struct {
	sync.RWMutex
	items      map[string]Item
	defaultTTL time.Duration
}

// New creates a new TTL cache
func New(defaultTTL time.Duration) *TTLCache {
	cache := &TTLCache{
		items:      make(map[string]Item),
		defaultTTL: defaultTTL,
	}
	// Start active background cleanup
	go cache.startCleanup(10 * time.Second)
	return cache
}

// Set adds a value with a specific TTL
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = Item{
		Value:      value,
		ExpiryTime: time.Now().Add(ttl),
	}
}

// Get retrieves a value if it hasn't expired
func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.RLock()
	item, exists := c.items[key]
	c.RUnlock()

	if !exists {
		return nil, false
	}

	// Lazy deletion check
	if time.Now().After(item.ExpiryTime) {
		c.Delete(key)
		return nil, false
	}

	return item.Value, true
}

// Delete removes a key from the cache
func (c *TTLCache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}

// startCleanup periodically removes expired items
func (c *TTLCache) startCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *TTLCache) cleanup() {
	c.Lock()
	defer c.Unlock()
	now := time.Now()
	for key, item := range c.items {
		if now.After(item.ExpiryTime) {
			delete(c.items, key)
		}
	}
}

func main() {
	cache := New(5 * time.Minute)
	cache.Set("status", "active", 1*time.Second)

	val, found := cache.Get("status")
	fmt.Printf("Initial: %v (found: %t)\n", val, found)

	time.Sleep(2 * time.Second)

	val, found = cache.Get("status")
	fmt.Printf("After 2s: %v (found: %t)\n", val, found)
}
```

```javascript
/**
 * TTLCache - simple implementation with lazy and active deletion.
 */
class TTLCache {
  constructor(defaultTTLMs = 60000) {
    this.cache = new Map();
    this.defaultTTL = defaultTTLMs;

    // Active deletion (optional background process)
    setInterval(() => this.cleanup(), 30000);
  }

  /**
   * Set a value with a specific TTL.
   */
  set(key, value, ttlMs = this.defaultTTL) {
    const expiry = Date.now() + ttlMs;
    this.cache.set(key, { value, expiry });
  }

  /**
   * Get a value (uses lazy deletion).
   */
  get(key) {
    const entry = this.cache.get(key);
    if (!entry) return undefined;

    if (Date.now() > entry.expiry) {
      this.cache.delete(key);
      return undefined;
    }

    return entry.value;
  }

  /**
   * Active cleanup of all expired entries.
   */
  cleanup() {
    const now = Date.now();
    for (const [key, entry] of this.cache) {
      if (now > entry.expiry) {
        this.cache.delete(key);
      }
    }
  }

  delete(key) {
    this.cache.delete(key);
  }
}

// Usage example
const cache = new TTLCache(500); // 500ms default TTL
cache.set("temp_token", "12345");
console.log(cache.get("temp_token")); // "12345"

setTimeout(() => {
  console.log(cache.get("temp_token")); // undefined (after 1s)
}, 1000);
```


```


## 🚀 Practical Problems
```go
package cache_ttl

import (
	"context"
	"fmt"
	"time"
)

// Example demonstrates the use of a cache with TTL
func Example() {
	// Create a cache with a 100 ms lifetime
	cache := New(100 * time.Millisecond)
	ctx := context.Background()

	// Add a value
	cache.Set(ctx, "user:1", "Boris Doe")
	fmt.Println("Added 'user:1' -> 'Boris Doe'")

	// Retrieve the value
	val, err := cache.Get(ctx, "user:1")
	if err == nil {
		fmt.Printf("Retrieved: %s\n", val)
	}

	// Wait for TTL expiration
	fmt.Println("Waiting for 200 ms...")
	time.Sleep(200 * time.Millisecond)

	// Attempt to retrieve the obsolete value
	val, err = cache.Get(ctx, "user:1")
	if err != nil {
		fmt.Printf("Retrieval error (expected): %v\n", err)
	} else {
		fmt.Printf("Unexpectedly retrieved: %s\n", val)
	}

	// Stop the cleanup
	cache.Stop()
}

<!-- QUIZ_START 
[
    {
        "question": "What does TTL stand for in caching?",
        "options": ["Total Transaction Length", "Time To Live", "Table Type Level", "True Traffic Limit"],
        "correctIndex": 1
    },
    {
        "question": "What is 'Lazy Deletion' in a TTL cache context?",
        "options": ["Deleting everything at once", "Checking and deleting an expired element only when it's accessed", "Never deleting anything", "Deleting elements based on their size"],
        "correctIndex": 1
    },
    {
        "question": "Why is a background worker (active deletion) often used alongside lazy deletion?",
        "options": ["To make the code more complex", "To ensure 'dead' data is eventually removed even if it's never accessed again", "To speed up the 'Get' operation", "To encrypt the timestamps"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```

