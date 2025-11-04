package merge_sort

/*
	Алгоритм сортировки слиянием (Merge Sort) использует принцип "разделяй и властвуй":
	1) Делим массив пополам.
	2) Рекурсивно сортируем каждую половину.
	3) Сливаем две отсортированные половины в одну.

	Time Complexity: O(n log n) — деление на log n уровней, слияние O(n) на каждом.
	Space Complexity: O(n) — для временных массивов left, right и result.
*/

// Функция для слияние двух отсортированных массивов
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Сравниваем элементы и добавляем меньший в результат
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Основная функция сортировки слиянием
func MergeSort(arr []int) []int {
	// Обработка краевых случаев
	if arr == nil {
		return nil
	}
	// Базовый случай: массив длиной 1 уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Разделяем массив пополам
	mid := len(arr) / 2

	// Рекурсивно сортируем правую и левую часть
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Сливаем отсортированные части
	return merge(left, right)
}
