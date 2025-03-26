package linked_list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Добавление элемента в начало списка
func (l *LinkedList) AddToFront(value int) {
	// 1-2-3-4
	newNode := &Node{
		Value: value,
		Next:  l.Head, // Указываем, что след элемент - текущая голова
	}

	l.Head = newNode
}

// Добавление элемента в конец списка
func (l *LinkedList) AddToBack(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	// Если список пут, новый узел становится головой
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// Идем до конца списка, чтобы добавить элемент
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	// У последнего узла указываем следующий элемент
	current.Next = newNode
}

// Печать списка
func (l *LinkedList) Print() {
	current := l.Head

	if current != nil {
		fmt.Print(current.Value, "->")
		current = current.Next
	}

	fmt.Println("nil")
}

// Нахождение элемента
func (l *LinkedList) FindElem(el int) *Node {
	current := l.Head

	for current != nil {
		if current.Value == el {
			return current
		}

		current = current.Next
	}

	return nil
}
