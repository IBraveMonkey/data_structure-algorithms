package ternary

import "fmt"

func Example() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	// Using the iterative version as it is safer for the stack
	result := TernarySearch(data, target)

	if result != -1 {
		fmt.Printf("Element %d found at position %d (Ternary Search)\n", target, result)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
