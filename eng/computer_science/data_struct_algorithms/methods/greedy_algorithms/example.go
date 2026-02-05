package greedy_algorithms

import (
	"fmt"
	"sort"
)

// Example demonstrates greedy algorithms
func Example() {
	// 1. Coin change (greediness works for canonical denomination systems)
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Printf("Min coins for %d with %v: %d\n", amount, coins, CoinChange(coins, amount))

	// 2. Non-overlapping intervals (Activity Selection Problem)
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Printf("Max number of non-overlapping intervals: %d\n", MaxNonOverlapping(intervals))
}

// Task 1: Minimum number of coins (for standard denominations)
// Note: works for (1, 2, 5, 10...), but may not work for (1, 3, 4)
func CoinChange(coins []int, amount int) int {
	count := 0
	remaining := amount

	// Sort coins in descending order to take the largest ones first
	// (they are already sorted in the example, but for universality sort could be added)
	// sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	// Since we use a slice of ints and iterate from the end, it's fine.

	for i := len(coins) - 1; i >= 0; i-- {
		for remaining >= coins[i] {
			remaining -= coins[i]
			count++
		}
	}

	if remaining != 0 {
		return -1
	}
	return count
}

// Task 2: Selecting the maximum number of non-overlapping intervals
// (Equivalent to the activity selection problem)
// Greedy strategy: always choose the interval that ends earliest to leave more space for others.
func MaxNonOverlapping(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// Sort by interval end
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		// If the start of the current interval >= the end of the last chosen -> they don't overlap
		if intervals[i][0] >= lastEnd {
			count++
			lastEnd = intervals[i][1]
		}
	}
	return count
}
