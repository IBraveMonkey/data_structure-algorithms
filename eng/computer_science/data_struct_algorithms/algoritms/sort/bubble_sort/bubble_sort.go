package bubble_sort

/*
Bubble Sort is a simple but inefficient sorting algorithm that repeatedly steps through the array, compares adjacent elements, and swaps them if they are in the wrong order. The algorithm gets its name because large elements gradually "bubble up" to the end of the array, much like air bubbles in water.

Algorithm Characteristics:
- Stability: Yes, the algorithm is stable, meaning it preserves the relative order of equal elements in the sorted array.

### Complexity

| Metric | Best (O) | Average/Worst (O) | Space (O) |
|:---|:---:|:---:|:---:|
| Time | O(n) | O(n²) | O(1) |

*O(n) is achieved if the array is already sorted (using an optimization flag).

- Applicability: Suitable for small arrays or educational purposes, but inefficient for large datasets compared to more complex algorithms (e.g., QuickSort).

*/

func BubbleSort(arr []int) {
	sorted := false

	for !sorted {
		sorted = true

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}
