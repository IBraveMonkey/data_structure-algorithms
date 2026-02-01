/*
Linked List (Связный список)

Что это такое?
Связный список — это линейная структура данных, в которой элементы (узлы) связаны между собой
последовательно. Каждый узел содержит данные и ссылку (указатель) на следующий узел в списке.
В отличие от массивов, элементы списка не хранятся в непрерывной области памяти.

Зачем это нужно?
- Эффективная вставка и удаление элементов в начало, середину и конец списка (O(1) при наличии ссылки)
- Динамическое изменение размера (не нужно заранее знать размер)
- Экономия памяти при частых изменениях структуры

### Сложность

| Операция | Временная сложность (O) | Пространственная сложность (O) |
|:---|:---:|:---:|
| Доступ (по индексу) | O(n) | O(1) |
| Вставка (в начало) | O(1) | O(1) |
| Вставка (в конец) | O(1)* | O(1) |
| Вставка (в середину) | O(n) | O(1) |
| Удаление (из начала) | O(1) | O(1) |
| Удаление (из конца) | O(n)** | O(1) |
| Удаление (из середины) | O(n) | O(1) |
| Поиск | O(n) | O(1) |
| Хранение | — | O(n) |

\*При наличии указателя на хвост (Tail). Без него — O(n).
\*\*В односвязном списке — O(n), так как нужно найти предпоследний узел. В двусвязном — O(1).

В чём смысл?
- Каждый узел содержит данные и указатель на следующий узел
- Нет необходимости в непрерывной памяти
- Операции вставки/удаления не требуют переноса других элементов

Когда использовать?
- Когда часто вставляются/удаляются элементы
- Когда размер структуры данных часто меняется
- Когда важна экономия памяти при частых изменениях

Как работает?
- Каждый узел содержит данные и указатель на следующий узел
- Для доступа к элементу нужно пройти от начала до нужной позиции
- Для вставки/удаления нужно изменить указатели соседних узлов

Как понять, что задача подходит под Linked List?
- Нужно часто вставлять/удалять элементы
- Не требуется быстрый доступ к произвольному индексу
- Размер данных динамически изменяется

Linked List в Go:

В Go связные списки реализуются с помощью структур и указателей:
- Узел (Node) содержит данные и указатель на следующий узел
- Список (LinkedList) содержит указатели на голову и, возможно, хвост
- Go автоматически управляет сборкой мусора, но нужно следить за циклическими ссылками
- Нет встроенной реализации, нужно реализовывать самостоятельно

Примеры задач с использованием Linked List:
Смотрите файл example.go
*/

package linked_list

import "fmt"

// Node - узел связного списка
type Node struct {
	Value int
	Next  *Node
}

// LinkedList - односвязный список
type LinkedList struct {
	Head *Node
	Tail *Node
	Size int // Добавим размер для удобства
}

// AddToFront - добавляет элемент в начало списка
func (l *LinkedList) AddToFront(value int) {
	newNode := &Node{
		Value: value,
		Next:  l.Head, // Указываем, что следующий элемент - текущая голова
	}

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}

	l.Size++
}

// AddToBack - добавляет элемент в конец списка
func (l *LinkedList) AddToBack(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	// Если список пуст, новый узел становится головой
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Добавляем новый узел после текущего tail
		l.Tail.Next = newNode
		// Обновляем tail чтобы он указывал на новый узел
		l.Tail = newNode
	}

	l.Size++
}

// InsertAt - вставляет элемент в указанную позицию
func (l *LinkedList) InsertAt(value int, index int) error {
	if index < 0 || index > l.Size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		l.AddToFront(value)
		return nil
	}

	if index == l.Size {
		l.AddToBack(value)
		return nil
	}

	// Находим узел перед позицией вставки
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode := &Node{
		Value: value,
		Next:  current.Next,
	}

	current.Next = newNode
	l.Size++

	return nil
}

// RemoveFromFront - удаляет элемент из начала списка
func (l *LinkedList) RemoveFromFront() {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next

	if l.Head == nil {
		l.Tail = nil
	}

	l.Size--
}

// RemoveFromBack - удаляет элемент из конца списка
func (l *LinkedList) RemoveFromBack() {
	if l.Head == nil {
		return
	}

	// Если в списке один элемент
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return
	}

	// Находим предпоследний элемент
	current := l.Head
	for current.Next != l.Tail {
		current = current.Next
	}

	// Удаляем последний элемент
	current.Next = nil
	l.Tail = current

	l.Size--
}

// Find - находит узел с заданным значением
func (l *LinkedList) Find(value int) *Node {
	current := l.Head

	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

// Get - возвращает значение узла по индексу
func (l *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= l.Size {
		return 0, fmt.Errorf("index out of bounds")
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}

// Print - печатает список
func (l *LinkedList) Print() {
	current := l.Head

	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// Reverse - разворачивает список
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.Head
	l.Tail = l.Head // После разворота старая голова станет хвостом

	for current != nil {
		nextTemp := current.Next // Сохраняем следующий узел
		current.Next = prev      // Меняем направление связи
		prev = current           // Перемещаем prev вперед
		current = nextTemp       // Перемещаем current вперед
	}

	l.Head = prev // Новый head - это был последний элемент
}

// RemoveValue - удаляет первое вхождение значения из списка
func (l *LinkedList) RemoveValue(value int) {
	if l.Head == nil {
		return
	}

	// Если удаляемый элемент - голова
	if l.Head.Value == value {
		l.RemoveFromFront()
		return
	}

	// Ищем узел перед удаляемым
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// Если нашли элемент для удаления
	if current.Next != nil {
		// Если удаляемый элемент - хвост, обновляем хвост
		if current.Next == l.Tail {
			l.Tail = current
		}

		current.Next = current.Next.Next
		l.Size--
	}
}
