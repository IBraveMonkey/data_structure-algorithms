# TwoPointers

**Description**: 
The Two Pointers method is an elegant technique that allows solving complex problems with a single pass through the data. Instead of using nested loops, we launch two indices ("pointers") moving toward each other or at different speeds.

- **How it works internally**: The core idea is narrowing the search space. Because the pointers move in a directed fashion, we can discard irrelevant options without checking them.
  - *Opposite Pointers*: Start at the ends of the array and move toward the center (ideal for checking palindromes or finding pairs in a sorted array).
  - *Fast and Slow (Hare and Tortoise)*: One moves faster than the other. Used for finding cycles in linked lists or pinpointing the middle of an array.
- **Analogy**: Imagine you and a friend are looking for each other on a long street. If you both start from opposite ends and walk toward each other, you will meet much faster than if one of you stood still while the other checked every single house.


### Pros and Cons
✅ **Pros**:
1. **Maximum Speed**: Transforms slow O(n²) algorithms into fast **O(n)** ones.
2. **Memory Efficiency**: Operates directly on the existing array, requiring no copies (O(1) space complexity).
3. **Implementation Simplicity**: Once the pointer logic is understood, the resulting code is typically very compact.

❌ **Cons**:
1. **Order Dependency**: The method often requires the data to be sorted. If it isn't, you must spend time sorting it first.
2. **Specialized Application**: Primarily suitable for linear structures (arrays, strings, linked lists).

---


## 🚀 Examples


### Task 1: Two Sum (Sorted Array)
Find two numbers in a sorted array that add up to `target`.

```go
func TwoSumSorted(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1} // Typically returns 1-based indices
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return []int{-1, -1}
}
```

```javascript
function twoSumSorted(numbers, target) {
  let left = 0, right = numbers.length - 1;
  while (left < right) {
    const sum = numbers[left] + numbers[right];
    if (sum === target) return [left + 1, right + 1];
    sum < target ? left++ : right--;
  }
  return [-1, -1];
}
```


### Task 2: Valid Palindrome
Check if a string is a palindrome, ignoring case and non-alphanumeric characters.

```go
import (
    "strings"
    "unicode"
)

func IsPalindrome(s string) bool {
    runes := []rune(strings.ToLower(s))
    left, right := 0, len(runes)-1
    
    for left < right {
        if !unicode.IsLetter(runes[left]) && !unicode.IsNumber(runes[left]) {
            left++
            continue
        }
        if !unicode.IsLetter(runes[right]) && !unicode.IsNumber(runes[right]) {
            right--
            continue
        }
        if runes[left] != runes[right] { return false }
        left++
        right--
    }
    return true
}
```

```javascript
function isPalindrome(s) {
  const cleaned = s.toLowerCase().replace(/[^a-z0-9]/g, '');
  let left = 0, right = cleaned.length - 1;
  while (left < right) {
    if (cleaned[left++] !== cleaned[right--]) return false;
  }
  return true;
}
```


## 🚀 Practical Problems
```go
package two_pointers

import (
	"fmt"
	"strings"
)

// Example demonstrates two pointers problems
func Example() {
	// 1. Two Sum (for sorted array)
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("TwoSum indices for %v, target %d: %v\n", nums, target, TwoSumSorted(nums, target))

	// 2. Palindrome check
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("IsPalindrome ('%s'): %v\n", s, IsPalindrome(s))

	// 3. Duplicate removal
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	lenAfter := RemoveDuplicates(arr)
	fmt.Printf("RemoveDuplicates - new len: %d, arr prefix: %v\n", lenAfter, arr[:lenAfter])
}

// Task 1: Two Sum (Input Array Is Sorted)
// Find two numbers that add up to a target. Return their indices (0-based in this implementation).
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++ // Sum is too small, we need more -> move left pointer to the right
		} else {
			right-- // Sum is too large, we need less -> move right pointer to the left
		}
	}

	return []int{}
}

// Task 2: Palindrome Validation
func IsPalindrome(s string) bool {
	// In a real task, it's better to handle runes and not create a new string for O(1) space,
	// but we'll simplify string preparation for this example.
	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r + 32 // toLower
		}
		return -1
	}, s)

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// Task 3: Removing duplicates from a sorted array
// Returns the new length (k).
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// If a new unique element is found
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// Task 4: In-Place Array Reversal
func ReverseArray(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

<!-- QUIZ_START 
[
    {
        "question": "What are the common types of Two Pointers mentioned in the description?",
        "options": ["Top and Bottom pointers", "Recursive and Iterative pointers", "Opposite Pointers (moving toward each other) and Fast/Slow Pointers", "Random and Static pointers"],
        "correctIndex": 2
    },
    {
        "question": "What is a major advantage of using Two Pointers for problems like Palindrome validation?",
        "options": ["It uses more memory but is safer", "It can transform O(n²) logic into O(n) with O(1) space complexity", "It allows working with unsorted data more efficiently than any other method", "It simplifies code by removing the need for 'for' loops"],
        "correctIndex": 1
    },
    {
        "question": "In what scenario are Fast and Slow pointers typically used?",
        "options": ["Checking for a palindrome", "Finding the sum of two numbers in a sorted array", "Finding cycles in linked lists or the middle of an array", "Sorting an array using bubble sort"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

