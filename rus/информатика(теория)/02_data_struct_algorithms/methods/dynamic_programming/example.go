package dynamic_programming

import "fmt"

// Example демонстрирует задачи на динамическое программирование
func Example() {
	// 1. Числа Фибоначчи
	n := 6
	fmt.Printf("Fib(%d) = %d\n", n, Fib(n))

	// 2. Задача о рюкзаке
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Printf("Рюкзак макс. стоимость: %d\n", Knapsack(weights, values, capacity))

	// 3. Самая длинная общая подпоследовательность (LCS)
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	fmt.Printf("LCS('%s', '%s') = %d\n", s1, s2, LongestCommonSubsequence(s1, s2))
}

// Задача 1: Числа Фибоначчи
// Найти n-е число Фибоначчи. Используем Bottom-Up подход (табуляцию).
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

// Задача 2: Задача о рюкзаке (0/1 Knapsack)
// Даны веса предметов, их стоимости и максимальный вес рюкзака.
// Найти максимальную стоимость, которую можно унести.
func Knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1) // dp[i][w] - макс. стоимость для i предметов и веса w

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				// Выбираем максимум: взять предмет или не брать
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][capacity]
}

// Задача 3: Самая длинная общая подпоследовательность (LCS)
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
