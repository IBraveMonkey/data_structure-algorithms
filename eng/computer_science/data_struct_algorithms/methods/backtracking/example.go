package backtracking

import "fmt"

// Example demonstrates backtracking problems
func Example() {
	// 1. Permutations
	nums := []int{1, 2, 3}
	perms := Permutations(nums)
	fmt.Printf("Permutations for %v: %v\n", nums, perms)

	// 2. Combinations
	k := 2
	combs := Combinations(nums, k)
	fmt.Printf("Combinations from %v of size %d: %v\n", nums, k, combs)

	// 3. N-Queens
	n := 4
	queens := SolveNQueens(n)
	fmt.Printf("Solutions for %d queens: \n%v\n", n, queens)
}

// Task 1: Generate all permutations
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

// Task 2: Combinations
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

// Task 3: N-Queens
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
