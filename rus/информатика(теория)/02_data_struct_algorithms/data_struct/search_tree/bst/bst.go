/*
Binary Search Tree (BST) - Двоичное дерево поиска

Что это такое?
BST — это специальная структура данных в виде дерева, где каждый узел имеет не более двух потомков:
левого и правого. Для каждого узла все значения в левом поддереве меньше значения узла,
а в правом — больше. Это свойство позволяет эффективно выполнять операции поиска, вставки и удаления.

Зачем это нужно?
- Быстрый поиск, вставка и удаление элементов (O(log n) в сбалансированном дереве)
- Поддержание упорядоченности данных
- Реализация ассоциативных массивов и множеств

В чём смысл?
- Каждый узел имеет не более двух потомков
- Левое поддерево содержит только значения меньше родителя
- Правое поддерево содержит только значения больше родителя

Когда использовать?
- Когда нужен быстрый доступ к упорядоченным данным
- Для реализации словарей и множеств
- Когда важна операция поиска

Как работает?
- Поиск: аналогично вставке, но возвращаем найденный узел
- Удаление: сложнее, особенно для узлов с двумя потомками

### Сложность

| Операция | Средняя (O) | Худшая (O)* | Пространственная (O) |
|:---|:---:|:---:|:---:|
| Вставка | O(log n) | O(n) | O(h) |
| Поиск | O(log n) | O(n) | O(h) |
| Удаление | O(log n) | O(n) | O(h) |
| Обход | O(n) | O(n) | O(h) |
| Хранение | — | — | O(n) |

\*Худший случай O(n) возникает, когда дерево вырождается в линейный список.
\*\*h — высота дерева. Пространственная сложность обусловлена стеком рекурсии.

Как понять, что задача подходит под BST?
- Нужен быстрый поиск в упорядоченных данных
- Требуется поддержка упорядоченности
- Часто выполняются операции вставки, поиска и удаления
*/

package bst

import (
	"container/list"
)

// Node - узел бинарного дерева поиска
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST - бинарное дерево поиска (обертка над узлами)
type BST struct {
	Root *Node
	Size int
}

// NewBST - создает новое пустое бинарное дерево поиска
func NewBST() *BST {
	return &BST{}
}

// Insert - вставляет значение в дерево
func (bst *BST) Insert(value int) {
	bst.Root = insertRecursive(bst.Root, value)
	bst.Size++
}

func insertRecursive(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}
	// Если value == node.Value, игнорируем дубликат

	return node
}

// Search - ищет значение в дереве
func (bst *BST) Search(value int) bool {
	return searchRecursive(bst.Root, value)
}

func searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return searchRecursive(node.Left, value)
	}

	return searchRecursive(node.Right, value)
}

// Delete - удаляет значение из дерева
func (bst *BST) Delete(value int) {
	bst.Root = deleteRecursive(bst.Root, value)
	bst.Size--
}

func deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// Нашли узел для удаления
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Узел с двумя потомками: находим минимальное значение в правом поддереве
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteRecursive(node.Right, minNode.Value)
	}

	return node
}

// findMin - находит узел с минимальным значением
func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InOrderTraversal - возвращает узлы в отсортированном порядке (лево -> корень -> право)
func (bst *BST) InOrderTraversal() []int {
	result := []int{}
	inOrderHelper(bst.Root, &result)
	return result
}

func inOrderHelper(node *Node, result *[]int) {
	if node != nil {
		inOrderHelper(node.Left, result)
		*result = append(*result, node.Value)
		inOrderHelper(node.Right, result)
	}
}

// BFS - обход дерева в ширину (по уровням)
func (bst *BST) BFS() []int {
	if bst.Root == nil {
		return []int{}
	}

	result := []int{}
	queue := list.New()
	queue.PushBack(bst.Root)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*Node)
		queue.Remove(element)

		result = append(result, node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return result
}

// Height - вычисляет высоту дерева
func (bst *BST) Height() int {
	return heightRecursive(bst.Root)
}

func heightRecursive(node *Node) int {
	if node == nil {
		return -1
	}

	leftHeight := heightRecursive(node.Left)
	rightHeight := heightRecursive(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsValidBST - проверяет, является ли дерево действительным BST
func (bst *BST) IsValidBST() bool {
	// Используем int64 или просто максимальные значения, но здесь упростим с bounds
	// Для int корректнее использовать null pointers для min/max, но int по умолчанию 0
	// Поэтому используем указатели в хелпере
	return isValidBSTHelper(bst.Root, nil, nil)
}

func isValidBSTHelper(node *Node, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
		return false
	}

	val := node.Value
	return isValidBSTHelper(node.Left, min, &val) &&
		isValidBSTHelper(node.Right, &val, max)
}

// RangeSum - сумма значений в заданном диапазоне
func (bst *BST) RangeSum(low, high int) int {
	return rangeSumHelper(bst.Root, low, high)
}

func rangeSumHelper(node *Node, low, high int) int {
	if node == nil {
		return 0
	}

	sum := 0
	if node.Value >= low && node.Value <= high {
		sum += node.Value
	}

	if node.Value > low {
		sum += rangeSumHelper(node.Left, low, high)
	}

	if node.Value < high {
		sum += rangeSumHelper(node.Right, low, high)
	}

	return sum
}

// KthSmallest - находит k-й наименьший элемент (k отсчитывается от 1)
func (bst *BST) KthSmallest(k int) int {
	count := 0
	return kthSmallestHelper(bst.Root, &count, k)
}

func kthSmallestHelper(node *Node, count *int, k int) int {
	if node == nil {
		return -1
	}

	// Обходим левое поддерево
	leftResult := kthSmallestHelper(node.Left, count, k)
	if leftResult != -1 {
		return leftResult
	}

	// Обрабатываем текущий узел
	*count++
	if *count == k {
		return node.Value
	}

	// Обходим правое поддерево
	return kthSmallestHelper(node.Right, count, k)
}
