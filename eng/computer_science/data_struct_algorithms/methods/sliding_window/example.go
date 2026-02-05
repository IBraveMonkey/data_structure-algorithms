package sliding_window

import (
	"fmt"
	"math"
)

// Example demonstrates sliding window problems
func Example() {
	// 1. Maximum average in a subarray of length k
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("Max average (k=%d): %.2f\n", k, FindMaxAverage(nums, k))

	// 2. Minimum length of a subarray with sum >= target
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Printf("Min subarray length (sum >= %d): %d\n", target, MinSubArrayLen(target, arr))
}

// Task 1: Find the maximum average value of a subarray of length k
// O(N) time, O(1) space
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	// 1. Initialize the sum of the first window
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	// 2. Slide the window (starting from the k-th element)
	for i := k; i < len(nums); i++ {
		sum += nums[i]   // Add new element (on the right)
		sum -= nums[i-k] // Remove old element (on the left)
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

// Task 2: Minimum subarray size (Variable Size Sliding Window)
// Find the minimum length of a contiguous subarray whose sum is >= target.
// If no such subarray exists, return 0.
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // Expand window to the right

		// Shrink window from the left while the condition is met
		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
