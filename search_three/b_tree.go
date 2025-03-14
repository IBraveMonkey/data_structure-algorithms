package search_three

import (
	"container/list"
	"fmt"
)

/* Бинарное дерево
-	Разновидность деревьев
- У каждого узла не может быть больше 2 потомков(левый и правый)
- Меньше проверок при рукурсивной обходе 0 < n < 2

Сложность - O(log n)
*/

// узел(Нода) дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Функция для создания нового узла
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Добавление нового элемента в BST
func (n *Node) Insert(value int) {
	if value < n.Value { // Идем в левое поддерево
		if n.Left == nil {
			n.Left = NewNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value { // Идем в правое поддерево
		if n.Right == nil {
			n.Right = NewNode(value)
		} else {
			n.Right.Insert(value)
		}
	}
}

// Поиск элемента в BST
func (n *Node) Search(value int) *Node {
	if n == nil || n.Value == value {
		return n // Если узел пустой или найдено совпадениe
	}

	if value < n.Value {
		return n.Left.Search(value)
	}

	return n.Right.Search(value)
}

/*
	Функция обхода в ширину (BFC)

Как работает BFS?
Начинаем с корня.
Сначала обрабатываем весь текущий уровень (слева направо).
Переходим на следующий уровень.
*/
func (n *Node) BFC() {
	if n == nil {
		return
	}

	queue := list.New() // Очередь для хранения узлов
	queue.PushBack(n)   // Добавляем корень в очередь

	for queue.Len() > 0 {
		// Достаем первый элемент из очереди
		element := queue.Front()
		node := element.Value.(*Node)
		queue.Remove(element) // Удаляем обработанный элемент

		fmt.Print(node.Value, " ")

		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

/* Обход в глубину (DFS)
DFS спускается вглубь дерева до самого нижнего узла, затем поднимается обратно. Работает на основе стека (LIFO - последний вошел, первый вышел).
*/
// Обход дерева (симметричный, InOrder - слева -> корень -> справа)
func (n *Node) InOrder() {
	if n == nil {
		return
	}

	n.Left.InOrder()
	fmt.Print(n.Value, " ")
	n.Right.InOrder()
}

// Обход - (узел -> слева -> справа)
func (n *Node) PreOrder() {
	if n == nil {
		return
	}

	fmt.Print(n.Value, " ")
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}

	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Println(n.Value, " ")
}

// Возвращает минимальный узел (самый левый в поддереве)
func (n *Node) minValueNode() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}

	return current
}

func (n *Node) maxValuNode() *Node {
	current := n

	for current.Right != nil {
		current = current.Right
	}

	return current
}

/* УДАЛЕНИЕ
- Удаление узла без потомков → просто удаляем.
- Удаление узла с одним потомком → заменяем потомком.
- Удаление узла с двумя потомками → заменяем минимальным узлом из правого поддерева.
*/

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	// 1. Ищем узел, который нужно удалить
	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else { // Найден узел, который нужно удалить
		//2. Случай узел без потомков(лист)
		if n.Left == nil && n.Right == nil {
			// Обрываем ссылку на найденный узел
			return nil
		}

		// 3. Случай: один потомок
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}

		// 4.Случай: два потомка
		minRight := n.Right.minValueNode()       // Находим минимальный узел в правом поддереве
		n.Value = minRight.Value                 // Копируем его значение в текущий узел
		n.Right = n.Right.Delete(minRight.Value) // Удаляем минимальный узел
	}

	return n
}

// Поиск минимальной глубины бинарного дерева DFS
func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}

	// если нет левого поддерева, идем только вправо
	if root.Left == nil {
		return 1 + MinDepth(root.Right)
	}

	// если нет правого поддерева, идем только влево
	if root.Right == nil {
		return 1 + MinDepth(root.Left)
	}
	// если есть оба поддерева, берем минимальную глубину из двух
	return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
}

// Являются ли 2 дерева одинаковыми
func IsSameTree(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true // Оба пустые, значит одинаковые
	}

	if p == nil || q == nil {
		return false // Один пустой, другой нет - разные
	}

	if p.Value != q.Value {
		return false // Значения не совпадают - разные
	}

	return IsSameTree(p.Left, p.Left) && IsSameTree(p.Right, q.Right)
}
