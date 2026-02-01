package red_black_tree

import "fmt"

// Example демонстрирует использование красно-черного дерева
func Example() {
	rbt := &RBTree{}

	fmt.Println("Вставка элементов: 10, 20, 30, 15")
	rbt.Insert(10)
	rbt.Insert(20)
	rbt.Insert(30)
	rbt.Insert(15)

	fmt.Println("Структура дерева (Val Color):")
	// Примечание: полноценная балансировка еще не реализована в red_black_tree.go
	if rbt.Root != nil {
		rbt.PrintTree(rbt.Root, "")
	} else {
		fmt.Println("Дерево пусто")
	}
}
