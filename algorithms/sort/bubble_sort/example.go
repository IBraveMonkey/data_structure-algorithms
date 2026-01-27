package bubble_sort

import "fmt"

// Example демонстрирует использование пузырьковой сортировки с различными примерами
func Example() {
	// Пример 1: Базовая пузырьковая сортировка
	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Исходный массив: %v\n", arr1)
	BubbleSort(arr1)
	fmt.Printf("Отсортированный массив: %v\n", arr1)

	// Пример 2: Уже отсортированный массив
	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Уже отсортированный массив: %v\n", arr2)
	BubbleSort(arr2)
	fmt.Printf("После сортировки: %v\n", arr2)

	// Пример 3: Массив, отсортированный в обратном порядке
	arr3 := []int{5, 4, 3, 2, 1}
	fmt.Printf("Обратно отсортированный массив: %v\n", arr3)
	BubbleSort(arr3)
	fmt.Printf("После сортировки: %v\n", arr3)
}

// Задача: Отсортировать массив целых чисел по возрастанию
// Это базовая задача для пузырьковой сортировки
func SortArray(nums []int) []int {
	// Создаем копию, чтобы не менять оригинал
	result := make([]int, len(nums))
	copy(result, nums)

	// Применяем пузырьковую сортировку
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(result)-1; i++ {
			if result[i] > result[i+1] {
				result[i], result[i+1] = result[i+1], result[i]
				sorted = false
			}
		}
	}

	return result
}

// Задача: Отсортировать массив с минимальным количеством перестановок
// Хотя bubble sort не оптимален для этого, он демонстрирует концепцию
func MinSwapsToSort(nums []int) int {
	// Создаем копию
	arr := make([]int, len(nums))
	copy(arr, nums)

	swaps := 0
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swaps++
				sorted = false
			}
		}
	}

	return swaps
}
