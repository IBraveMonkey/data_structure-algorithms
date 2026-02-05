package stack

import "fmt"

// Example demonstrates the use of a stack with various examples
func Example() {
	// Create a new stack
	stack := &Stack{}

	// Add elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack after adding 1, 2, 3:")

	// View the top element
	if val, ok := stack.Peek(); ok {
		fmt.Printf("Top element: %d\n", val)
	}

	// Remove elements
	if val, ok := stack.Pop(); ok {
		fmt.Printf("Removed element: %d\n", val)
	}

	// Check if the stack is empty
	fmt.Printf("Is stack empty? %t\n", stack.IsEmpty())

	// Palindrome test
	testStr := "racecar"
	fmt.Printf("Is '%s' a palindrome? %t\n", testStr, IsPalindrome(testStr))

	testStr2 := "hello"
	fmt.Printf("Is '%s' a palindrome? %t\n", testStr2, IsPalindrome(testStr2))
}

// Problem 1: Valid Parentheses
// Determine if a string with various types of brackets is valid.
func IsValidParentheses(s string) bool {
	stack := &Stack{}

	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack.Push(char) // save as rune
		case ')', '}', ']':
			topInterface, ok := stack.Peek()
			if !ok {
				return false
			}
			top := topInterface.(rune)

			if top != mapping[char] {
				return false
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

// Problem 2: Reverse Polish Notation (RPN)
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
func EvalRPN(tokens []string) int {
	stack := &Stack{}

	for _, token := range tokens {
		switch token {
		case "+":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) + val2.(int))
		case "-":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) - val2.(int))
		case "*":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) * val2.(int))
		case "/":
			val2, _ := stack.Pop()
			val1, _ := stack.Pop()
			stack.Push(val1.(int) / val2.(int))
		default:
			// Token is a number
			var num int
			fmt.Sscanf(token, "%d", &num)
			stack.Push(num)
		}
	}

	result, _ := stack.Pop()
	return result.(int)
}

// Problem 3: Daily Temperatures
// For each day, return the number of days until it gets warmer.
func DailyTemperatures(temperatures []int) []int {
	stack := &Stack{} // Stack stores indices
	result := make([]int, len(temperatures))

	for i, temp := range temperatures {
		// As long as the stack is not empty and the current temperature is higher than the temperature at the index in the stack
		for !stack.IsEmpty() {
			topIndexInter, _ := stack.Peek()
			topIndex := topIndexInter.(int)

			if temp <= temperatures[topIndex] {
				break
			}
			// Extract the index and calculate the difference
			stack.Pop()
			result[topIndex] = i - topIndex
		}
		// Add the current index
		stack.Push(i)
	}

	return result
}

// Problem 4: Palindrome Check (via Stack)
func IsPalindrome(a string) bool {
	q := &Stack{}

	for _, val := range a {
		q.Push(string(val))
	}

	for _, val := range a {
		popVal, _ := q.Pop()
		if string(val) != popVal.(string) {
			return false
		}
	}

	return true
}

// Problem 5: Palindrome Check (via 2 Pointers)
func IsPalindromePointer(a string) bool {
	left := 0
	right := len(a) - 1

	for left < right {
		if a[left] != a[right] {
			return false
		}

		left++
		right--
	}

	return true
}
