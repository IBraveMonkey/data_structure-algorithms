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
	cache.Set(ctx, "user:1", "John Doe")
	fmt.Println("Added 'user:1' -> 'John Doe'")

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
