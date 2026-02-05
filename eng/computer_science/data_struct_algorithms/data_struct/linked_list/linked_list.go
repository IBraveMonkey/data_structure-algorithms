/*
Linked List

What is it?
A Linked List is a linear data structure in which elements (nodes) are linked together sequentially. Each node contains data and a reference (pointer) to the next node in the list. Unlike arrays, list elements are not stored in a contiguous memory area.

Why is it needed?
- Efficient insertion and deletion of elements at the beginning, middle, or end of the list (O(1) if a pointer is available).
- Dynamic resizing (no need to know the size in advance).
- Memory efficiency during frequent structural changes.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Access (by index) | O(n) | O(1) |
| Insertion (at front) | O(1) | O(1) |
| Insertion (at back) | O(1)* | O(1) |
| Insertion (in middle) | O(n) | O(1) |
| Deletion (from front) | O(1) | O(1) |
| Deletion (from back) | O(n)** | O(1) |
| Deletion (from middle) | O(n) | O(1) |
| Search | O(n) | O(1) |
| Storage | — | O(n) |

*If a Tail pointer is present. Without it — O(n).
**In a singly linked list — O(n), as you need to find the second-to-last node. In a doubly linked list — O(1).

What's the point?
- Each node contains data and a pointer to the next node.
- No need for contiguous memory.
- Insertion/deletion operations do not require shifting other elements.

When to use?
- When elements are frequently inserted or deleted.
- When the size of the data structure changes often.
- When memory conservation during frequent changes is important.

How does it work?
- Each node contains data and a pointer to the next node.
- To access an element, you must traverse from the beginning to the desired position.
- To insert/delete, you need to modify the pointers of neighboring nodes.

How to know if a problem fits Linked List?
- You need to insert/delete elements frequently.
- Fast access to an arbitrary index is not required.
- Data size changes dynamically.

Linked List in Go:

In Go, linked lists are implemented using structs and pointers:
- A Node contains data and a pointer to the next node.
- A LinkedList contains pointers to the head and, possibly, the tail.
- Go automatically manages garbage collection, but you need to be careful with cyclic references.
- There is no built-in implementation; you need to implement it yourself.

Examples of using Linked List:
See the example.go file.
*/

package linked_list

import "fmt"

// Node represents a node in a linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
	Tail *Node
	Size int // Adding size for convenience
}

// AddToFront adds an element to the beginning of the list
func (l *LinkedList) AddToFront(value int) {
	newNode := &Node{
		Value: value,
		Next:  l.Head, // Point next to the current head
	}

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}

	l.Size++
}

// AddToBack adds an element to the end of the list
func (l *LinkedList) AddToBack(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	// If the list is empty, the new node becomes the head
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Add the new node after the current tail
		l.Tail.Next = newNode
		// Update the tail to point to the new node
		l.Tail = newNode
	}

	l.Size++
}

// InsertAt inserts an element at the specified position
func (l *LinkedList) InsertAt(value int, index int) error {
	if index < 0 || index > l.Size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		l.AddToFront(value)
		return nil
	}

	if index == l.Size {
		l.AddToBack(value)
		return nil
	}

	// Find the node before the insertion position
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode := &Node{
		Value: value,
		Next:  current.Next,
	}

	current.Next = newNode
	l.Size++

	return nil
}

// RemoveFromFront removes an element from the beginning of the list
func (l *LinkedList) RemoveFromFront() {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next

	if l.Head == nil {
		l.Tail = nil
	}

	l.Size--
}

// RemoveFromBack removes an element from the end of the list
func (l *LinkedList) RemoveFromBack() {
	if l.Head == nil {
		return
	}

	// If there is only one element in the list
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return
	}

	// Find the second-to-last element
	current := l.Head
	for current.Next != l.Tail {
		current = current.Next
	}

	// Remove the last element
	current.Next = nil
	l.Tail = current

	l.Size--
}

// Find finds a node with the given value
func (l *LinkedList) Find(value int) *Node {
	current := l.Head

	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

// Get returns the value of the node at the specified index
func (l *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= l.Size {
		return 0, fmt.Errorf("index out of bounds")
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}

// Print prints the list
func (l *LinkedList) Print() {
	current := l.Head

	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// Reverse reverses the list
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.Head
	l.Tail = l.Head // After reversing, the old head will become the tail

	for current != nil {
		nextTemp := current.Next // Save next node
		current.Next = prev      // Reverse the pointer direction
		prev = current           // Move prev forward
		current = nextTemp       // Move current forward
	}

	l.Head = prev // New head is the former last element
}

// RemoveValue removes the first occurrence of a value from the list
func (l *LinkedList) RemoveValue(value int) {
	if l.Head == nil {
		return
	}

	// If the element to be removed is the head
	if l.Head.Value == value {
		l.RemoveFromFront()
		return
	}

	// Look for the node before the one to be removed
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// If the element to be removed was found
	if current.Next != nil {
		// If the element to be removed is the tail, update top level tail
		if current.Next == l.Tail {
			l.Tail = current
		}

		current.Next = current.Next.Next
		l.Size--
	}
}
