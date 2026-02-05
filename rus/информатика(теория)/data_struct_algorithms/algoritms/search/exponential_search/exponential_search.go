package exponential_search

import (
	"math"
)

/*
Exponential Search (Экспоненциальный поиск)

Что это такое?
Это алгоритм поиска в отсортированном, бесконечном или очень большом массиве, где размер неизвестен или слишком велик.
Он работает в два этапа: сначала находит диапазон, где может быть элемент (умножая индекс на 2), а затем запускает бинарный поиск в этом диапазоне.

Зачем это нужно?
- Для поиска в неограниченных (unbounded) или потоковых массивах.
- Работает быстрее бинарного поиска, если искомый элемент находится близко к началу массива (O(log i) против O(log n)).

В чём смысл?
- Мы "скачем" по массиву индексами 1, 2, 4, 8, 16... пока текущий элемент меньше искомого.
- Как только перепрыгнули (arr[i] > target), мы знаем, что элемент где-то между [i/2, i].

Когда использовать?
- Массив отсортирован, но размер огромен/неизвестен.
- Есть высокая вероятность, что элементы находятся в начале.

Как работает?
1. Проверяем 0-й элемент.
2. Инициализируем bound = 1.
3. Пока bound < len(arr) и arr[bound] < target:
   bound *= 2.
4. Запускаем Binary Search в диапазоне [bound/2, min(bound, len)].

Сложность:
- Время: O(log i), где i — индекс искомого элемента.
- Память: O(1).
*/

func ExponentialSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// 1. Проверяем первый элемент
	if arr[0] == target {
		return 0
	}

	// 2. Находим диапазон
	bound := 1
	for bound < n && arr[bound] <= target {
		bound *= 2
	}

	// 3. Бинарный поиск
	// Диапазон: [bound/2, min(bound, n-1)]
	left := bound / 2
	right := int(math.Min(float64(bound), float64(n-1)))

	return binarySearch(arr, left, right, target)
}

func binarySearch(arr []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
