package counting_sort

/*
Counting Sort

What is it?
It is an integer sorting algorithm that operates without comparing elements to each other.
Instead of comparisons, it "counts" the number of occurrences of each value and uses this information to place integers directly into their correct positions.

Why is it needed?
- To sort data INSTANTLY (O(n)), bypassing the mathematical limit of O(n log n) for comparison-based sorts.
- However, this only applies if the range of the integers (K) is relatively small.

What's the core idea?
- If we know that the number "5" appears 3 times and there are exactly 10 numbers smaller than five, we know for certain that the fives will occupy indices 10, 11, and 12.

When to use?
- When sorting integers (or objects that can be mapped to integer keys).
- When the range of values (K = max-min) does not significantly exceed the number of elements (N). For example, sorting a million grades ranging from 1 to 5.

How does it work?
1. Find the min and max values.
2. Create a count array of size (max-min + 1).
3. Iterate through the original array and increment the counter for each number.
4. (Optional for a stable version) Transform count into a prefix sum array (positional indices).
5. Fill the resulting array.

Complexity:
- Time: O(N + K).
- Space: O(K) (for the count array).
*/

func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// Create the count array
	rangeSize := max - min + 1
	count := make([]int, rangeSize)

	// Count frequencies
	for _, v := range arr {
		count[v-min]++
	}

	// Overwrite the original array (for the simplified version)
	sortedIndex := 0
	for valOffset, frequency := range count {
		val := valOffset + min
		for frequency > 0 {
			arr[sortedIndex] = val
			sortedIndex++
			frequency--
		}
	}

	return arr
}
