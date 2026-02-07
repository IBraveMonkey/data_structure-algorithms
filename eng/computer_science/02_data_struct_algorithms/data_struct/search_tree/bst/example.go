package bst

import "fmt"

// Example demonstrates the use of a binary search tree (BST)
func Example() {
	bst := NewBST()

	// Insert values
	values := []int{5, 3, 8, 2, 4, 7, 9}
	for _, v := range values {
		bst.Insert(v)
	}

	fmt.Println("In-order traversal (sorted):", bst.InOrderTraversal())
	fmt.Println("BFS traversal (level-order):", bst.BFS())
	fmt.Println("Tree height:", bst.Height())
	fmt.Println("Is valid BST:", bst.IsValidBST())

	// Search
	fmt.Println("Search for 4:", bst.Search(4))
	fmt.Println("Search for 10:", bst.Search(10))

	// Delete
	bst.Delete(3)
	fmt.Printf("After deleting 3 (In-order): %v\n", bst.InOrderTraversal())

	// Range sum
	fmt.Printf("Sum in range [4, 9]: %d\n", bst.RangeSum(4, 9))

	// K-th smallest element
	k := 2
	fmt.Printf("%d-th smallest element: %d\n", k, bst.KthSmallest(k))
}
