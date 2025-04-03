package main

import (
	"fmt"

	"github.com/IBraveMonkey/data-structure-algorithms/quick_sort"
)

const (
	Sunday = iota
	Monday
	Tuesday
)

func main() {
	arr := []int{8, 0, 10, 4, 15, 2, 24}

	fmt.Println(quick_sort.QuickSort2(arr))

	// Работа с арифметикой указателей
	// arr := [3]int{10, 20, 30}
	// p := unsafe.Pointer(&arr[0])
	// p2 := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
	// fmt.Println(*(*int)(p2))
}
