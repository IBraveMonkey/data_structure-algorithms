package bubble_sort

import "fmt"

// Example demonstrates the use of bubble sort with various examples
func Example() {
	// Example 1: Basic bubble sort
	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original array: %v\n", arr1)
	BubbleSort(arr1)
	fmt.Printf("Sorted array: %v\n", arr1)

	// Example 2: Already sorted array
	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Already sorted array: %v\n", arr2)
	BubbleSort(arr2)
	fmt.Printf("After sorting: %v\n", arr2)

	// Example 3: Reverse sorted array
	arr3 := []int{5, 4, 3, 2, 1}
	fmt.Printf("Reverse sorted array: %v\n", arr3)
	BubbleSort(arr3)
	fmt.Printf("After sorting: %v\n", arr3)
}

// Problem: Sort an integer array in ascending order
// This is a basic problem for bubble sort
func SortArray(nums []int) []int {
	// Create a copy to avoid modifying the original
	result := make([]int, len(nums))
	copy(result, nums)

	// Apply bubble sort
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(result)-1; i++ {
			if result[i] > result[i+1] {
				result[i], result[i+1] = result[i+1], result[i]
				sorted = false
			}
		}
	}

	return result
}

// Problem: Sort array with the minimum number of swaps
// Although bubble sort is not optimal for this, it demonstrates the concept
func MinSwapsToSort(nums []int) int {
	// Create a copy
	arr := make([]int, len(nums))
	copy(arr, nums)

	swaps := 0
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
				sorted = false
			}
		}
	}

	return swaps
}
