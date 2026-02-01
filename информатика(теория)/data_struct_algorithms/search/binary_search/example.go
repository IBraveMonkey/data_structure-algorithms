package binary_search

import "fmt"

// Example демонстрирует использование бинарного поиска с различными примерами
func Example() {
	// Пример 1: Базовый бинарный поиск
	arr := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := BinarySearch(arr, target)
	fmt.Printf("Бинарный поиск числа %d в %v: индекс %d\n", target, arr, result)

	// Пример 2: Бинарный поиск несуществующего элемента
	target = 2
	result = BinarySearch(arr, target)
	fmt.Printf("Бинарный поиск числа %d в %v: индекс %d\n", target, arr, result)

	// Пример 3: Бинарный поиск квадратного корня
	fmt.Printf("Квадратный корень из 9: %d\n", binarySearchSqrt(9))
	fmt.Printf("Квадратный корень из 21: %d\n", binarySearchSqrt(21))
}

// Задача: Найти пиковый элемент (Find Peak Element)
// Пиковый элемент — это элемент, который строго больше своих соседей.
// Учитывая 0-индексированный массив целых чисел nums, найти пиковый элемент и вернуть его индекс.
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// Пик находится в левой части (включая mid)
			right = mid
		} else {
			// Пик находится в правой части
			left = mid + 1
		}
	}

	return left
}

// Задача: Поиск во вращающемся отсортированном массиве (Search in Rotated Sorted Array)
// Учитывая массив nums после вращения и целое число target, вернуть индекс target, если он есть в nums, или -1, если нет.
func SearchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Проверяем, отсортирована ли левая половина
		if nums[left] <= nums[mid] {
			// Target в левой половине
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				// Target в правой половине
				left = mid + 1
			}
		} else {
			// Правая половина отсортирована
			if nums[mid] < target && target <= nums[right] {
				// Target в правой половине
				left = mid + 1
			} else {
				// Target в левой половине
				right = mid - 1
			}
		}
	}

	return -1
}

// Задача: Найти минимум во вращающемся отсортированном массиве (Find Minimum in Rotated Sorted Array)
// Учитывая отсортированный вращающийся массив nums уникальных элементов, вернуть минимальный элемент.
func FindMinInRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Минимум в правой части
			left = mid + 1
		} else {
			// Минимум в левой части (включая mid)
			right = mid
		}
	}

	return nums[left]
}
