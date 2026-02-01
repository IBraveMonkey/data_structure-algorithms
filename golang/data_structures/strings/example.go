package strings

import (
	"fmt"
	stdstrings "strings"
)

// Example демонстрирует алгоритмы поиска строк и задачи
func Example() {
	// ==========================================
	// 1. Демонстрация свойств строк
	// ==========================================
	StringInfo()
	fmt.Println("\n-------------------")

	// ==========================================
	// 2. Алгоритмы поиска (Brute Force, BMH)
	// ==========================================
	text := "exampletext"
	pattern := "ext"

	// Brute Force
	idx := BruteForce(text, pattern)
	fmt.Printf("BruteForce: '%s' в '%s' найден на индексе %d\n", pattern, text, idx)

	// Boyer-Moore-Horspool
	idxBMH := BMHSearch(text, pattern)
	fmt.Printf("BMHSearch: '%s' в '%s' найден на индексе %d\n", pattern, text, idxBMH)

	fmt.Println("-------------------")

	// ==========================================
	// 3. Задачи на строки
	// ==========================================

	// Задача: Самый длинный общий префикс
	strs := []string{"flower", "flow", "flight"}
	lcp := LongestCommonPrefix(strs)
	fmt.Printf("Самый длинный общий префикс для %v: '%s'\n", strs, lcp)
}

// -----------------------------------------------------------
// Алгоритмы (перенесены из strings.go)
// -----------------------------------------------------------

/*
Метод грубой силы (Brute Force)
  - Начинаем с первого символа в строке
  - Сравниваем слева направо каждый символ из строки с подстрокой
  - Time Complexity: O(n * m)
*/
func BruteForce(text, pattern string) int {
	tRune := []rune(text)
	pRune := []rune(pattern)

	n := len(tRune)
	m := len(pRune)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && tRune[i+j] == pRune[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

/*
Алгоритм Бойера - Мура - Хорспула (BMH)
  - Использует таблицу смещений для пропуска заведомо несовпадающих частей.
  - Time Complexity: O(n * m) в худшем, но на практике часто быстрее.
*/
func createShiftTable(pattern string) map[byte]int {
	m := len(pattern)
	table := make(map[byte]int)

	for i := 0; i < m-1; i++ {
		table[pattern[i]] = m - 1 - i
	}
	return table
}

func BMHSearch(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m > n {
		return -1
	}

	shiftTable := createShiftTable(pattern)
	i := 0

	for i <= n-m {
		j := m - 1

		for j >= 0 && text[i+j] == pattern[j] {
			j--
		}

		if j < 0 {
			return i
		}

		shift, exists := shiftTable[text[i+m-1]]
		if exists {
			i += shift
		} else {
			i += m
		}
	}

	return -1
}

// -----------------------------------------------------------
// Задачи
// -----------------------------------------------------------

// Задача: Самый длинный общий префикс
// Написать функцию для поиска самого длинного общего префикса строки среди массива строк.
// Если общего префикса нет, вернуть пустую строку "".
//
// Пример: ["flower","flow","flight"] => "fl"
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !stdstrings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
