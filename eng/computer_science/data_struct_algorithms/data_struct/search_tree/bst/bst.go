/*
Binary Search Tree (BST)

What is it?
A BST is a special tree data structure where each node has at most two children: Left and Right. For each node, all values in the left subtree are less than the node's value, and all values in the right subtree are greater. This property allows for efficient searching, insertion, and deletion.

Why is it needed?
- Fast search, insertion, and deletion of elements (O(log n) in a balanced tree).
- Maintaining ordered data.
- Implementation of associative arrays and sets.

What's the point?
- Each node has at most two children.
- Left subtree contains only values smaller than the parent.
- Right subtree contains only values larger than the parent.

When to use?
- When fast access to ordered data is needed.
- For implementing dictionaries and sets.
- When search operations are critical.

How does it work?
- Search: Similar to insertion, but returns the found node.
- Deletion: More complex, especially for nodes with two children.

### Complexity

| Operation | Average (O) | Worst (O)* | Space Complexity (O) |
|:---|:---:|:---:|:---:|
| Insertion | O(log n) | O(n) | O(h) |
| Search | O(log n) | O(n) | O(h) |
| Deletion | O(log n) | O(n) | O(h) |
| Traversal | O(n) | O(n) | O(h) |
| Storage | — | — | O(n) |

*The worst case O(n) occurs when the tree degenerates into a linear list.
**h — tree height. Space complexity is due to the recursion stack.

How to know if a problem fits BST?
- Fast search in ordered data is needed.
- Maintaining order is required.
- Frequent insertion, search, and deletion operations.
*/

package bst

import (
	"container/list"
)

// Node - binary search tree node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST - binary search tree (wrapper around nodes)
type BST struct {
	Root *Node
	Size int
}

// NewBST - creates a new empty binary search tree
func NewBST() *BST {
	return &BST{}
}

// Insert - inserts a value into the tree
func (bst *BST) Insert(value int) {
	bst.Root = insertRecursive(bst.Root, value)
	bst.Size++
}

func insertRecursive(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}
	// If value == node.Value, ignore duplicate

	return node
}

// Search - searches for a value in the tree
func (bst *BST) Search(value int) bool {
	return searchRecursive(bst.Root, value)
}

func searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return searchRecursive(node.Left, value)
	}

	return searchRecursive(node.Right, value)
}

// Delete - deletes a value from the tree
func (bst *BST) Delete(value int) {
	bst.Root = deleteRecursive(bst.Root, value)
	bst.Size--
}

func deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// Found the node to delete
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node with two children: find minimum value in the right subtree
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteRecursive(node.Right, minNode.Value)
	}

	return node
}

// findMin - finds the node with the minimum value
func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InOrderTraversal - returns nodes in sorted order (left -> root -> right)
func (bst *BST) InOrderTraversal() []int {
	result := []int{}
	inOrderHelper(bst.Root, &result)
	return result
}

func inOrderHelper(node *Node, result *[]int) {
	if node != nil {
		inOrderHelper(node.Left, result)
		*result = append(*result, node.Value)
		inOrderHelper(node.Right, result)
	}
}

// BFS - breadth-first traversal (level-order)
func (bst *BST) BFS() []int {
	if bst.Root == nil {
		return []int{}
	}

	result := []int{}
	queue := list.New()
	queue.PushBack(bst.Root)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*Node)
		queue.Remove(element)

		result = append(result, node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return result
}

// Height - calculates the tree height
func (bst *BST) Height() int {
	return heightRecursive(bst.Root)
}

func heightRecursive(node *Node) int {
	if node == nil {
		return -1
	}

	leftHeight := heightRecursive(node.Left)
	rightHeight := heightRecursive(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsValidBST - checks if the tree is a valid BST
func (bst *BST) IsValidBST() bool {
	// Using pointers in helper to correctly handle int bounds
	return isValidBSTHelper(bst.Root, nil, nil)
}

func isValidBSTHelper(node *Node, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
		return false
	}

	val := node.Value
	return isValidBSTHelper(node.Left, min, &val) &&
		isValidBSTHelper(node.Right, &val, max)
}

// RangeSum - sum of values within a given range
func (bst *BST) RangeSum(low, high int) int {
	return rangeSumHelper(bst.Root, low, high)
}

func rangeSumHelper(node *Node, low, high int) int {
	if node == nil {
		return 0
	}

	sum := 0
	if node.Value >= low && node.Value <= high {
		sum += node.Value
	}

	if node.Value > low {
		sum += rangeSumHelper(node.Left, low, high)
	}

	if node.Value < high {
		sum += rangeSumHelper(node.Right, low, high)
	}

	return sum
}

// KthSmallest - finds the k-th smallest element (k is 1-indexed)
func (bst *BST) KthSmallest(k int) int {
	count := 0
	return kthSmallestHelper(bst.Root, &count, k)
}

func kthSmallestHelper(node *Node, count *int, k int) int {
	if node == nil {
		return -1
	}

	// Traverse left subtree
	leftResult := kthSmallestHelper(node.Left, count, k)
	if leftResult != -1 {
		return leftResult
	}

	// Process current node
	*count++
	if *count == k {
		return node.Value
	}

	// Traverse right subtree
	return kthSmallestHelper(node.Right, count, k)
}
