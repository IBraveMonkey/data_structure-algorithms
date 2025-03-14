package merge_sort

/*
	Делается через - разделяй и властвуй

	1) Делим массив пополам
	2) Сортируем каждую половину независимо
	3) Объединяем две отсортированные последовательности
	4) Алгоритм слияние - 2 указателя

	TimeComplexity - O(n log n)
*/

// Функция для слияние двух отсортированных массивов
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Сравниваем элементы двух массивов
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
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
	// Базовый случай: массив длиной 1 уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Разделяем массив пополам
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Сливаем отсортированные части
	return merge(left, right)
}
