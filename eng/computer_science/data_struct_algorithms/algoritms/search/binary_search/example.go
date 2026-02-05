package binary_search

import "fmt"

// Example demonstrates the use of binary search with various examples
func Example() {
	// Example 1: Basic binary search
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := BinarySearch(arr, target)
	fmt.Printf("Binary search for number %d in %v: index %d\n", target, arr, result)

	// Example 2: Binary search for a non-existent element
	target = 2
	result = BinarySearch(arr, target)
	fmt.Printf("Binary search for number %d in %v: index %d\n", target, arr, result)

	// Example 3: Binary search for square root
	fmt.Printf("Square root of 9: %d\n", binarySearchSqrt(9))
	fmt.Printf("Square root of 21: %d\n", binarySearchSqrt(21))
}

// Problem: Find Peak Element
// A peak element is an element that is strictly greater than its neighbors.
// Given a 0-indexed integer array nums, find a peak element and return its index.
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// The peak is in the left part (including mid)
			right = mid
		} else {
			// The peak is in the right part
			left = mid + 1
		}
	}

	return left
}

// Problem: Search in Rotated Sorted Array
// Given an array nums after rotation and an integer target, return the index of target if it is in nums, or -1 if it is not.
func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Check if the left half is sorted
		if nums[left] <= nums[mid] {
			// Target is in the left half
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				// Target is in the right half
				left = mid + 1
			}
		} else {
			// The right half is sorted
			if nums[mid] < target && target <= nums[right] {
				// Target is in the right half
				left = mid + 1
			} else {
				// Target is in the left half
				right = mid - 1
			}
		}
	}

	return -1
}

// Problem: Find Minimum in Rotated Sorted Array
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
func FindMinInRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// The minimum is in the right part
			left = mid + 1
		} else {
			// The minimum is in the left part (including mid)
			right = mid
		}
	}

	return nums[left]
}
