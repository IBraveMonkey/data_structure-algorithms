package insertion_sort

/*
Insertion Sort

What is it?
Insertion Sort is a simple algorithm that builds a sorted array one element at a time.
It works much like the way you sort playing cards in your hands: you take a new card and insert it into its correct position among the already sorted cards.

Why is it needed?
- To sort small arrays.
- To finish sorting nearly sorted arrays (works in O(n) in the best case).
- Used as a component of more complex algorithms (e.g., TimSort uses Insertion Sort for small subarrays).

What's the core idea?
- We divide the array into "sorted" (left) and "unsorted" (right) parts.
- We take an element from the right part and shift the elements of the left part until we find its correct position.

When to use?
- When the array is small (N < 50).
- When the array is already partially sorted (adaptivity).
- When a simple, stable, in-place sort is needed without allocating extra memory.

How does it work?
1. Start with the second element (index 1).
2. Compare it with the elements to its left.
3. Shift all larger elements to the right.
4. Insert the current element into the newly freed position.

### Complexity

| Metric | Best (O) | Average/Worst (O) | Space (O) |
|:---|:---:|:---:|:---:|
| Time | O(n) | O(n²) | O(1) |

*O(n) is achieved on nearly sorted data.
*/

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Shift elements of arr[0..i-1] that are greater than key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
