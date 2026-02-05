package dynamic_programming

import "fmt"

// Example demonstrates dynamic programming problems
func Example() {
	// 1. Fibonacci Numbers
	n := 6
	fmt.Printf("Fib(%d) = %d\n", n, Fib(n))

	// 2. Knapsack Problem
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Printf("Knapsack max value: %d\n", Knapsack(weights, values, capacity))

	// 3. Longest Common Subsequence (LCS)
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	fmt.Printf("LCS('%s', '%s') = %d\n", s1, s2, LongestCommonSubsequence(s1, s2))
}

// Task 1: Fibonacci Numbers
// Find the n-th Fibonacci number. Uses a Bottom-Up approach (tabulation).
func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Task 2: 0/1 Knapsack Problem
// Given weights of items, their values, and the maximum capacity of a knapsack.
// Find the maximum value that can be carried.
func Knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1) // dp[i][w] - max value for i items and weight w

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				// Choose max: take the item or don't take it
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][capacity]
}

// Task 3: Longest Common Subsequence (LCS)
func LongestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
