package recurse

import "fmt"

// Example demonstrates recursion examples
func Example() {
	// 1. Factorial
	fmt.Printf("5! = %d\n", Factorial(5))

	// 2. Power
	fmt.Printf("2^10 = %d\n", Power(2, 10))

	// 3. Tree Traversal
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Printf("Number of nodes in the tree: %d\n", CountNodes(root))
}

// Task 1: Factorial
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Task 2: Fast exponentiation (recursive)
func Power(x, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / Power(x, -n)
	}
	// Optimization: x^n = (x^(n/2))^2
	half := Power(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return x * half * half
}

// Task 3: Counting nodes (Example of working with trees)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

// Simple Fibonacci (as an example of inefficient recursion without optimization, for educational purposes)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
