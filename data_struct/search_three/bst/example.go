package bst

import "fmt"

// Example демонстрирует использование бинарного дерева поиска (BST)
func Example() {
	bst := NewBST()

	// Вставляем значения
	values := []int{5, 3, 8, 2, 4, 7, 9}
	for _, v := range values {
		bst.Insert(v)
	}

	fmt.Println("Обход In-order (отсортированный):", bst.InOrderTraversal())
	fmt.Println("Обход BFS (в ширину):", bst.BFS())
	fmt.Println("Высота дерева:", bst.Height())
	fmt.Println("Валидное ли BST:", bst.IsValidBST())

	// Поиск
	fmt.Println("Поиск 4:", bst.Search(4))
	fmt.Println("Поиск 10:", bst.Search(10))

	// Удаление
	bst.Delete(3)
	fmt.Printf("После удаления 3 (In-order): %v\n", bst.InOrderTraversal())

	// Диапазонная сумма
	fmt.Printf("Сумма в диапазоне [4, 9]: %d\n", bst.RangeSum(4, 9))

	// K-й наименьший элемент
	k := 2
	fmt.Printf("%d-й наименьший элемент: %d\n", k, bst.KthSmallest(k))
}
