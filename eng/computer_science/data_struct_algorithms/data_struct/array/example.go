package array

import "fmt"

// Examples shows practical techniques for working with arrays
func Examples() {
	// Array Reverse (In-place)
	// Complexity: O(n) time, O(1) space
	data := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	fmt.Println("Reverse:", data)

	// Finding the maximum
	// Complexity: O(n)
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	fmt.Println("Maximum:", max)

	// Filtering (creating a new slice)
	// Complexity: O(n)
	even := make([]int, 0)
	for _, v := range data {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	fmt.Println("Evens:", even)
}

/*
Important points about Slices in Go:
1. len(s) - number of elements in the slice.
2. cap(s) - capacity of the underlying array.
3. A slice is a descriptor (pointer to array, length, capacity).
   Passing a slice to a function copies the descriptor, but not the data themselves.
*/
