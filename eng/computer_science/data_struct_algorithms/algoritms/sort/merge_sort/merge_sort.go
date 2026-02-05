package merge_sort

/*
Merge Sort

What is it?
Merge Sort is a sorting algorithm based on the "Divide and Conquer" principle. It recursively divides the array in half until pieces of length 1 remain (which are already sorted), and then it "merges" these pieces back together in the correct order.

Why is it needed?
- Guaranteed time complexity of O(n log n) even in the worst case (unlike Quick Sort).
- Stable sort (preserves the order of equal elements).
- Well-suited for sorting linked lists or data that does not fit into memory (External Sorting).

What's the core idea?
- Merging two sorted arrays into one is very simple and fast (at O(n)).
- The entire complexity of the algorithm lies in this merge operation.

When to use?
- When guaranteed O(n log n) performance is required.
- When stability is important.
- When working with Linked Lists (as it doesn't require random access).

How does it work?
1. If the array length is 0 or 1, return it.
2. Divide the array into Left and Right halves at the midpoint.
3. Recursively call MergeSort(Left) and MergeSort(Right).
4. Merge the results using the Merge function.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time (Always) | O(n log n) |
| Space | O(n) |
| Stability | ✅ (Stable) |

*O(n) space is required for a temporary buffer during the merge.
*/

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}
