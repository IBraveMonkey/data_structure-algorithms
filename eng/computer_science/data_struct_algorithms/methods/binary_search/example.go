package binary_search

import "fmt"

// Example demonstrates the use of binary search
func Example() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	idx := BinarySearch(arr, target)
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Searching for number %d: found at index %d\n", target, idx)

	// Task: Search Insert Position
	nums := []int{1, 3, 5, 6}
	val := 5
	pos := SearchInsert(nums, val)
	fmt.Printf("Where to insert %d in %v? Index %d\n", val, nums, pos)

	val = 2
	pos = SearchInsert(nums, val)
	fmt.Printf("Where to insert %d in %v? Index %d\n", val, nums, pos)
}

// Task: Search Insert Position
// Given a sorted array and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// Example: [1,3,5,6], 5 -> 2
// Example: [1,3,5,6], 2 -> 1
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// Task: First Bad Version
// Imagine you are a product manager. Releasing a bad version means all subsequent versions are also bad.
// You need to find the first bad version with the minimum number of API checks isBadVersion(version).
func FirstBadVersion(n int, isBadVersion func(int) bool) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid // Bad version could be this or to the left
		} else {
			left = mid + 1 // This version is good, so the bad one is definitely to the right
		}
	}

	return left
}
