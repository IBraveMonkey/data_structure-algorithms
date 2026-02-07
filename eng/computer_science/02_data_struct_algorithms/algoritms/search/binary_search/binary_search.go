package binary_search

/*
Binary Search is an algorithm for finding an element within a sorted array. Its essence is that at each step, the array is divided in half, and it is checked which half the desired element is in. This allows reducing the number of checks.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(log n) |
| Space | O(1) iterative / O(log n) recursive |

Pros: High speed on large datasets.
Cons: Requires a pre-sorted array.
*/

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	if target < arr[left] || target > arr[right] {
		return -1
	}

	for left <= right {
		mid := (right + left) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			right = mid + 1
		} else {
			left = mid - 1
		}
	}

	return -1
}

/*
Square Root Calculation - write a function that finds the square root of a number or the nearest smaller integer

For example - for 9 it will be 3
For 21 will be 4
5 won't work because 5 squared is 25, which is greater than 21
*/
func binarySearchSqrt(target int) int {
	left := 0
	right := target

	for left <= right {
		middle := (left + right) / 2

		if middle*middle > target {
			right = middle - 1
			continue
		}

		if middle*middle < target {
			left = middle + 1
			continue
		}

		return middle
	}

	return right
}
