package main

import (
	"fmt"
	"unsafe"
)

func empty_and_nil_slice() {
	var data []string
	fmt.Println("var data []string:")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// empty=true nil=true size=24 data=0x0; пустой = да, nil = да

	data = []string(nil)
	fmt.Println("data = []string(nil):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// 	empty=true nil=true size=24 data=0x0; пустой = да, nil = да

	data = []string{}
	fmt.Println("data = []string{}:")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// empty=true nil=false size=24 data=0x58f740; пустой = да, nil = нет

	data = make([]string, 0)
	fmt.Println("data = make([]string, 0):")
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data))
	// empty=true nil=false size=24 data=0x58f740; пустой = да, nil = нет

	empty := struct{}{}
	fmt.Println("empty struct address:", unsafe.Pointer(&empty))
	// empty struct address: 0x58f740
}

// лучше проверять вот так на пустоту и nil
func is_empty_arr(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	return false
}
