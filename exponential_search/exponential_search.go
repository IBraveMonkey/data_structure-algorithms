package exponential_search

/* Используется для уточнения диапазона поиска
В итоговом диапазоне поиск осуществляется с помощью бинарного поиска
TimeComplexity - O(log(m)) m - индекс элемента, который нужно найти

Как работает, берем border - это первый элемент, потом его увеличиваем в 2 раза, если этот элемент больше, чем target, то в этом отрезке делаем бинарный поиск
*/

/*
Задача - Поиск результирующего отрезка

# Дан массив целых чисел, отсортированнй по возрастанию, и некоторое число target

Необходимо установить отрезок массива, в котором может распологаться это число
*/
func ExponentialSearch(arr []int, target int) []int {
	border := 1
	lastElement := len(arr) - 1

	for border < lastElement && arr[border] < target {
		if arr[border] == target {
			return []int{border, border * 2}
		}

		border = border * 2

		if border > lastElement {
			return []int{border / 2, lastElement}
		}
	}

	return []int{(border / 2), border}
}

// [1,4,7,9,12,18,29], target - 12
func test(arr []int, target int) []int {
	border := 1
	lastElement := len(arr) - 1

	for border < lastElement && arr[border] < target {
		if arr[border] == target {
			return []int{border, border * 2}
		}

		border = border * 2

		if border > lastElement {
			return []int{border / 2, lastElement}
		}
	}

	return []int{border / 2, border}
}
