package lru

import "fmt"

// Example demonstrates the use of an LRU cache
func Example() {
	// Create an LRU cache with a capacity of 2
	cache := New(2)

	// Add elements
	cache.Put(1, 1) // cache: [1=1]
	fmt.Println("Put(1, 1)")

	cache.Put(2, 2) // cache: [2=2, 1=1]
	fmt.Println("Put(2, 2)")

	// Retrieve an element
	val := cache.Get(1)
	fmt.Printf("Get(1): %d (expected 1)\n", val) // cache: [1=1, 2=2]

	// Add another item (triggers deletion of the least used — 2)
	cache.Put(3, 3) // cache: [3=3, 1=1]
	fmt.Println("Put(3, 3) - evicts key 2")

	val = cache.Get(2)
	fmt.Printf("Get(2): %d (expected -1)\n", val) // 2 was removed

	cache.Put(4, 4) // cache: [4=4, 3=3] - evicts 1
	fmt.Println("Put(4, 4) - evicts key 1")

	val = cache.Get(1)
	fmt.Printf("Get(1): %d (expected -1)\n", val)

	val = cache.Get(3)
	fmt.Printf("Get(3): %d (expected 3)\n", val)

	val = cache.Get(4)
	fmt.Printf("Get(4): %d (expected 4)\n", val)
}
