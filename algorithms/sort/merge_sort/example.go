package merge_sort

import "fmt"

func Example() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("Original: %v\n", arr)

	sorted := MergeSort(arr)
	fmt.Printf("Sorted:   %v\n", sorted)

	// Задача: Сортировка большого массива (симуляция)
	// В Merge Sort часто удобно видеть этапы слияния, но здесь просто демонстрация
}
