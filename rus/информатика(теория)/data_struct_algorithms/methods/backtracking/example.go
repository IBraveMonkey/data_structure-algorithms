package backtracking

import "fmt"

// Example демонстрирует задачи на backtracking
func Example() {
	// 1. Перестановки
	nums := []int{1, 2, 3}
	perms := Permutations(nums)
	fmt.Printf("Перестановки для %v: %v\n", nums, perms)

	// 2. Комбинации
	k := 2
	combs := Combinations(nums, k)
	fmt.Printf("Комбинации из %v по %d: %v\n", nums, k, combs)

	// 3. N Ферзей
	n := 4
	queens := SolveNQueens(n)
	fmt.Printf("Решения для %d ферзей: \n%v\n", n, queens)
}

// Задача 1: Генерация всех перестановок
func Permutations(nums []int) [][]int {
	var result [][]int
	backtrackPermutation(nums, 0, &result)
	return result
}

func backtrackPermutation(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrackPermutation(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// Задача 2: Комбинации
func Combinations(nums []int, k int) [][]int {
	var result [][]int
	var current []int
	backtrackCombination(nums, k, 0, current, &result)
	return result
}

func backtrackCombination(nums []int, k int, start int, current []int, result *[][]int) {
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrackCombination(nums, k, i+1, current, result)
		current = current[:len(current)-1]
	}
}

// Задача 3: N Ферзей
func SolveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	backtrackQueen(board, 0, n, &result)
	return result
}

func backtrackQueen(board [][]string, row int, n int, result *[][]string) {
	if row == n {
		temp := make([]string, n)
		for i := range board {
			temp[i] = ""
			for j := range board[i] {
				temp[i] += board[i][j]
			}
		}
		*result = append(*result, temp)
		return
	}

	for col := 0; col < n; col++ {
		if isValidPosition(board, row, col, n) {
			board[row][col] = "Q"
			backtrackQueen(board, row+1, n, result)
			board[row][col] = "."
		}
	}
}

func isValidPosition(board [][]string, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
