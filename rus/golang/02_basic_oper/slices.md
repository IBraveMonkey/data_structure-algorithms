# 🍕 Слайсы и Массивы: Полный разбор (Slices & Arrays Deep Dive)

В Go массивы имеют фиксированный размер, а слайсы — это динамические дескрипторы поверх них. Ниже приведены все детальные примеры работы с ними, включая паники и ошибки компиляции.

---


### 🚦 Причины паники и ошибок (Panic & Error Causes)

| Операция | Массив | Слайс |
| :--- | :--- | :--- |
| **Индекс вне границ** | 🔥 **PANIC** (или compile error) | 🔥 **PANIC** |
| **Доступ к Nil** | N/A (всегда инициализирован) | 🔥 **PANIC** |
| **Make с отр. размером** | Compile error | 🔥 **PANIC** (в runtime) |
| **Реслайс (high > cap)** | N/A | 🔥 **PANIC** |

---


### 💻 Примеры кода: Массивы (Array Examples)

```go
package main

import (
	"fmt"
	"unsafe"
)

// Ошибка: panic при доступе к элементу массива с индексом вне диапазона
func accessToArrayElement1() {
	data := [3]int{1, 2, 3}
	idx := 4               
	// fmt.Println(data[idx]) // 🔥 Panic
	// fmt.Println(data[4])   // ❌ Compilation error
}

// Ошибка: panic при доступе к элементу массива с отрицательным индексом
func accessToArrayElement2() {
	data := [3]int{1, 2, 3}
	idx := -1              
	// fmt.Println(data[idx]) // 🔥 Panic
	// fmt.Println(data[-1])  // ❌ Compilation error
}

// Работает: возвращает длину массива
func arrayLen() {
	data := [10]int{}      
	fmt.Println(len(data)) // 10
}

// Работает: возвращает емкость массива
func capArray() {
	var data [10]int       
	fmt.Println(cap(data)) // 10
}

// Работает: сравнивает массивы (сравнимы только если элементы сравнимы)
func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}
	fmt.Println(first == second) // true
}

// Работает: размер массива в байтах
func emptyArray() {
	var data [10]byte                
	fmt.Println(unsafe.Sizeof(data)) // 10
}

// Ошибка: создание массива с переменной длиной
func arrayCreation() {
	length1 := 100
	// var data1 [length1]int // ❌ Compilation error
	
	const length2 = 100 
	var data2 [length2]int // ✅ OK: константа известна при компиляции
}
```

---


### 💻 Примеры кода: Слайсы (Slice Examples)

```go
// Ошибка: доступ вне границ слайса
func accessToSliceElement1() {
	data := make([]int, 3)
	// fmt.Println(data[4]) // 🔥 Panic
}

// Ошибка: доступ вне длины (len), даже если меньше емкости (cap)
func accessToSliceElement2() {
	data := make([]int, 3, 6)
	// fmt.Println(data[4]) // 🔥 Panic
}

// Ошибка: доступ к nil-слайсу
func accessToNilSlice1() {
	var data []int
	// _ = data[0] // 🔥 Panic
}

// Нормально: append обрабатывает nil-слайс
func appendToNilSlice() {
	var data []int
	data = append(data, 10) // ✅ OK: слайс инициализируется
}

// Нормально: итерация по nil-слайсу
func rangeByNilSlice() {
	var data []int
	for range data { // ✅ OK: 0 итераций
	}
}

// Ошибка: создание слайса с неверными параметрами
func makeSliceErrors() {
	// _ = make([]int, 10, 5) // ❌ Compilation error: len > cap
	
	size := -5
	// _ = make([]int, size)  // 🔥 Panic: negative size
}

// Работает: реслайсинг в пределах емкости (capacity)
func sliceMoreThanSize() {
	data := make([]int, 2, 6) // len=2, cap=6
	slice1 := data[1:6]       // ✅ OK
}

// Ошибка: паника при попытке увеличить емкость через реслайс сверх исходной
func increaseCapacityError() {
	data := make([]int, 0, 10)
	// data = data[:10:100] // 🔥 Panic!
}
```

> [!IMPORTANT]
> Для проверки "пустоты" (и nil, и `{}`) всегда используйте `len(s) == 0`. Внутренне `nil` и `[]T{}` различаются адресом данных, но оба имеют нулевую длину.

<!-- QUIZ_START 

[
    {
        "question": "В чем главное отличие массива от слайса в Go?",
        "options": [
            "Слайс всегда быстрее",
            "Массив имеет фиксированный размер, а слайс — это динамический дескриптор поверх массива",
            "Массивы всегда хранятся в куче",
            "Слайсы нельзя передавать в функции"
        ],
        "correctIndex": 1
    },
    {
        "question": "Что произойдет при попытке обратиться по индексу, который больше len, но меньше cap слайса?",
        "options": [
            "Вернется нулевое значение",
            "Слайс автоматически расширится",
            "Произойдет panic (index out of range)",
            "Вернется ошибка"
        ],
        "correctIndex": 2
    },
    {
        "question": "Как функция append() ведет себя с nil-слайсом?",
        "options": [
            "Вызывает panic",
            "Ничего не делает",
            "Корректно инициализирует слайс и добавляет в него элементы",
            "Возвращает nil"
        ],
        "correctIndex": 2
    }
]

QUIZ_END -->

