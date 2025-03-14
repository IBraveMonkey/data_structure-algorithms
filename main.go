package main

import (
	"fmt"
)

func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// 1, 8
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	count := make([]int, max-(max+1))
	for _, v := range arr {
		count[v-min]++
	}

	sortedIndex := 0
	for i, c := range count {
		for c > 0 {
			arr[sortedIndex] = i + min
			sortedIndex++
			c--
		}
	}

	return arr
}

func main() {
	fmt.Println("Hello")
	unsorted := []int{4, 2, 2, 8, 3, 3, 1}
	sorted := CountingSort(unsorted)
	fmt.Println(sorted)
}
