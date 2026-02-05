package red_black_tree

import "fmt"

// Example demonstrates the use of a red-black tree
func Example() {
	rbt := &RBTree{}

	fmt.Println("Inserting elements: 10, 20, 30, 15")
	rbt.Insert(10)
	rbt.Insert(20)
	rbt.Insert(30)
	rbt.Insert(15)

	fmt.Println("Tree structure (Val Color):")
	// Note: full balancing is not yet implemented in red_black_tree.go
	if rbt.Root != nil {
		rbt.PrintTree(rbt.Root, "")
	} else {
		fmt.Println("Tree is empty")
	}
}
