# Lru

**Description**: 
LRU Cache (Least Recently Used) is one of the most popular caching strategies, helping programs "remember" what's important and discard what's no longer needed in a timely manner. Its core principle is: "If we haven't used something in a long time, we probably won't need it in the future either."

- **How it works internally**: LRU is a tandem of two structures. A **Hash Table** provides instantaneous data lookup by key (O(1)), while a **Doubly Linked List** stores elements in the order they were used.
  - When you read or write an item, it moves to the "head" of the list (the most recent).
  - Items that haven't been accessed in a while gradually "slide" toward the tail.
  - Once the cache reaches capacity, the item at the very tail (the "oldest") is removed forever.
- **Analogy**: Imagine your desk. The most important documents are right in front of you. Those you pulled out yesterday are a bit further away. If your desk fills up completely, you'll take the piece of paper you haven't touched the longest and throw it in the trash to make room for something new.


### Pros and Cons
✅ **Pros**:
1. **Fixed Memory**: You know exactly how much memory the cache consumes, and it will never grow beyond its limit.
2. **High Speed**: Lookup, insertion, and deletion all operate in **O(1)**.
3. **Predictability**: Performs excellently in scenarios where the same data is requested frequently (principle of locality).

❌ **Cons**:
1. **Implementation Complexity**: Requires careful management of list nodes and the map simultaneously.
2. **Memory Overhead**: Storing pointers (Next/Prev) in the list requires more memory than a simple array would.

---


## 💻 Implementation

```go
package lru

// Node represents a node in the doubly linked list
type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

// LRUCache implements an LRU cache with O(1) operations
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // Virtual head (most recent)
	tail     *Node // Virtual tail (least recent)
}

// New creates a new LRU cache with the specified capacity
func New(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

// remove deletes a node from the linked list
func (lru *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// addToHead inserts a node at the front of the list
func (lru *LRUCache) addToHead(node *Node) {
	node.Next = lru.head.Next
	node.Prev = lru.head
	lru.head.Next.Prev = node
	lru.head.Next = node
}

// Get retrieves a value and moves the item to the front
func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if !exists {
		return -1
	}
	lru.remove(node)
	lru.addToHead(node)
	return node.Value
}

// Put adds or updates a value, evicting the oldest if needed
func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		node.Value = value
		lru.remove(node)
		lru.addToHead(node)
	} else {
		newNode := &Node{Key: key, Value: value}
		if len(lru.cache) >= lru.capacity {
			last := lru.tail.Prev
			lru.remove(last)
			delete(lru.cache, last.Key)
		}
		lru.addToHead(newNode)
		lru.cache[key] = newNode
	}
}
```

```javascript
/**
 * Node for the doubly linked list.
 */
class Node {
  constructor(key, value) {
    this.key = key;
    this.value = value;
    this.prev = null;
    this.next = null;
  }
}

/**
 * LRUCache implementation using Map and Doubly Linked List.
 */
class LRUCache {
  constructor(capacity) {
    this.capacity = capacity;
    this.cache = new Map();
    this.head = new Node(0, 0); // Virtual head
    this.tail = new Node(0, 0); // Virtual tail
    this.head.next = this.tail;
    this.tail.prev = this.head;
  }

  _remove(node) {
    node.prev.next = node.next;
    node.next.prev = node.prev;
  }

  _addToHead(node) {
    node.next = this.head.next;
    node.prev = this.head;
    this.head.next.prev = node;
    this.head.next = node;
  }

  get(key) {
    if (!this.cache.has(key)) return -1;
    const node = this.cache.get(key);
    this._remove(node);
    this._addToHead(node);
    return node.value;
  }

  put(key, value) {
    if (this.cache.has(key)) {
      const node = this.cache.get(key);
      node.value = value;
      this._remove(node);
      this._addToHead(node);
    } else {
      const newNode = new Node(key, value);
      if (this.cache.size >= this.capacity) {
        const last = this.tail.prev;
        this._remove(last);
        this.cache.delete(last.key);
      }
      this._addToHead(newNode);
      this.cache.set(key, newNode);
    }
  }
}
```


## 🚀 Practical Problems
```go
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

<!-- QUIZ_START 
[
    {
        "question": "What is the core principle of the LRU (Least Recently Used) cache?",
        "options": ["Keep everything forever", "Discard the items that haven't been accessed for the longest time", "Discard items randomly to save space", "Only store items that are small in size"],
        "correctIndex": 1
    },
    {
        "question": "Which two data structures are typically combined to implement an efficient LRU cache?",
        "options": ["Stack and Queue", "Hash Table and Doubly Linked List", "Binary Search Tree and Array", "Two Hash Tables"],
        "correctIndex": 1
    },
    {
        "question": "What is the time complexity for 'Get' and 'Put' operations in a well-implemented LRU cache?",
        "options": ["O(n)", "O(log n)", "O(1)", "O(n²)"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

