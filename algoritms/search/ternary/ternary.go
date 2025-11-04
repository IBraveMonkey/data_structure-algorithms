package ternary

// TimeComplexity O(log(n))
func TernarySearch(data []int, left, right, target int) int {
	if left >= right {
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3

		if data[m1] == target {
			return m1
		}

		if data[m2] == target {
			return m2
		}

		if target < data[m1] {
			return TernarySearch(data, left, m1-1, target)
		} else if target > data[m2] {
			return TernarySearch(data, m2+1, right, target)
		} else {
			return TernarySearch(data, m1+1, m2-1, target)
		}
	}

	return -1
}
