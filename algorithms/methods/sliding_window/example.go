package sliding_window

import (
	"fmt"
	"math"
)

// Example демонстрирует задачи на Sliding Window
func Example() {
	// 1. Максимальное среднее в подмассиве длины k
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("Макс. среднее (k=%d): %.2f\n", k, FindMaxAverage(nums, k))

	// 2. Минимальная длина подмассива с суммой >= target
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Printf("Мин. длина подмассива (сумма >= %d): %d\n", target, MinSubArrayLen(target, arr))
}

// Задача 1: Найти максимальное среднее значение подмассива длины k
// O(N) time, O(1) space
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	// 1. Инициализируем сумму первого окна
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	// 2. Скользим (начинаем с k-го элемента)
	for i := k; i < len(nums); i++ {
		sum += nums[i]   // Добавляем новый элемент (справа)
		sum -= nums[i-k] // Удаляем старый элемент (слева)
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

// Задача 2: Минимальный размер подмассива (Variable Size Sliding Window)
// Найти минимальную длину непрерывного подмассива, сумма которого >= target.
// Если такого нет, вернуть 0.
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // Расширяем окно вправо

		// Сжимаем окно слева, пока условие выполняется
		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
