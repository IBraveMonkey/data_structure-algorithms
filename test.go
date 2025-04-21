package main

func Bubble(arr []int) {
	sorted := false

	for !sorted {
		sorted = true

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func Selected(arr []int) {
	n := len(arr) - 1

	for i := 0; i < n; i++ {
		min := i

		for j := i + 1; j <= n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
}

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i

		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
}

func Merge(arr []int) []int {
	if arr == nil {
		return nil
	}

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := []int{}
	right := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] <= arr[mid] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	leftSide := Merge(left)
	rightSide := Merge(right)

	return merge(leftSide, rightSide)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func medianOfThree(arr []int, low, hight int) int {
	mid := (low + hight) / 2

	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}

	if arr[low] > arr[hight] {
		arr[low], arr[hight] = arr[hight], arr[low]

	}

	if arr[mid] > arr[hight] {
		arr[mid], arr[hight] = arr[hight], arr[mid]
	}

	return mid
}

func Quick(arr []int) []int {
	if arr == nil {
		return nil
	}

	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := medianOfThree(arr, 0, len(arr)-1)
	pivot := arr[pivotIndex]
	left := []int{}
	right := []int{}

	for i := 0; i < len(arr); i++ {
		if pivotIndex == i {
			continue
		}

		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	leftSide := Quick(left)
	rightSide := Quick(right)

	return append(append(leftSide, pivot), rightSide...)
}

// func Binary(arr []int, target int) int {
// 	left := 0
// 	right := len(arr)-1
// }
