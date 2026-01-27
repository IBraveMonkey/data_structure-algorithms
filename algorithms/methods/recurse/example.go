package recurse

import "fmt"

// Example демонстрирует примеры рекурсии
func Example() {
	// 1. Факториал
	fmt.Printf("5! = %d\n", Factorial(5))

	// 2. Степень
	fmt.Printf("2^10 = %d\n", Power(2, 10))

	// 3. Обход дерева
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
	fmt.Printf("Количество узлов в дереве: %d\n", CountNodes(root))
}

// Задача 1: Факториал
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Задача 2: Быстрое возведение в степень (рекурсивно)
func Power(x, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / Power(x, -n)
	}
	// Оптимизация: x^n = (x^(n/2))^2
	half := Power(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return x * half * half
}

// Задача 3: Подсчет узлов (Пример работы с деревьями)
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

// Простой Фибоначчи (как пример плохой рекурсии без оптимизации, для обучения)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
