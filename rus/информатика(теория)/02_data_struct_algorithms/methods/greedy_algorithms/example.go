package greedy_algorithms

import (
	"fmt"
	"sort"
)

// Example демонстрирует жадные алгоритмы
func Example() {
	// 1. Сдача монетами (для канонических систем номиналов жадность работает)
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Printf("Мин. монет для %d с %v: %d\n", amount, coins, CoinChange(coins, amount))

	// 2. Непересекающиеся интервалы (Activity Selection Problem)
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Printf("Макс. кол-во непересекающихся интервалов: %d\n", MaxNonOverlapping(intervals))
}

// Задача 1: Минимальное количество монет (для удобных номиналов)
// Примечание: работает для (1, 2, 5, 10...), но может не работать для (1, 3, 4)
func CoinChange(coins []int, amount int) int {
	count := 0
	remaining := amount

	// Сортируем монеты по убыванию, чтобы брать самые крупные сначала
	// (в примере уже отсортированы, но для универсальности можно добавить sort)
	// sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	// Но так как у нас слайс интов и мы идем с конца - это ок

	for i := len(coins) - 1; i >= 0; i-- {
		for remaining >= coins[i] {
			remaining -= coins[i]
			count++
		}
	}

	if remaining != 0 {
		return -1
	}
	return count
}

// Задача 2: Выбор максимального числа непересекающихся интервалов
// (Аналог задачи о выборе заявок в аудитории)
// Жадная стратегия: всегда выбираем интервал, который заканчивается раньше всех,
// чтобы оставить больше места для остальных.
func MaxNonOverlapping(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// Сортируем по концу интервала
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		// Если начало текущего >= конца последнего выбранного -> они не пересекаются
		if intervals[i][0] >= lastEnd {
			count++
			lastEnd = intervals[i][1]
		}
	}
	return count
}
