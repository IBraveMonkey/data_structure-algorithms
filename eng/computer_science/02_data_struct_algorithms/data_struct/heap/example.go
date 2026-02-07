package heap

import (
	"container/heap"
	"fmt"
)

// Example demonstrates the use of a Heap
func Example() {
	// Task: Find the K-th largest element in an array
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := FindKthLargest(nums, k)
	fmt.Printf("For array %v, the %d-th largest element is: %d\n", nums, k, result)

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	result2 := FindKthLargest(nums2, k2)
	fmt.Printf("For array %v, the %d-th largest element is: %d\n", nums2, k2, result2)
}

// Problem 1: K-th Largest Element in an Array
// Given an integer array nums and an integer k, return the k-th largest element in the array.
// Note that it is the k-th largest element in sorted order, not the k-th unique element.
//
// Example: [3,2,1,5,6,4], k = 2 => 5
// Complexity: O(N log K)
func FindKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		// If the heap size exceeds k, remove the minimum element (root)
		// Since IntHeap is a min-heap, we keep the K largest elements.
		// The root will be the smallest of these K, which is the K-th largest overall.
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Return the root of the heap, which is the K-th largest element
	// (being the minimum among the K largest).
	return (*h)[0]
}
