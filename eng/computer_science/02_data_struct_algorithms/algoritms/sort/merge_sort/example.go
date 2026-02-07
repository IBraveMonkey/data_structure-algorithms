package merge_sort

import "fmt"

func Example() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("Original: %v\n", arr)

	sorted := MergeSort(arr)
	fmt.Printf("Sorted:   %v\n", sorted)

	// Problem: Sorting a large array (simulation)
	// In Merge Sort, it's often useful to see the merging stages, but this is just a demonstration
}
