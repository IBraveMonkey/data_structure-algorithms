package exponential_search

import "fmt"

func Example() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 20}
	target := 12

	idx := ExponentialSearch(arr, target)

	fmt.Printf("Array: ... (huge sorted)\n")
	if idx != -1 {
		fmt.Printf("Element %d found at index %d\n", target, idx)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
