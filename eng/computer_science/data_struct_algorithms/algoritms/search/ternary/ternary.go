package ternary

/*
Ternary Search

What is it?
Similar to binary search, but divides the array into 3 parts instead of 2 using two points, m1 and m2.

Why is it needed?
- Primarily used for finding an extremum (minimum or maximum) in UNIMODAL functions (functions that either increase then decrease, or vice versa).
- For searching in a sorted array, it is usually worse than binary search (more comparisons), so it is used less frequently.

What's the core idea?
- We discard 1/3 of the search area at each step.

When to use?
- Finding the peak of a function f(x) (e.g., a parabola).
- If you need to find an element in a sorted array (though Binary Search is typically better).

How it works (for a function):
1. m1 = left + (right-left)/3
2. m2 = right - (right-left)/3
3. If f(m1) < f(m2) (and searching for a maximum), then the maximum is definitely not in the first third -> left = m1.
4. Otherwise -> right = m2.

How it works (for an array):
1. Compare target with arr[m1] and arr[m2].
2. If a match is found - we're done.
3. If it's less than arr[m1] -> move to the left third.
4. If it's greater than arr[m2] -> move to the right third.
5. Otherwise -> move to the middle third.

Complexity:
- O(log3 n), which is equivalent to O(log n). (The logarithmic base does not affect the asymptotic behavior).
*/

// TernarySearch - search for an element in a sorted array
func TernarySearch(data []int, target int) int {
	left := 0
	right := len(data) - 1

	for left <= right {
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3

		if data[m1] == target {
			return m1
		}
		if data[m2] == target {
			return m2
		}

		if target < data[m1] {
			right = m1 - 1 // Left third
		} else if target > data[m2] {
			left = m2 + 1 // Right third
		} else {
			left = m1 + 1 // Middle third
			right = m2 - 1
		}
	}

	return -1
}
