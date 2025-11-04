package two_pointers

import "strings"

/*
Two Pointer: Метод двух указателей

	Кратко: Чтобы решать задачи быстрее, чем перебором, используя два индекса, которые двигаются по массиву или строке.
	Смысл: Вместо проверки всех пар (O(n²)) двигаем указатели умно (O(n)), опираясь на условие задачи.
*/

/* Left-Right: Отсортировано? Симметрия? Два конца? → Используй!
Slow-Fast: Циклы? Дубликаты? Один бежит, другой догоняет? → Используй!
*/
// Найти числа, который дадут в сумме target
func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

// Является ли строка паллиндромом
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	cleaned := ""

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			cleaned += string(char)
		}
	}

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

// Развернуть массив in-place
func ReverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// Удалить дубликаты, вернуть кол-во уникальный эл-ов
// [1,1,2,3,3]
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0                                 // Медленный указатель - позиция для уникального элемента
	for fast := 1; fast < len(nums); fast++ { // Быстрый указатель - проверяет след элемент
		if nums[fast] != nums[slow] { // Если нашли уникальный элемент
			slow++                  // Двигаем slow вперед
			nums[fast] = nums[slow] // Записываем уникальный элемент
		}
	}

	return slow + 1 // Новая длина - кол-во уникальных элементов
}
