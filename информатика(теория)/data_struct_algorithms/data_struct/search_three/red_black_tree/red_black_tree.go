package red_black_tree

import "fmt"

/* Red-Black Tree (Красно-Черное Дерево)
Это самобалансирующееся бинарное дерево поиска. Оно гарантирует, что высота дерева остается логарифмической даже после множества вставок и удалений

Это обычное бинарное дерево поиска(BST), НО с дополнительным свойством балансировки, обеспечивающий равномерный рост

- Каждый узел хранит один дополнительный бит - цвет(красный или черный)

 Основные свойства красно-черного дерева:
	1️⃣ Каждый узел окрашен либо в КРАСНЫЙ, либо в ЧЕРНЫЙ.
	2️⃣ Корень всегда черный.
	3️⃣ Каждый путь от корня до nil-узла содержит одинаковое количество черных узлов (это называется черная высота).
	4️⃣ Красный узел не может иметь красных потомков (то есть красные узлы идут только через черные).
	5️⃣ Каждый новый узел добавляется КРАСНЫМ.

🔹 Эти правила обеспечивают балансировку, предотвращая перекос.

### Сложность

| Операция | Средняя (O) | Худшая (O) | Пространственная (O) |
|:---|:---:|:---:|:---:|
| Вставка | O(log n) | O(log n) | O(log n) |
| Поиск | O(log n) | O(log n) | O(1) итеративно / O(log n) рекурсивно |
| Удаление | O(log n) | O(log n) | O(log n) |
| Хранение | — | — | O(n) |

\*Гарантирует логарифмическую высоту h < 2 * log2(n + 1).
*/

const (
	RED   = true
	BLACK = false
)

type RBTNode struct {
	Val    int
	Color  bool
	Left   *RBTNode
	Right  *RBTNode
	Parent *RBTNode
}

type RBTree struct {
	Root *RBTNode
}

func NewNode(val int) *RBTNode {
	return &RBTNode{Val: val, Color: RED} // Всегда вставляем красный
}

// Левый поворот
func (t *RBTree) LeftRotate(x *RBTNode) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

// Правый поворот
func (t *RBTree) RightRotate(x *RBTNode) {
	y := x.Left
	x.Left = y.Right

	if y.Right != nil {
		y.Right.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

// Вставка нового узла
func (t *RBTree) Insert(val int) {
	newNode := NewNode(val)
	if t.Root == nil {
		newNode.Color = BLACK // Корень всегда черный
		t.Root = newNode
	} else {
		t.insertRecursive(t.Root, newNode)
	}
	t.fixInsert(newNode) // Балансировка
}

// Рекурсивная вставка в BST
func (t *RBTree) insertRecursive(root, node *RBTNode) {
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			t.insertRecursive(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			t.insertRecursive(root.Right, node)
		}
	}
}

// Балансировка после вставки
func (t *RBTree) fixInsert(node *RBTNode) {
	// TODO: реализация правил балансировки
}

// PrintTree - вывод дерева (упрощенный)
func (t *RBTree) PrintTree(node *RBTNode, indent string) {
	if node != nil {
		t.PrintTree(node.Right, indent+"   ")
		color := "R"
		if node.Color == BLACK {
			color = "B"
		}
		fmt.Println(indent, node.Val, color)
		t.PrintTree(node.Left, indent+"   ")
	}
}
