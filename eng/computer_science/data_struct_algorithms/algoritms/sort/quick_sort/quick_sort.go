package quick_sort

/*
Quick Sort

What is it?
Quick Sort is an efficient sorting algorithm based on the "Divide and Conquer" principle. It selects a "pivot" element and reorders the array so that all elements smaller than the pivot are on the left, and all larger ones are on the right. It then recursively sorts these two parts.

Why is it needed?
- To sort arrays as quickly as possible on average.
- The de facto standard for in-place sorting (without significant extra memory if implemented correctly).

What's the core idea?
- Partitioning is the key stage. We "distribute" elements relative to the pivot.
- After partitioning, the pivot is in its final position in the sorted array.

When to use?
- When a fast, general-purpose sort is needed (like in standard libraries).
- When speed in the average case O(n log n) is important.
- When stability is not required (the relative order of equal elements may change).

How does it work?
1. Select a pivot (first, last, middle, or median-of-three).
2. Iterate through the array and move elements: smaller than pivot -> left, larger -> right.
3. Recursively run Quick Sort on the left and right sub-parts.

### Complexity

| Metric | Best/Average (O) | Worst (O) | Space (O) |
|:---|:---:|:---:|:---:|
| Time | O(n log n) | O(n²) | O(log n) / O(n) |
| Stability | ❌ (Unstable) | — | — |

*The worst case O(n²) occurs with a poor choice of the pivot element.
**Space O(log n) is for the call stack in an optimal in-place implementation.
*/

func medianOfThree(arr []int, low, high int) int {
	mid := low + (high-low)/2

	// Compare three elements: first, middle, and last to select the median
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	return mid
}

// QuickSort - sorting implementation.
// Note: This implementation creates new slices (not in-place) for conceptual simplicity.
// For production code, an in-place version with a partition function is typically used.
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := medianOfThree(arr, 0, len(arr)-1)
	pivot := arr[pivotIndex]

	left := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)
	pivots := make([]int, 0) // For handling pivot duplicates

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			pivots = append(pivots, v)
		}
	}

	return append(append(QuickSort(left), pivots...), QuickSort(right)...)
}
