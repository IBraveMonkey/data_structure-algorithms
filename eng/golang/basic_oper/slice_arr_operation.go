package main

import (
	"fmt"
	"unsafe"
)

// Error: panic on accessing array element with index out of range
func accessToArrayElement1() {
	data := [3]int{1, 2, 3}

	idx := 4               // can hide (trick), but will catch error at runtime
	fmt.Println(data[idx]) // panic

	fmt.Println(data[4]) // compilation error
}

// Error: panic on accessing array element with negative index
func accessToArrayElement2() {
	data := [3]int{1, 2, 3}

	idx := -1              // can hide (trick), but will catch error at runtime, we are not Python by the way (if I'm not mistaken)
	fmt.Println(data[idx]) // panic

	fmt.Println(data[-1]) // compilation error
}

// Works: returns array length
func arrayLen() {
	data := [10]int{}      // created with 10 zeros
	fmt.Println(len(data)) // 10
}

// Works: returns array capacity
func capArray() {
	var data [10]int       // creates with 10 zeros
	fmt.Println(cap(data)) // 10
}

// Works: compares arrays
func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	//	[<, <=, >, >=]  ->  compilation error
}

// Works: returns array size
func emptyArray() {
	var data [10]byte                // created with 10 zeros -> byte = uint8 -> default for uint8 = 0
	fmt.Println(unsafe.Sizeof(data)) // 10

	//data == nil // compilation error
}

// Works: returns array size
func zeroArray() {
	var data [0]int
	fmt.Println(unsafe.Sizeof(data)) // 0
}

// Error: compilation error when creating array with negative length
func negativeArray() {
	var data [-1]int // compilation error
	_ = data
}

// Error: compilation error when creating array with variable length
func arrayCreation() {
	length1 := 100
	var data1 [length1]int // compilation error
	_ = length1
	_ = data1

	const length2 = 100 // this works because we know the value at compile time
	var data2 [length2]int
	_ = data2
}

// Error: compilation error when using make for array (only for slice\map\chan)
func makeArray() {
	_ = make([10]int, 10) // compilation error
}

// Error: only for slice
func appendToArray() {
	_ = append([10]int{}, 10) // compilation error
}

// Error
func accessToSliceElement1() {
	data := make([]int, 3)
	fmt.Println(data[4]) // panic
}

// Error
func accessToSliceElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4]) // panic
}

// Error
func accessToElement3() {
	data := make([]int, 3, 6)
	_ = data[-1] // compilation error
}

// Error
func accessToNilSlice1() {
	var data []int
	_ = data[0] // panic
}

// Error
func accessToNilSlice2() {
	var data []int
	data[0] = 10 // panic
}

// Normal
func appendToNilSlice() {
	var data []int
	data = append(data, 10) // ok
}

// Normal
func rangeByNilSlice() {
	var data []int
	for range data { // ok
		// but accessing is still an error
	}
}

// Works: creates slice of zero length
func makeZeroSlice() {
	data := make([]int, 0)
	fmt.Println(len(data)) // 0
	fmt.Println(cap(data)) // 0
}

// Error: compilation error or panic when creating slice with invalid parameters
func makeSlice() {
	_ = make([]int, -5)    // compilation error
	_ = make([]int, 10, 5) // compilation error

	size := -5
	_ = make([]int, size) // panic

	size = 5
	_ = make([]int, size*2, size) // panic
}

// Works: creates slice with indices within capacity
func sliceMoreThanSize() {
	data := make([]int, 2, 6) // [0,0],0,0,0,0

	slice1 := data[1:6] // ok // [0,0,0,0,0,0]
	_ = slice1
}

// Error: panic or compilation error when creating slice with invalid indices
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

// Works / Error: normal operation or panic when creating slice from nil slice
func sliceWithNilSlice() {
	var data []int

	slice := data[:]  // ok
	slice = data[0:0] // ok
	slice = data[0:1] // panic
	_ = slice
}

// Error: panic when attempting to increase max capacity beyond original
func increaseCapacity() {
	data := make([]int, 0, 10)
	data = data[:10:100] // panic
}
