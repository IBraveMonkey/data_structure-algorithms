package counting_sort

import "fmt"

func Example() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Printf("Original: %v\n", arr)

	// Counting Sort often returns a new array or modifies the current one
	// In our implementation, it modifies the current one in-place (pseudo) and returns it
	sorted := CountingSort(arr)
	fmt.Printf("Sorted:   %v\n", sorted)

	// Example with negative numbers (our implementation supports this via offset)
	arr2 := []int{-5, -10, 0, -3, 8, 5, -1, 10}
	fmt.Printf("Sorted (with negatives): %v\n", CountingSort(arr2))
}
