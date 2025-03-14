package bubble_sort

/*

	Алгоритм устойчив

Алгоритм:
	-Проходим по массиву, сравнивая пары элементов.
	-Если текущий элемент больше следующего, меняем их местами.
	-Повторяем процесс, пока массив не станет отсортированным.
	-На каждой итерации самый большой элемент «всплывает» в конец.

	TimeComplexity - O(n^2)
*/

func BubbleSort(arr []int) {
	sorted := false

	for !sorted {
		sorted = true

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}
