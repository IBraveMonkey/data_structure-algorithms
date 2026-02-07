### 🍕 Slices and Arrays: Deep Dive

In Go, arrays have a fixed size, while slices are dynamic descriptors over them. Below are all detailed examples of working with them, including panics and compilation errors.

---

### 🚦 Causes of Panic and Errors (Panic & Error Causes)

| Operation | Array | Slice |
| :--- | :--- | :--- |
| **Index Out of Range** | 🔥 **PANIC** (or compile error) | 🔥 **PANIC** |
| **Nil Access** | N/A (always initialized) | 🔥 **PANIC** |
| **Make with Neg. Size** | Compile error | 🔥 **PANIC** (at runtime) |
| **Reslice (high > cap)** | N/A | 🔥 **PANIC** |

---

### 💻 Code Examples: Arrays (Array Examples)

```go
package main

import (
	"fmt"
	"unsafe"
)

// Error: panic on accessing array element with index out of range
// Ошибка: panic при доступе к элементу массива с индексом вне диапазона
func accessToArrayElement1() {
	data := [3]int{1, 2, 3}
	idx := 4               
	// fmt.Println(data[idx]) // 🔥 Panic
	// fmt.Println(data[4])   // ❌ Compilation error
}

// Error: panic on accessing array element with negative index
// Ошибка: panic при доступе к элементу массива с отрицательным индексом
func accessToArrayElement2() {
	data := [3]int{1, 2, 3}
	idx := -1              
	// fmt.Println(data[idx]) // 🔥 Panic
	// fmt.Println(data[-1])  // ❌ Compilation error
}

// Works: returns array length
// Работает: возвращает длину массива
func arrayLen() {
	data := [10]int{}      
	fmt.Println(len(data)) // 10
}

// Works: returns array capacity
// Работает: возвращает емкость массива
func capArray() {
	var data [10]int       
	fmt.Println(cap(data)) // 10
}

// Works: compares arrays (only if element types are comparable)
// Работает: сравнивает массивы (сравнимы только если элементы сравнимы)
func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}
	fmt.Println(first == second) // true
}

// Works: returns array size in bytes
// Работает: размер массива в байтах
func emptyArray() {
	var data [10]byte                
	fmt.Println(unsafe.Sizeof(data)) // 10
}

// Error: creating array with variable length
// Ошибка: создание массива с переменной длиной
func arrayCreation() {
	length1 := 100
	// var data1 [length1]int // ❌ Compilation error
	
	const length2 = 100 
	var data2 [length2]int // ✅ OK: constant is known at compile time
}
```

---

### 💻 Code Examples: Slices (Slice Examples)

```go
// Error: access beyond slice bounds
// Ошибка: доступ вне границ слайса
func accessToSliceElement1() {
	data := make([]int, 3)
	// fmt.Println(data[4]) // 🔥 Panic
}

// Error: access beyond len, even if within cap
// Ошибка: доступ вне длины (len), даже если меньше емкости (cap)
func accessToSliceElement2() {
	data := make([]int, 3, 6)
	// fmt.Println(data[4]) // 🔥 Panic
}

// Error: access to nil slice
// Ошибка: доступ к nil-слайсу
func accessToNilSlice1() {
	var data []int
	// _ = data[0] // 🔥 Panic
}

// Normal: append handles nil slice
// Нормально: append обрабатывает nil-слайс
func appendToNilSlice() {
	var data []int
	data = append(data, 10) // ✅ OK: slice is initialized
}

// Normal: range over nil slice
// Нормально: итерация по nil-слайсу
func rangeByNilSlice() {
	var data []int
	for range data { // ✅ OK: 0 iterations
	}
}

// Error: creating slice with invalid parameters
// Ошибка: создание слайса с неверными параметрами
func makeSliceErrors() {
	// _ = make([]int, 10, 5) // ❌ Compilation error: len > cap
	
	size := -5
	// _ = make([]int, size)  // 🔥 Panic: negative size
}

// Works: reslicing within capacity
// Работает: реслайсинг в пределах емкости (capacity)
func sliceMoreThanSize() {
	data := make([]int, 2, 6) // len=2, cap=6
	slice1 := data[1:6]       // ✅ OK
}

// Error: panic when attempting to increase capacity beyond original limit
// Ошибка: паника при попытке увеличить емкость через реслайс сверх исходной
func increaseCapacityError() {
	data := make([]int, 0, 10)
	// data = data[:10:100] // 🔥 Panic!
}
```

---

### 🧩 Nil vs Empty Slice (Internal Structure)

| State | Code | `len` | `cap` | `== nil` | Data Address |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Nil** | `var s []int` | 0 | 0 | ✅ `true` | `0x0` |
| **Empty** | `s := []int{}` | 0 | 0 | ❌ `false` | `0x58f740` (static) |

```go
func compareSlices() {
	var nilS []string
	emptyS := []string{}
	
	fmt.Printf("Nil slice:   len=%d nil=%t\n", len(nilS), nilS == nil)
	fmt.Printf("Empty slice: len=%d nil=%t\n", len(emptyS), emptyS == nil)
}
```

> [!IMPORTANT]
> To check for "emptiness" (both nil and `{}`), always use `len(s) == 0`. Internally, `nil` and `[]T{}` differ by the data address, but both have zero length.
