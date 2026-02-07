package two_pointers

import (
	"fmt"
	"strings"
)

// Example demonstrates two pointers problems
func Example() {
	// 1. Two Sum (for sorted array)
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("TwoSum indices for %v, target %d: %v\n", nums, target, TwoSumSorted(nums, target))

	// 2. Palindrome check
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("IsPalindrome ('%s'): %v\n", s, IsPalindrome(s))

	// 3. Duplicate removal
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	lenAfter := RemoveDuplicates(arr)
	fmt.Printf("RemoveDuplicates - new len: %d, arr prefix: %v\n", lenAfter, arr[:lenAfter])
}

// Task 1: Two Sum (Input Array Is Sorted)
// Find two numbers that add up to a target. Return their indices (0-based in this implementation).
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++ // Sum is too small, we need more -> move left pointer to the right
		} else {
			right-- // Sum is too large, we need less -> move right pointer to the left
		}
	}

	return []int{}
}

// Task 2: Palindrome Validation
func IsPalindrome(s string) bool {
	// In a real task, it's better to handle runes and not create a new string for O(1) space,
	// but we'll simplify string preparation for this example.
	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r + 32 // toLower
		}
		return -1
	}, s)

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// Task 3: Removing duplicates from a sorted array
// Returns the new length (k).
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// If a new unique element is found
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// Task 4: In-Place Array Reversal
func ReverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
