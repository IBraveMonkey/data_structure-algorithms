/*
Queue

What is it?
A Queue is a linear data structure that follows the FIFO (First In, First Out) principle — the first one in is the first one out. Elements are added to the end of the queue and removed from the beginning of the queue. This is similar to a regular queue in a store: the first person in line is served first.

Why is it needed?
- Managing tasks that should be executed in the order of their arrival.
- Graph traversal algorithms (BFS).
- Data buffering between processes.
- Handling requests in web servers.

What's the point?
- Adding elements to the end (enqueue).
- Removing elements from the beginning (dequeue).
- Maintaining the order in which elements arrive.

When to use?
- When elements need to be processed in the order of their receipt.
- For algorithms requiring breadth-first search (BFS).
- When simulating queuing systems.

How does it work?
- Dequeueing an element happens from the beginning of the queue.
- Accessing the first element (peek) without removing it.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Enqueue (insertion) | O(1) | O(1) |
| Dequeue (removal) | O(1) | O(1) |
| Peek (view) | O(1) | O(1) |
| Search | O(n) | O(1) |
| Storage | — | O(n) |

How to know if a problem fits Queue?
- You need to process elements in the order of their receipt.
- A breadth-first traversal algorithm is required.
- There are tasks for simulating processes with a queue.

Queue in Go:

In Go, queues can be implemented in several ways:
- Using slices: a simple implementation, but less efficient for dequeueing.
- Using linked lists: a more efficient implementation.
- Using the container/list package: a standard doubly linked list implementation.

Examples of using Queue:
See the example.go file.
*/

package queue

import "fmt"

// ArrayQueue - a queue implemented using a slice
type ArrayQueue struct {
	Data []interface{}
	Size int
}

// Push - adds an element to the end of the queue
func (queue *ArrayQueue) Push(data interface{}) {
	queue.Data = append(queue.Data, data)
	queue.Size++
}

// Pop - removes and returns an element from the beginning of the queue
func (queue *ArrayQueue) Pop() (interface{}, bool) {
	if queue.Size <= 0 {
		return nil, false
	}

	firstElem := queue.Data[0]
	queue.Data = queue.Data[1:]
	queue.Size--

	return firstElem, true
}

// Peek - returns the first element without removal
func (queue *ArrayQueue) Peek() (interface{}, bool) {
	if queue.Size <= 0 {
		return nil, false
	}

	return queue.Data[0], true
}

// IsEmpty - checks if the queue is empty
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.Size == 0
}

// QueueNode - a node for a linked implementation of a queue
type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// LinkedListQueue - a queue implemented using a linked list
type LinkedListQueue struct {
	Head *QueueNode // Beginning of the queue
	Tail *QueueNode // End of the queue
	Size int
}

// Enqueue - adds an element to the end of the queue
func (q *LinkedListQueue) Enqueue(value interface{}) {
	newNode := &QueueNode{
		Value: value,
		Next:  nil,
	}

	// If the queue is empty, the new node becomes both the head and the tail
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		// Otherwise, add the new node to the end (after the tail)
		q.Tail.Next = newNode
		q.Tail = newNode
	}

	q.Size++
}

// Dequeue - removes and returns an element from the beginning of the queue
func (q *LinkedListQueue) Dequeue() (interface{}, bool) {
	if q.Head == nil {
		return nil, false
	}

	value := q.Head.Value
	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	q.Size--

	return value, true
}

// Front - returns the first element without removal
func (q *LinkedListQueue) Front() (interface{}, bool) {
	if q.Head == nil {
		return nil, false
	}

	return q.Head.Value, true
}

// IsEmpty - checks if the queue is empty
func (q *LinkedListQueue) IsEmpty() bool {
	return q.Head == nil
}

// Print - prints the contents of the queue
func (q *LinkedListQueue) Print() {
	current := q.Head

	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// CircularQueue - a circular queue (Ring Buffer)
type CircularQueue struct {
	data     []int
	front    int
	rear     int
	size     int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:     make([]int, capacity),
		front:    -1,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

func (cq *CircularQueue) Enqueue(value int) bool {
	if cq.size == cq.capacity {
		return false // Queue is full
	}

	if cq.front == -1 { // First insertion
		cq.front = 0
		cq.rear = 0
	} else {
		cq.rear = (cq.rear + 1) % cq.capacity
	}

	cq.data[cq.rear] = value
	cq.size++
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.size == 0 {
		return 0, false // Queue is empty
	}

	value := cq.data[cq.front]

	if cq.front == cq.rear { // Last element
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.capacity
	}

	cq.size--
	return value, true
}

func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}
