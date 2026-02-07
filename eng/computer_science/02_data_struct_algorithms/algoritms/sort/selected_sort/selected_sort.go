package selected_sort

/*
Selection Sort

What is it?
Selection Sort is an algorithm that divides an array into sorted and unsorted parts.
It repeatedly finds the minimum element from the unsorted part and places it at the end of the sorted part.

Why is it needed?
- It is one of the simplest algorithms to understand.
- It performs the minimum number of writes (swaps) — exactly N (or N-1), which can be useful if writing to memory is very expensive (e.g., flash memory).

What's the core idea?
- "Find the smallest, put it first. Find the next smallest, put it second..."

When to use?
- For educational purposes.
- When the number of write operations is critical (though Cycle Sort performs even fewer, Selection Sort achieves O(N)).
- In practice, it almost always loses to Insertion Sort or Quick Sort.

How it works?
1. In a loop from 0 to N-1:
   - Find the index of the minimum element in the subarray [i...N-1].
   - Perform swap(arr[i], arr[minIdx]).

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time (Always) | O(n²) |
| Space | O(1) |
| Stability | ❌ (Unstable) |
*/

func SelectedSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		// Look for the minimum in the remaining part
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap if a new minimum was found
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
