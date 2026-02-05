package array

/*
Array and Slice

What is it?
An Array is a fixed-size data structure containing elements of the same type, arranged in a contiguous memory area.
In Go, Slices are more commonly used — these are dynamic arrays that can change their size.

Why is it needed?
- For storing ordered lists of data of the same type.
- When fast access to elements by index is required.
- To provide high performance due to:
  - Direct access to elements.
  - Insertion and deletion of elements.
  - Dynamic expansion (via append).

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Access (by index) | O(1) | O(1) |
| Insertion (at end) | O(1) amortized | O(1) |
| Insertion (in middle) | O(n) | O(1) |
| Deletion (from end) | O(1) | O(1) |
| Deletion (from middle) | O(n) | O(1) |
| Search (unsorted) | O(n) | O(1) |
| Storage | — | O(n) |

What's the point?
- Understanding how slices work in Go.
- Efficient memory usage.
- Choosing correctly between an array and a slice.
- Cache-friendly: predictable data access patterns.
- Direct access formula: address = start + index * element_size.

When to use?
- When the number of elements is known in advance or changes infrequently.
- When read/write speed by index is critical (O(1)).
- When data must be processed sequentially.

Operation Complexity:
- Access by index: O(1)
- Search (unsorted): O(n)
- Search (sorted, binary): O(log n)
- Insertion (at start/middle): O(n) — requires shifting elements.
- Insertion (at end): O(1) (amortized for slices).
- Deletion: O(n) — requires shifting elements.

How to know if a problem fits Array/Slice?
- Data is a simple sequence.
- Frequent access by index is required.
- Memory must be allocated as a single block for efficiency.
*/

import "fmt"

// BasicOperations demonstrates basic operations with arrays and slices
func BasicOperations() {
	// 1. Array declaration (fixed size)
	var arr [5]int
	arr[0] = 10 // O(1)

	// 2. Slice declaration (dynamic array)
	slice := []int{1, 2, 3}

	// 3. Appending to the end (Append)
	// Average O(1), but if capacity is exhausted,
	// reallocation and copying of the array will occur (O(n)).
	slice = append(slice, 4)

	// 4. Insertion in the middle (O(n))
	index := 1
	value := 99
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value

	// 5. Deletion from the middle (O(n))
	// Preserving order by shifting elements
	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println("Final slice:", slice)
}
