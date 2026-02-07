/*
Heap

What is it?
A Heap is a specialized tree-based data structure that satisfies the heap property: in a max-heap, for any given node C, if P is a parent node of C, then the value of P is greater than or equal to the value of C. In a min-heap, the value of P is less than or equal to the value of C. It is used for fast access to the minimum or maximum element.

Why is it needed?
- Fast retrieval of the minimum or maximum (O(1)).
- Efficient implementation of priority queues.
- Sorting data (Heap Sort).

What's the point?
- The tree is always balanced (a complete binary tree).
- The root always contains the extreme element (min or max).

When to use?
- When you need to frequently retrieve the minimum/maximum element.
- In graph algorithms (Dijkstra, Prim).
- In task schedulers (selecting the task with the highest priority).
- For finding the K largest/smallest elements in a data stream.

How does it work?
- Insertion: The element is added to the end of the tree and "sifts up" to its correct position.
- Removal: The root is replaced by the last element, which then "sifts down".
- Heapify: Converting an array into a heap in O(n) time.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Insertion | O(log n) | O(1) |
| Root removal (min/max) | O(log n) | O(1) |
| Get min/max | O(1) | O(1) |
| Build heap (Heapify) | O(n) | O(1) |
| Storage | — | O(n) |

How to know if a problem fits Heap?
- You need to find the "K largest/smallest" elements.
- You need to constantly retrieve the current minimum/maximum from a dynamically changing data set.
- The task involves merging K sorted arrays/lists.
*/

package heap

// IntHeap is our array of integers that will become a heap.
// It implements the heap.Interface from the standard container/heap library.
type IntHeap []int

// Len returns the number of elements.
func (h IntHeap) Len() int { return len(h) }

// Less compares elements (for min-heap: i < j)
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap swaps the elements
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the end.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the last element (the root is already extracted by the standard library before this call).
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
