package binary_search

import "fmt"

// Example демонстрирует использование бинарного поиска
func Example() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	idx := BinarySearch(arr, target)
	fmt.Printf("Массив: %v\n", arr)
	fmt.Printf("Поиск числа %d: найден на индексе %d\n", target, idx)

	// Задача: Поиск позиции вставки
	nums := []int{1, 3, 5, 6}
	val := 5
	pos := SearchInsert(nums, val)
	fmt.Printf("Куда вставить %d в %v? Индекс %d\n", val, nums, pos)

	val = 2
	pos = SearchInsert(nums, val)
	fmt.Printf("Куда вставить %d в %v? Индекс %d\n", val, nums, pos)
}

// Задача: Поиск позиции вставки (Search Insert Position)
// Учитывая отсортированный массив и целевое значение, вернуть индекс, если цель найдена.
// Если нет, вернуть индекс, где она была бы, если бы была вставлена по порядку.
// Пример: [1,3,5,6], 5 -> 2
// Пример: [1,3,5,6], 2 -> 1
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// Задача: Первый плохой релиз (First Bad Version)
// Представьте, что вы product manager. Выпустить плохую версию означает, что все последующие версии тоже плохие.
// Нужно найти первую плохую версию с минимальным количеством проверок API isBadVersion(version).
func FirstBadVersion(n int, isBadVersion func(int) bool) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid // Плохая версия может быть этой или левее
		} else {
			left = mid + 1 // Эта версия хорошая, значит плохая точно правее
		}
	}

	return left
}
