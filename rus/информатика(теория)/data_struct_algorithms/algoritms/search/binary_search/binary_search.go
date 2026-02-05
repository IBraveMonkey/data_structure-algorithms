package binary_search

/*
Бинарный поиск — это алгоритм поиска элемента в отсортированном массиве. Его суть в том, что на каждом шаге массив делится пополам, и проверяется, в какой половине находится искомый элемент. Это позволяет сократить количество проверок.

### Сложность

| Метрика | Сложность (O) |
|:---|:---:|
| Время | O(log n) |
| Пространственная | O(1) итеративно / O(log n) рекурсивно |

Плюсы: высокая скорость на больших данных.
Минусы: требует предварительной сортировки массива.
*/

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	if target < arr[left] || target > arr[right] {
		return -1
	}

	for left <= right {
		mid := (right + left) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			right = mid + 1
		} else {
			left = mid - 1
		}
	}

	return -1
}

/*
	Найти корень числа - написать функцию, которая находит корень числа или ближайшее подходящее наименьшее целое число

Например - для 9 это будет 3
Для 21 будет 4
5 не подойдет, потому что квадрат 5 будет 25, что больше, чем 21
*/
func binarySearchSqrt(target int) int {
	left := 0
	right := target

	for left <= right {
		middle := (left + right) / 2

		if middle*middle > target {
			right = middle - 1
			continue
		}

		if middle*middle < target {
			left = middle + 1
			continue
		}

		return middle
	}

	return right
}
