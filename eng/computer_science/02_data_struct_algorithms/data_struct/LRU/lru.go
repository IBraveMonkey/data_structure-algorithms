/*
LRU Cache (Least Recently Used Cache)

What is it?
An LRU (Least Recently Used) cache is a data structure that limits the number of items it can store and, when the limit is reached, removes the items that have not been used for the longest time. When an item is accessed, it becomes "recently used".

Why is it needed?
- To limit memory usage.
- To speed up access to frequently used data.
- To maintain data relevance in the cache.

What's the point?
- Store a limited number of items.
- Mark an item as "recently used" when it's accessed.
- Remove the least recently used items when the cache overflows.

When to use?
- When you need to limit the memory used by the cache.
- When the speed of access to frequently used data is important.
- When you need to automatically delete "old" data.

How does it work?
- A doubly linked list is used to track the order of use.
- A hash table is used for quick access to elements.
- When an item is accessed, it moves to the head of the list (as recently used).
- When the cache overflows, the item at the tail of the list (the least recently used) is removed.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Get (retrieval) | O(1) | O(1) |
| Put (insertion/update) | O(1) | O(1) |
| Deletion | O(1) | O(1) |
| Storage | — | O(n) |

*O(1) is achieved through a combination of a hash table (for lookup) and a doubly linked list (to maintain order).

How to know if a problem fits LRU Cache?
- You need to limit the cache size.
- You need to automatically delete "old" data.
- The frequency of item usage is important.

LRU Cache in Go:

Implementing an LRU cache in Go typically combines:
- A doubly linked list to track usage order.
- A hash table (map) for fast element access.
- Synchronization for multi-threaded access.
- Memory management for efficiency.

Examples of using LRU Cache:
See the example.go file.
*/

package lru

// Node - a doubly linked list node for the LRU cache
type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

// LRUCache - an LRU cache implementation using a doubly linked list and a hash table
type LRUCache struct {
	Capacity int
	Data     map[int]*Node // Hash table for fast access
	Head     *Node         // Virtual head of the list
	Tail     *Node         // Virtual tail of the list
}

// Constructor - creates a new LRU cache with a given capacity
func New(capacity int) LRUCache {
	data := make(map[int]*Node, capacity)

	// Create virtual head and tail to simplify operations
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Capacity: capacity,
		Data:     data,
		Head:     head,
		Tail:     tail,
	}
}

// remove - removes a node from the doubly linked list
func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// addToHead - adds a node to the head of the list (making it recently used)
func (this *LRUCache) addToHead(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

// Get - returns the value by key, marking the item as recently used
func (this *LRUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1 // Key not found
	}

	// Move the node to the head (as recently used)
	this.remove(node)
	this.addToHead(node)
	return node.Value
}

// Put - adds or updates an item in the cache
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		// Update existing item
		node.Value = value
		// Move to head as recently used
		this.remove(node)
		this.addToHead(node)
		return
	}

	// Check if capacity is exceeded
	if len(this.Data) >= this.Capacity {
		// Remove the least recently used element (at the tail of the list)
		leastUsed := this.Tail.Prev
		this.remove(leastUsed)
		delete(this.Data, leastUsed.Key)
	}

	// Create a new node
	newNode := &Node{
		Key:   key,
		Value: value,
	}

	// Add to the head of the list and to the hash table
	this.addToHead(newNode)
	this.Data[key] = newNode
}
