package linked_list

import "fmt"

// Example demonstrates the use of a linked list with various examples
func Example() {
	// Create a new linked list
	list := &LinkedList{}

	// Add elements
	list.AddToBack(1)
	list.AddToBack(2)
	list.AddToBack(3)
	list.AddToBack(4)
	list.AddToBack(5)

	fmt.Printf("Original list: ")
	list.Print()

	// Find an element
	node := list.Find(3)
	if node != nil {
		fmt.Printf("Found node with value: %d\n", node.Value)
	} else {
		fmt.Println("Node not found")
	}

	// Geat element by index
	value, err := list.Get(2)
	if err == nil {
		fmt.Printf("Value at index 2: %d\n", value)
	}

	// Reverse the list
	list.Reverse()
	fmt.Printf("Reversed list: ")
	list.Print()

	// Find the middle
	middle := Middle(list)
	if middle != nil {
		fmt.Printf("Middle element: %d\n", middle.Value)
	}
}

// Problem 1: Cycle Detection (Floyd's Algorithm)
// Checks if there is a cycle in the list.
func HasCycle(l *LinkedList) bool {
	if l.Head == nil || l.Head.Next == nil {
		return false
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// Problem 2: Find the Middle Element (Two-Pointer Algorithm)
// Finds the middle element of the list.
func Middle(l *LinkedList) *Node {
	if l.Head == nil {
		return nil
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Problem 3: Remove N-th Node from End
// Given a list, remove the n-th node from the end and return the head.
func RemoveNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{Value: 0, Next: head}
	fast := dummy
	slow := dummy

	// Move fast n+1 steps forward
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the n-th node from the end
	slow.Next = slow.Next.Next

	return dummy.Next
}

// Problem 4: Merge Two Sorted Lists
// Merges two sorted linked lists into one sorted list.
func MergeTwoLists(list1 *Node, list2 *Node) *Node {
	dummy := &Node{}
	current := dummy

	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Value <= p2.Value {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}

	// Add the remaining nodes
	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}

	return dummy.Next
}

// Problem 5: Palindrome Linked List
// Check if a singly linked list is a palindrome.
func IsPalindromeList(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find the middle of the list
	slow, fast := head, head
	var prev *Node

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the list
	if prev != nil {
		prev.Next = nil // Splitting
	}

	// Reverse the second half (starting from slow)
	var secondHalf *Node
	curr := slow
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = secondHalf
		secondHalf = curr
		curr = nextTemp
	}

	// Compare the two halves
	firstHalf := head
	p2 := secondHalf

	for firstHalf != nil && p2 != nil {
		if firstHalf.Value != p2.Value {
			return false
		}
		firstHalf = firstHalf.Next
		p2 = p2.Next
	}

	return true
}
