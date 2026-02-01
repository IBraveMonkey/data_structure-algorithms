package array

import "fmt"

// Example демонстрирует работу с массивами и срезами через задачи
func Example() {
	// Задача: Переместить нули в конец
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	fmt.Printf("После перемещения нулей: %v\n", nums)

	// Задача: Максимальная прибыль
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(prices)
	fmt.Printf("Максимальная прибыль: %d\n", profit)
}

// Задача 1: Перемещение нулей
// Дан массив целых чисел nums. Написать функцию, которая перемещает все 0 в конец массива,
// сохраняя относительный порядок ненулевых элементов.
//
// Пример: [0,1,0,3,12] => [1,3,12,0,0]
// Сложность: O(n), Пространство: O(1)
func MoveZeroes(nums []int) {
	insertPos := 0

	// Перемещаем все ненулевые элементы в начало
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	// Заполняем оставшуюся часть нулями
	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

// Задача 2: Лучшее время для покупки и продажи акций
// Дан массив цен. Нужно найти максимальную прибыль (покупка в один день, продажа в другой, более поздний).
//
// Пример: [7,1,5,3,6,4] => 5 (покупка за 1, продажа за 6)
// Сложность: O(n), Пространство: O(1)
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// Задача 3: Объединить отсортированные массивы
// Объединить два отсортированных массива nums1 и nums2 в nums1.
// Предполагается, что nums1 имеет достаточно места (m + n).
//
// Пример: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3 => [1,2,2,3,5,6]
func Merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// Если остались элементы в nums2
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
