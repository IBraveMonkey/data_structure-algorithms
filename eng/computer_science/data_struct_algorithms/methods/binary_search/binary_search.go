package binary_search

/*
Binary Search

What is it?
Binary Search is an algorithm for finding an element in a SORTED array by repeatedly dividing the search interval in half.

Why is it needed?
- To find an element in an array as quickly as possible (O(log n) complexity).
- To find a boundary or conditions in a monotonic function.

What's the point?
- We always check the middle element.
- If the middle element is greater than the target, then all elements to the right are also greater (discard the right half).
- If the middle element is smaller than the target, discard the left half.

When to use?
- Data is sorted.
- Problem of finding an element, first occurrence, or last occurrence.
- Search for an answer in a range (binary search by answer).

How does it work?
1. Define the limits `left` and `right`.
2. While `left <= right`:
   - Find `mid`.
   - If `arr[mid] == target`, return the index.
   - If `arr[mid] > target`, search left (`right = mid - 1`).
   - Otherwise search right (`left = mid + 1`).
3. If the loop ends, the element is not found.

Complexity: O(log n)
*/

// BinarySearch - classic binary search
// Returns the index of the element or -1 if not found.
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// mid := (left + right) / 2 can cause integer overflow for very large left and right
		// Safe way to calculate the middle:
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			right = mid - 1 // Search in the left part
		} else {
			left = mid + 1 // Search in the right part
		}
	}

	return -1
}
