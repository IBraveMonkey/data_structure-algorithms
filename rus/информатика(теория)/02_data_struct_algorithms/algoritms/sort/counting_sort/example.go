package counting_sort

import "fmt"

func Example() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Printf("Original: %v\n", arr)

	// Counting Sort часто возвращает новый массив или модифицирует текущий
	// В нашей реализации он меняет текущий in-place (псевдо) и возвращает его
	sorted := CountingSort(arr)
	fmt.Printf("Sorted:   %v\n", sorted)

	// Пример с отрицательными числами (наша реализация это поддерживает благодаря offset)
	arr2 := []int{-5, -10, 0, -3, 8, 5, -1, 10}
	fmt.Printf("Отсортированный (с отрицательными): %v\n", CountingSort(arr2))
}
