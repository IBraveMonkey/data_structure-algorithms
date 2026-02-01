package insertion_sort

import "fmt"

func Example() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Printf("Original: %v\n", arr)

	InsertionSort(arr)
	fmt.Printf("Sorted:   %v\n", arr)
}
