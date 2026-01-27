package stack

import "fmt"

// Example демонстрирует использование стека на различных примерах
func Example() {
	// Создаем новый стек
	stack := &Stack{}

	// Добавляем элементы
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Стек после добавления 1, 2, 3:")

	// Смотрим верхний элемент
	if val, ok := stack.Peek(); ok {
		fmt.Printf("Верхний элемент: %d\n", val)
	}

	// Удаляем элементы
	if val, ok := stack.Pop(); ok {
		fmt.Printf("Удаленный элемент: %d\n", val)
	}

	// Проверяем пуст ли стек
	fmt.Printf("Стек пуст? %t\n", stack.IsEmpty())

	// Тест палиндрома
	testStr := "racecar"
	fmt.Printf("Является ли '%s' палиндромом? %t\n", testStr, IsPalindrome(testStr))

	testStr2 := "hello"
	fmt.Printf("Является ли '%s' палиндромом? %t\n", testStr2, IsPalindrome(testStr2))
}

// Задача 1: Валидные скобки
// Определить, является ли строка с различными типами скобок валидной.
func IsValidParentheses(s string) bool {
	stack := &Stack{}

	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack.Push(char) // сохраняем rune
		case ')', '}', ']':
			topInterface, ok := stack.Peek()
			if !ok {
				return false
			}
			top := topInterface.(rune)

			if top != mapping[char] {
				return false
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

// Задача 2: Обратная Польская Нотация (RPN)
// Вычислить значение арифметического выражения в обратной польской записи.
func EvalRPN(tokens []string) int {
	stack := &Stack{}

	for _, token := range tokens {
		switch token {
		case "+":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) + val2.(int))
		case "-":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) - val2.(int))
		case "*":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) * val2.(int))
		case "/":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) / val2.(int))
		default:
			// Токен это число
			var num int
			fmt.Sscanf(token, "%d", &num)
			stack.Push(num)
		}
	}

	result, _ := stack.Pop()
	return result.(int)
}

// Задача 3: Ежедневные температуры
// Для каждого дня вернуть количество дней до потепления.
func DailyTemperatures(temperatures []int) []int {
	stack := &Stack{} // Стек хранит индексы
	result := make([]int, len(temperatures))

	for i, temp := range temperatures {
		// Пока стек не пуст и текущая температура выше температуры по индексу в стеке
		for !stack.IsEmpty() {
			topIndexInter, _ := stack.Peek()
			topIndex := topIndexInter.(int)

			if temp <= temperatures[topIndex] {
				break
			}
			// Извлекаем индекс и считаем разницу
			stack.Pop()
			result[topIndex] = i - topIndex
		}
		// Добавляем текущий индекс
		stack.Push(i)
	}

	return result
}

// Задача 4: Проверка палиндрома (через Стек)
func IsPalindrome(a string) bool {
	q := &Stack{}

	for _, val := range a {
		q.Push(string(val))
	}

	for _, val := range a {
		popVal, _ := q.Pop()
		if string(val) != popVal.(string) {
			return false
		}
	}

	return true
}

// Задача 5: Проверка палиндрома (через 2 указателя)
func IsPalindromePointer(a string) bool {
	left := 0
	right := len(a) - 1

	for left < right {
		if a[left] != a[right] {
			return false
		}

		left++
		right--
	}

	return true
}
