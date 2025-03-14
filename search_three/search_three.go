package search_three

/* Дерево - иерархическая структура данных

-Состоит из нод(узлов)
-Узел - хранилище для данных и указателей на след узлы(потомки)
-Лист - узел, у которого нет потомков
-Корневой узел - вершина дерева
-Всегда может быть только один родитель и множество дочерних нод
-Высота дерева - расстояние от корня до самого нижнего элемента
*/

/* Бинарное дерево
-Разновидность деревьев
- У каждого узла не может быть больше 2 потомков(левый и правый)
- Меньше проверок при рукурсивной обходе 0 < n < 2

Сложность - O(log n)
*/

// Обход в ширину(BFT) Breath-First Traversal

// Обход в глубину(DFT) Depth-First Traversal

/*  Двоичная Куча(max-куча)

 */

/*  Очередь с приоритетом

 */

type Node2 struct {
	Value int
	Left  *Node2
	Right *Node2
}

func (tr *Node2) NewNode(value int) *Node2 {
	return &Node2{Value: value}
}

func (tr *Node2) Add(value int) {
	if value < tr.Value {
		if tr.Left == nil {
			tr.Left = tr.NewNode(value)
		} else {
			tr.Left.Add(value)
		}
	} else if value > tr.Value {
		if tr.Right == nil {
			tr.Right = tr.NewNode(value)
		} else {
			tr.Right.Add(value)
		}
	}

}
