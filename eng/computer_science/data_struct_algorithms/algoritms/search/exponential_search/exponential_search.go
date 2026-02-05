package exponential_search

import (
	"math"
)

/*
Exponential Search

What is it?
It is a search algorithm for finding an element in a sorted, infinite, or very large array where the size is unknown or too large.
It works in two stages: first, it finds the range where the element might be (by doubling the index), and then it runs a binary search within that range.

Why is it needed?
- For searching in unbounded or streaming arrays.
- It works faster than binary search if the desired element is close to the beginning of the array (O(log i) vs O(log n)).

What's the core idea?
- We "jump" through the array with indices 1, 2, 4, 8, 16... as long as the current element is less than the desired one.
- Once we've jumped past it (arr[i] > target), we know the element is somewhere between [i/2, i].

When to use?
- The array is sorted, but the size is huge or unknown.
- There is a high probability that the elements are at the beginning.

How does it work?
1. Check the 0-th element.
2. Initialize bound = 1.
3. While bound < len(arr) and arr[bound] < target:
   bound *= 2.
4. Run Binary Search in the range [bound/2, min(bound, len)].

Complexity:
- Time: O(log i), where i is the index of the desired element.
- Space: O(1).
*/

func ExponentialSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// 1. Check the first element
	if arr[0] == target {
		return 0
	}

	// 2. Find the range
	bound := 1
	for bound < n && arr[bound] <= target {
		bound *= 2
	}

	// 3. Binary Search
	// Range: [bound/2, min(bound, n-1)]
	left := bound / 2
	right := int(math.Min(float64(bound), float64(n-1)))

	return binarySearch(arr, left, right, target)
}

func binarySearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
