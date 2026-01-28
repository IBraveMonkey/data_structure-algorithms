package main

import (
	"fmt"
	"unsafe"
)

// Ошибка: panic при доступе к элементу массива с индексом вне диапазона
func accessToArrayElement1() {
	data := [3]int{1, 2, 3}

	idx := 4               // можно скрыть(обмануть), но в runtime поймаем ошибку
	fmt.Println(data[idx]) // panic

	fmt.Println(data[4]) // compilation error
}

// Ошибка: panic при доступе к элементу массива с отрицательным индексом
func accessToArrayElement2() {
	data := [3]int{1, 2, 3}

	idx := -1              // можно скрыть(обмануть), но в runtime поймаем ошибку, мы не Питухон если что(если не ошибаюсь)
	fmt.Println(data[idx]) // panic

	fmt.Println(data[-1]) // compilation error
}

// Работает: возвращает длину массива
func arrayLen() {
	data := [10]int{}      // создается на 10 нулей
	fmt.Println(len(data)) // 10
}

// Работает: возвращает емкость массива
func capArray() {
	var data [10]int       // создает на 10 нулей
	fmt.Println(cap(data)) // 10
}

// Работает: сравнивает массивы
func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	//	[<, <=, >, >=]  ->  compilation error
}

// Работает: возвращает размер массива
func emptyArray() {
	var data [10]byte                // создается на 10 нулей -> byte = uin8 -> default for uint8 = 0
	fmt.Println(unsafe.Sizeof(data)) // 10

	//data == nil // compilation error
}

// Работает: возвращает размер массива
func zeroArray() {
	var data [0]int
	fmt.Println(unsafe.Sizeof(data)) // 0
}

// Ошибка: compilation error при создании массива с отрицательной длиной
func negativeArray() {
	var data [-1]int // compilation error
	_ = data
}

// Ошибка: compilation error при создании массива с переменной длиной
func arrayCreation() {
	length1 := 100
	var data1 [length1]int // compilation error
	_ = length1
	_ = data1

	const length2 = 100 // так работает, потому что во время compile мы знаем значение
	var data2 [length2]int
	_ = data2
}

// Ошибка: compilation error при использовании make для массива(только для slice\map\chan)
func makeArray() {
	_ = make([10]int, 10) // compilation error
}

// Ошибка: только для slice
func appendToArray() {
	_ = append([10]int{}, 10) // compilation error
}

// Ошибка
func accessToSliceElement1() {
	data := make([]int, 3)
	fmt.Println(data[4]) // panic
}

// Ошибка
func accessToSliceElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4]) // panic
}

// Ошибка
func accessToElement3() {
	data := make([]int, 3, 6)
	_ = data[-1] // compilation error
}

// Ошибка
func accessToNilSlice1() {
	var data []int
	_ = data[0] // panic
}

// Ошибка
func accessToNilSlice2() {
	var data []int
	data[0] = 10 // panic
}

// Нормально
func appendToNilSlice() {
	var data []int
	data = append(data, 10) // ok
}

// Нормально
func rangeByNilSlice() {
	var data []int
	for range data { // ok
	}
}

// Работает: создает срез нулевой длины
func makeZeroSlice() {
	data := make([]int, 0)
	fmt.Println(len(data)) // 0
	fmt.Println(cap(data)) // 0
}

// Ошибка: compilation error или panic при создании среза с неверными параметрами
func makeSlice() {
	_ = make([]int, -5)    // compilation error
	_ = make([]int, 10, 5) // compilation error

	size := -5
	_ = make([]int, size) // panic

	size = 5
	_ = make([]int, size*2, size) // panic
}

// Работает: создает срез с индексами в пределах capacity
func sliceMoreThanSize() {
	data := make([]int, 2, 6) // [0,0],0,0,0,0

	slice1 := data[1:6] // ok // [0,0,0,0,0,0]
	_ = slice1
}

// Ошибка: panic или compilation error при создании среза с неверными индексами
func sliceWithIncorrectIndeces() {
	data := make([]int, 2, 6)

	slice2 := data[1:7] // panic
	_ = slice2

	slice3 := data[2:1] // compilation error
	_ = slice3

	left := 2
	right := 1
	slice4 := data[left:right] // panic
	_ = slice4
}

// Работает / Ошибка: нормальная работа или panic при создании среза из nil среза
func sliceWithNilSlice() {
	var data []int

	slice := data[:]  // ok
	slice = data[0:0] // ok
	slice = data[0:1] // panic
	_ = slice
}

// Ошибка: panic при попытке увеличить максимальную емкость за пределы исходной
func increaseCapacity() {
	data := make([]int, 0, 10)
	data = data[:10:100] // panic
}
