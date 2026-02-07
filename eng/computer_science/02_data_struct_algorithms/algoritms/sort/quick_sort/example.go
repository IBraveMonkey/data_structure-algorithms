package quick_sort

import "fmt"

func Example() {
	arr := []int{12, 7, 14, 9, 10, 11}
	fmt.Printf("Original: %v\n", arr)

	sorted := QuickSort(arr)
	fmt.Printf("Sorted: %v\n", sorted)

	// Example with duplicates
	arr2 := []int{5, 1, 9, 1, 5, 2}
	fmt.Printf("Original (with duplicates): %v\n", arr2)
	fmt.Printf("Sorted: %v\n", QuickSort(arr2))
}
