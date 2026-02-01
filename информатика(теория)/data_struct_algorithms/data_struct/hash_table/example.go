package hash_table

import (
	"fmt"
	"sort"
)

// Example демонстрирует использование хеш-таблицы на различных примерах
func Example() {
	// Пример 1: Найти лишнюю букву
	a := "abcd"
	b := "abcde"
	result := ExtraLetter(a, b)
	fmt.Printf("Лишняя буква в '%s' по сравнению с '%s': %s\n", b, a, result)

	// Пример 2: Сумма двух элементов
	nums := []int{2, 7, 11, 15}
	target := 9
	indices := TwoSum(nums, target)
	fmt.Printf("Индексы двух чисел для суммы %d в %v: %v\n", target, nums, indices)

	// Пример 3: Группировка анаграмм
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagramGroups := GroupAnagrams(strs)
	fmt.Printf("Группы анаграмм для %v: %v\n", strs, anagramGroups)

	// Пример 4: Подсчет частоты
	numsFreq := []int{1, 2, 3, 2, 1, 1}
	freqMap := CountFrequency(numsFreq)
	fmt.Printf("Частота чисел для %v: %v\n", numsFreq, freqMap)

	// Пример 5: Проверка дубликатов
	hasDup := ContainsDuplicate([]int{1, 2, 3, 1})
	fmt.Printf("Есть дубликаты в [1,2,3,1]: %t\n", hasDup)

	hasDup = ContainsDuplicate([]int{1, 2, 3, 4})
	fmt.Printf("Есть дубликаты в [1,2,3,4]: %t\n", hasDup)
}

// Задача 1: Найти разницу между двумя строками
// На вход подается две строки a и b.
// Строка b образована из строки a путем перемешивания и добавления одной буквы.
// Необходимо вернуть эту букву.
//
// Пример: a = "abc", b = "beca" => 'e'
// Временная сложность: O(n), пространственная сложность: O(n)
func ExtraLetter(a, b string) string {
	hTable := make(map[rune]int, len(b))

	// Подсчитываем частоту символов в строке b
	for _, v := range b {
		hTable[v]++
	}

	// Уменьшаем счетчики для символов из строки a
	for _, v := range a {
		hTable[v]--
	}

	// Находим символ с ненулевым счетчиком
	for key, val := range hTable {
		if val > 0 {
			return string(key)
		}
	}

	return ""
}

// Задача 2: Сумма двух элементов массива
// Дан неотсортированный массив целых чисел и некоторое число target.
// Необходимо написать функцию, которая найдет два таких элемента в массиве,
// сумма которых будет равна target.
//
// Пример: data = [2, 7, 11, 15], target = 9 => [0, 1] (data[0] + data[1] = 2 + 7 = 9)
// Временная сложность: O(n), пространственная сложность: O(n)
func TwoSum(data []int, target int) []int {
	// Создаем хеш-таблицу для хранения значений и их индексов
	cache := make(map[int]int)

	for i, num := range data {
		complement := target - num // Искомое дополнение до target

		// Проверяем, есть ли дополнение в хеш-таблице
		if index, ok := cache[complement]; ok {
			return []int{index, i} // Возвращаем индексы двух элементов
		}

		// Сохраняем текущий элемент и его индекс
		cache[num] = i
	}

	return []int{} // Если пара не найдена
}

// Задача 3: Массив Анаграмм
// Необходимо найти и сгруппировать слова-анаграммы вместе.
// Анаграмма - это слово или фраза, полученная путем перестановки букв другого слова или фразы.
//
// Пример: ["eat", "tea", "tan", "ate", "nat", "bat"] => [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
// Временная сложность: O(n * m * log(m)), где n - количество строк, m - средняя длина строки
func GroupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		// Сортируем символы строки, чтобы получить общий ключ для анаграмм
		sortedStr := sortString(str)

		// Добавляем строку в группу с соответствующим ключом
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// Преобразуем map в слайс слайсов
	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

// Вспомогательная функция для сортировки символов в строке
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Задача 4: Подсчет частоты элементов
// Подсчитать, сколько раз встречается каждый элемент в массиве
//
// Пример: [1, 2, 3, 2, 1, 1] => {1: 3, 2: 2, 3: 1}
// Временная сложность: O(n), пространственная сложность: O(k), где k - количество уникальных элементов
func CountFrequency(nums []int) map[int]int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++ // Увеличиваем счетчик для текущего элемента
	}

	return freqMap
}

// Задача 5: Проверка наличия дубликатов
// Проверить, есть ли в массиве повторяющиеся элементы
//
// Пример: [1, 2, 3, 1] => true, [1, 2, 3, 4] => false
// Временная сложность: O(n), пространственная сложность: O(n)
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true // Найден дубликат
		}
		seen[num] = true
	}

	return false // Дубликатов нет
}

// Задача 6: Пересечение двух массивов
// Найти элементы, которые присутствуют в обоих массивах
//
// Пример: [1, 2, 2, 1], [2, 2] => [2]
// Временная сложность: O(n + m), пространственная сложность: O(min(n, m))
func Intersection(nums1, nums2 []int) []int {
	set1 := make(map[int]bool)
	for _, num := range nums1 {
		set1[num] = true
	}

	set2 := make(map[int]bool)
	for _, num := range nums2 {
		if set1[num] {
			set2[num] = true
		}
	}

	result := make([]int, 0, len(set2))
	for num := range set2 {
		result = append(result, num)
	}

	return result
}

// Задача 7: Первый неповторяющийся символ
// Найти первый символ, который встречается в строке только один раз
//
// Пример: "leetcode" => 'l', "loveleetcode" => 'v'
// Временная сложность: O(n), пространственная сложность: O(k), где k - количество уникальных символов
func FirstUniqueChar(s string) int {
	charCount := make(map[rune]int)

	// Подсчитываем частоту каждого символа
	for _, char := range s {
		charCount[char]++
	}

	// Находим первый символ с частотой 1
	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1 // Нет уникального символа
}

// Задача 8: Палиндром из перестановки
// Проверить, можно ли переставить символы строки так, чтобы получить палиндром
//
// Пример: "aab" => true ("aba"), "carerac" => true ("racecar")
// Временная сложность: O(n), пространственная сложность: O(k), где k - количество уникальных символов
func CanPermutePalindrome(s string) bool {
	charCount := make(map[rune]int)

	// Подсчитываем частоту каждого символа
	for _, char := range s {
		charCount[char]++
	}

	// Подсчитываем количество символов с нечетной частотой
	oddCount := 0
	for _, count := range charCount {
		if count%2 == 1 {
			oddCount++
		}
	}

	// Палиндром возможен, если не более одного символа имеет нечетную частоту
	return oddCount <= 1
}

// Задача 9: Содержит дубликат II
// Дан массив целых чисел nums и целое число k. Вернуть true, если в массиве есть два различных индекса i и j,
// таких что nums[i] == nums[j] и abs(i - j) <= k.
func ContainsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, num := range nums {
		if prevIndex, exists := indexMap[num]; exists {
			if i-prevIndex <= k {
				return true
			}
		}
		// Обновляем последний индекс этого числа
		indexMap[num] = i
	}

	return false
}

// Задача 10: Валидный Судоку
// Определить, является ли доска Судоку 9x9 валидной. Необходимо проверить только заполненные ячейки.
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// Инициализация карт
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}

			// Проверка строки
			if rows[i][cell] {
				return false
			}
			rows[i][cell] = true

			// Проверка столбца
			if cols[j][cell] {
				return false
			}
			cols[j][cell] = true

			// Проверка квадрата 3x3
			boxIndex := (i/3)*3 + j/3
			if boxes[boxIndex][cell] {
				return false
			}
			boxes[boxIndex][cell] = true
		}
	}

	return true
}

// Задача 11: Группировка сдвинутых строк
// Дана строка, мы можем сдвинуть ее вправо на одну позицию, чтобы получить новую строку.
// Например, "abc" -> "bcd". Строка группируется с другими, которые могут быть сдвинуты, чтобы сформировать друг друга.
func GroupShiftedStrings(strings []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strings {
		// Создаем ключ на основе относительных сдвигов от первого символа
		key := ""
		if len(s) > 0 {
			base := rune(s[0])

			for _, c := range s {
				// Вычисляем относительный сдвиг, обрабатывая цикличность (26 символов)
				shift := (c - base + 26) % 26
				key += string(shift) + ","
			}
		}

		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
