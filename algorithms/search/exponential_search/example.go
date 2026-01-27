package exponential_search

import "fmt"

func Example() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 20}
	target := 12

	idx := ExponentialSearch(arr, target)

	fmt.Printf("Массив: ... (огромный отсортированный)\n")
	if idx != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, idx)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
