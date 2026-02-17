# SlidingWindow

**Description**: 
The Sliding Window method is an elegant way to transform a slow brute-force search into a lightning-fast process. It is used wherever you need to analyze contiguous groups of elements in arrays or strings.

- **How it works internally**: Instead of rescanning a subarray for every possible position, we maintain a "window" using two pointers (`left` and `right`). As the window slides to the right, we simply add the new element on the right and "pop out" the old element from the left. This reduces the complexity from **O(n*k)** to a linear time of **O(n)**.
- **Analogy**: Imagine looking out the window of a moving train. You only see a specific section of the landscape at any given time. As the train moves, new trees appear on the right side of the window, and old trees disappear on the left. You don't need to get off the train and re-scan the entire path from the beginning to your current location—you simply update the picture right in front of you.


### Pros and Cons
✅ **Pros**:
1. **Extreme Efficiency**: Reduces repetitive calculations to the absolute minimum.
2. **Low Memory Footprint**: Usually requires only a few variables to track the window's sum or state.

❌ **Cons**:
1. **Contiguous Only**: The method is useless if the elements of your subarray are scattered across the main array.
2. **Complexity of Conditions**: It can sometimes be tricky to correctly tune the exact moment when the left boundary should start "shrinking" (especially in variable-sized windows).

---


## 🚀 Examples


### Task 1: Maximum Average Subarray (Fixed Window)
Find a contiguous subarray whose length is equal to `k` that has the maximum average value.

```go
func FindMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ { sum += nums[i] }
    
    maxSum := sum
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        if sum > maxSum { maxSum = sum }
    }
    return float64(maxSum) / float64(k)
}
```

```javascript
function findMaxAverage(nums, k) {
  let sum = 0;
  for (let i = 0; i < k; i++) sum += nums[i];

  let maxSum = sum;
  for (let i = k; i < nums.length; i++) {
    sum += nums[i] - nums[i - k];
    maxSum = Math.max(maxSum, sum);
  }
  return maxSum / k;
}
```


### Task 2: Minimum Size Subarray Sum (Variable Window)
Find the minimal length of a subarray whose sum is greater than or equal to `target`.

```go
func MinSubArrayLen(target int, nums []int) int {
    left, sum := 0, 0
    minLen := len(nums) + 1

    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            if right-left+1 < minLen {
                minLen = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }
    if minLen > len(nums) { return 0 }
    return minLen
}
```

```javascript
function minSubArrayLen(target, nums) {
  let left = 0, sum = 0, minLen = Infinity;

  for (let right = 0; right < nums.length; right++) {
    sum += nums[right];
    while (sum >= target) {
      minLen = Math.min(minLen, right - left + 1);
      sum -= nums[left++];
    }
  }
  return minLen === Infinity ? 0 : minLen;
}
```


## 🚀 Practical Problems
```go
package sliding_window

import (
	"fmt"
	"math"
)

// Example demonstrates sliding window problems
func Example() {
	// 1. Maximum average in a subarray of length k
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("Max average (k=%d): %.2f\n", k, FindMaxAverage(nums, k))

	// 2. Minimum length of a subarray with sum >= target
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Printf("Min subarray length (sum >= %d): %d\n", target, MinSubArrayLen(target, arr))
}

// Task 1: Find the maximum average value of a subarray of length k
// O(N) time, O(1) space
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	// 1. Initialize the sum of the first window
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	// 2. Slide the window (starting from the k-th element)
	for i := k; i < len(nums); i++ {
		sum += nums[i]   // Add new element (on the right)
		sum -= nums[i-k] // Remove old element (on the left)
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

// Task 2: Minimum subarray size (Variable Size Sliding Window)
// Find the minimum length of a contiguous subarray whose sum is >= target.
// If no such subarray exists, return 0.
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // Expand window to the right

		// Shrink window from the left while the condition is met
		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

<!-- QUIZ_START 
[
    {
        "question": "What is the main goal of the Sliding Window technique?",
        "options": ["To sort an array in O(n log n)", "To reduce complexity from O(n*k) to linear time O(n) for processing contiguous subarrays", "To find the maximum element in a group of scattered items", "To implement an infinite loop safely"],
        "correctIndex": 1
    },
    {
        "question": "How is the 'window' typically maintained as it slides to the right?",
        "options": ["By copying the entire subarray", "By adding the new element on the right and removing the old element from the left", "By clearing the entire array and re-scanning", "By using a stack to store all elements"],
        "correctIndex": 1
    },
    {
        "question": "In which case is the Sliding Window method NOT suitable?",
        "options": ["Finding the maximum sum of a fixed-size subarray", "Analyzing contiguous groups in a string", "Processing elements that are scattered (non-contiguous) across the array", "Finding a minimum length subarray meeting a condition"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

