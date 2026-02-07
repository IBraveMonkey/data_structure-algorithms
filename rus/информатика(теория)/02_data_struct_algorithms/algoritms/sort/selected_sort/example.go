package selected_sort

import "fmt"

func Example() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Printf("Original: %v\n", arr)

	SelectedSort(arr)
	fmt.Printf("Sorted:   %v\n", arr)
}
