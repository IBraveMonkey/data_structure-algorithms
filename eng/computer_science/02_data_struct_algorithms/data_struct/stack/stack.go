/*
Stack

What is it?
A Stack is an ordered collection of elements that operates on the LIFO principle (Last In, First Out).
Imagine a stack of plates: you can only place a new plate on top and take a plate only from the top.

Why is it needed?
- For tracking the order of operations (e.g., function calls, recursion).
- For undoing actions (Undo).
- For parsing expressions and syntax checking (e.g., parentheses).

What's the point?
- Access is provided only to one element — the top one.
- Strict order of addition and removal.

When to use?
- When elements need to be processed in the reverse order of their arrival.
- In traversal algorithms (DFS).
- For checking the balance of parentheses.

How does it work?
- Push: Add an element to the top of the stack.
- Pop: Remove and return the element from the top of the stack.
- Peek (Top): Look at the top element without removing it.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Push (insertion) | O(1) | O(1) |
| Pop (removal) | O(1) | O(1) |
| Peek (view) | O(1) | O(1) |
| Search | O(n) | O(1) |
| Storage | — | O(n) |

How to know if a problem fits Stack?
- You need to find a "pair" for an element (e.g., a closing parenthesis).
- You need to process data in reverse order.
- The problem involves nested structures.
*/

package stack

// ArrayStack - Stack implementation based on a slice (more performant in Go due to cache locality)
type ArrayStack struct {
	Data []interface{}
}

// Push - adds an element to the stack
func (stack *ArrayStack) Push(data interface{}) {
	stack.Data = append(stack.Data, data)
}

// Pop - removes and returns the top element of the stack
func (stack *ArrayStack) Pop() (interface{}, bool) {
	if len(stack.Data) == 0 {
		return nil, false
	}

	lastIndex := len(stack.Data) - 1
	lastElem := stack.Data[lastIndex]

	stack.Data = stack.Data[0:lastIndex]
	return lastElem, true
}

// Peek - returns the top element without removal
func (stack *ArrayStack) Peek() (interface{}, bool) {
	if len(stack.Data) == 0 {
		return nil, false
	}

	lastIndex := len(stack.Data) - 1
	return stack.Data[lastIndex], true
}

// IsEmpty - checks if the stack is empty
func (stack *ArrayStack) IsEmpty() bool {
	return len(stack.Data) == 0
}

// Size - returns the size of the stack
func (stack *ArrayStack) Size() int {
	return len(stack.Data)
}

// Node - node for linked implementation of the stack
type Node struct {
	Value interface{}
	Next  *Node
}

// Stack - Stack implementation based on a linked list (less memory efficient due to pointers)
type Stack struct {
	Top     *Node
	SizeVal int
}

// Push - adds an element to the stack
func (s *Stack) Push(value interface{}) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
	s.SizeVal++
}

// Pop - removes and returns the top element of the stack
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	removedValue := s.Top.Value
	s.Top = s.Top.Next
	s.SizeVal--
	return removedValue, true
}

// Peek - returns the top element without removal
func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s.Top.Value, true
}

// IsEmpty - checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// Size - returns the size of the stack
func (s *Stack) Size() int {
	return s.SizeVal
}
