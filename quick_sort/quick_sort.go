package quick_sort

/*
 Быстрая сортировка

	Выбираем опорный элемент - pivot
	В идеале - медиану
	Элементы, меньшие опорного элемента, идут в одну часть,
	Элементы, большие или равные опорному - в другую

	После этого обе части рекурсивно сортируются

	spaceComplexity in place работает по памяти, но из-за рекурсий будет O(log n) в худшем случае будет O(n)

	timeComplexity
	в худшем О(n^2), если опорный элемент постоянно делит массив крайне неравномерно(например, если массив уже отсортирован, а мы всегда вырираем первый элемент)

	в лучшем О(n log n)
*/

func medianOfThree(arr []int, low, hight int) int {
	mid := (low + hight) / 2

	// Сравнение трех элементов: первый, средний, последний
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}

	if arr[low] > arr[hight] {
		arr[low], arr[hight] = arr[hight], arr[low]
	}

	if arr[mid] > arr[hight] {
		arr[mid], arr[hight] = arr[hight], arr[mid]
	}

	// Медиана теперь в arr[mid], возвращаем индекс
	return mid
}

// В это случае будет O(n log n)
func QuickSort(arr []int) []int {
	// Базовый случай, если длина массива <= 1 он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (pivot)
	pivotIndex := medianOfThree(arr, 0, len(arr)-1)
	pivot := arr[pivotIndex]
	left := []int{}
	right := []int{}

	// разделяем массив на элементы меньше или больше pivot
	for i := 0; i < len(arr); i++ {
		if i == pivotIndex {
			continue
		}

		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	sortedLeft := QuickSort(left)   // сортируем левую часть
	sortedRight := QuickSort(right) // сортируем правую часть

	// Складываем результат, левый массив + pivot + правый массив
	result := append(sortedLeft, pivot)
	result = append(result, sortedRight...)

	return result
	// Рекурсивно сортируем левую и правую части и объуединяем pivot
	// return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
