package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [3]int{10, 20, 30}
	p := unsafe.Pointer(&arr[0])
	p2 := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
	fmt.Println(*(*int)(p2))
}
