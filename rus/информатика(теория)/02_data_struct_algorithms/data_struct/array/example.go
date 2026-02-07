package array

import "fmt"

// Examples показывает практические приемы работы с массивами
func Examples() {
	// Реверс массива (In-place)
	// Сложность: O(n) по времени, O(1) по памяти
	data := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	fmt.Println("Реверс:", data)

	// Поиск максимума
	// Сложность: O(n)
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	fmt.Println("Максимум:", max)

	// Фильтрация (создание нового среза)
	// Сложность: O(n)
	even := make([]int, 0)
	for _, v := range data {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	fmt.Println("Четные:", even)
}

/*
Важные моменты про Slices в Go:
1. len(s) - количество элементов в срезе.
2. cap(s) - емкость базового массива.
3. Срез — это дескриптор (указатель на массив, длина, емкость).
   Передача среза в функцию копирует дескриптор, но не сами данные.
*/
