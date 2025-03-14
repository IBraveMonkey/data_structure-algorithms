package array

// Приходит отсортированный массив и нам надо найти target/ arr = [1,2,3,4,5,6,7,8,9,10] и target = 8. Можно решить через 2 указателя
func Sum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{arr[left], arr[right]}
		}

		if sum < target {
			left++
			continue
		}

		right--
	}

	return nil
}

// Развернуть массив без аллокаций памяти через два указателя
func ReverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// Cлияние 2 отсортированных массивов(по возрастанию) - два указателя
func MergeTwoArray(arr1, arr2 []int) []int {
	i := 0
	j := 0
	mergedArray := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArray = append(mergedArray, arr1[i])
			i++
		} else {
			mergedArray = append(mergedArray, arr2[j])
			j++
		}
	}

	mergedArray = append(mergedArray, arr1[i:]...)
	mergedArray = append(mergedArray, arr2[j:]...)

	return mergedArray
}
