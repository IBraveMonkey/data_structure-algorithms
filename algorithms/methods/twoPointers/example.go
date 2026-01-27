package two_pointers

import (
	"fmt"
	"strings"
)

// Example демонстрирует задачи на Two Pointers
func Example() {
	// 1. Two Sum (для отсортированного массива)
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("TwoSum индексы для %v, target %d: %v\n", nums, target, TwoSumSorted(nums, target))

	// 2. Проверка палиндрома
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("IsPalindrome ('%s'): %v\n", s, IsPalindrome(s))

	// 3. Удаление дубликатов
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	lenAfter := RemoveDuplicates(arr)
	fmt.Printf("RemoveDuplicates - new len: %d, arr prefix: %v\n", lenAfter, arr[:lenAfter])
}

// Задача 1: Two Sum (Input Array Is Sorted)
// Найти два числа, которые в сумме дают target. Вернуть их индексы (1-based, как в LeetCode, или 0-based).
// Здесь возвращаем 0-based.
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++ // Сумма слишком маленькая, нужно больше -> двигаем левый вправо
		} else {
			right-- // Сумма слишком большая, нужно меньше -> двигаем правый влево
		}
	}

	return []int{}
}

// Задача 2: Валидация палиндрома
func IsPalindrome(s string) bool {
	// В реальной задаче лучше обрабатывать руны и не создавать новую строку для O(1) space,
	// но для примера упростим подготовку строки.
	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r + 32 // toLower
		}
		return -1
	}, s)

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// Задача 3: Удаление дубликатов из отсортированного массива
// Возвращает новую длину (k).
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// Если нашли новый уникальный элемент
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// Задача 4: Разворот массива In-Place
func ReverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
