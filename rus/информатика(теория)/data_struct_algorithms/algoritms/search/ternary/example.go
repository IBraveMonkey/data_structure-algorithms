package ternary

import "fmt"

func Example() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	// Используем итеративную версию, так как она безопаснее для стека
	result := TernarySearch(data, target)

	if result != -1 {
		fmt.Printf("Элемент %d найден на позиции %d (Ternary Search)\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
