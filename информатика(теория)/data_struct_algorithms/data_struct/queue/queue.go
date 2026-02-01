/*
Queue (Очередь)

Что это такое?
Очередь — это линейная структура данных, которая следует принципу FIFO (First In, First Out) —
первым пришел, первым ушел. Элементы добавляются в конец очереди и удаляются из начала очереди.
Это похоже на обычную очередь в магазине: первый в очереди обслуживается первым.

Зачем это нужно?
- Управление задачами, которые должны выполняться в порядке поступления
- Алгоритмы обхода графов (BFS)
- Буферизация данных между процессами
- Обработка запросов в веб-серверах

В чём смысл?
- Добавление элементов в конец (enqueue)
- Удаление элементов из начала (dequeue)
- Поддержание порядка поступления элементов

Когда использовать?
- Когда нужно обрабатывать элементы в порядке их поступления
- Для алгоритмов, требующих обработки в ширину (BFS)
- При симуляции систем массового обслуживания

Как работает?
- Удаление элемента (dequeue) происходит из начала очереди
- Доступ к первому элементу (peek) без удаления

### Сложность

| Операция | Временная сложность (O) | Пространственная сложность (O) |
|:---|:---:|:---:|
| Enqueue (вставка) | O(1) | O(1) |
| Dequeue (удаление) | O(1) | O(1) |
| Peek (просмотр) | O(1) | O(1) |
| Поиск | O(n) | O(1) |
| Хранение | — | O(n) |

Как понять, что задача подходит под Queue?
- Нужно обрабатывать элементы в порядке их поступления
- Требуется алгоритм обхода в ширину
- Есть задачи на симуляцию процессов с очередью

Queue в Go:

В Go очереди можно реализовать несколькими способами:
- С использованием срезов (slайсов) - простая реализация, но менее эффективная для dequeue
- С использованием связных списков - более эффективная реализация
- С использованием контейнера container/list - стандартная реализация двусвязного списка

Примеры задач с использованием Queue:
Смотрите файл example.go
*/

package queue

import "fmt"

// ArrayQueue - очередь, реализованная на основе среза
type ArrayQueue struct {
	Data []interface{}
	Size int
}

// Push - добавляет элемент в конец очереди
func (queue *ArrayQueue) Push(data interface{}) {
	queue.Data = append(queue.Data, data)
	queue.Size++
}

// Pop - удаляет и возвращает элемент из начала очереди
func (queue *ArrayQueue) Pop() (interface{}, bool) {
	if queue.Size <= 0 {
		return nil, false
	}

	firstElem := queue.Data[0]
	queue.Data = queue.Data[1:]
	queue.Size--

	return firstElem, true
}

// Peek - возвращает первый элемент без удаления
func (queue *ArrayQueue) Peek() (interface{}, bool) {
	if queue.Size <= 0 {
		return nil, false
	}

	return queue.Data[0], true
}

// IsEmpty - проверяет, пуста ли очередь
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.Size == 0
}

// QueueNode - узел для связной реализации очереди
type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// LinkedListQueue - очередь, реализованная на основе связного списка
type LinkedListQueue struct {
	Head *QueueNode // Начало очереди
	Tail *QueueNode // Конец очереди
	Size int
}

// Enqueue - добавляет элемент в конец очереди
func (q *LinkedListQueue) Enqueue(value interface{}) {
	newNode := &QueueNode{
		Value: value,
		Next:  nil,
	}

	// Если очередь пуста, новый узел становится и головой и хвостом
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		// Иначе добавляем новый узел в конец (после tail)
		q.Tail.Next = newNode
		q.Tail = newNode
	}

	q.Size++
}

// Dequeue - удаляет и возвращает элемент из начала очереди
func (q *LinkedListQueue) Dequeue() (interface{}, bool) {
	if q.Head == nil {
		return nil, false
	}

	value := q.Head.Value
	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	q.Size--

	return value, true
}

// Front - возвращает первый элемент без удаления
func (q *LinkedListQueue) Front() (interface{}, bool) {
	if q.Head == nil {
		return nil, false
	}

	return q.Head.Value, true
}

// IsEmpty - проверяет, пуста ли очередь
func (q *LinkedListQueue) IsEmpty() bool {
	return q.Head == nil
}

// Print - печатает содержимое очереди
func (q *LinkedListQueue) Print() {
	current := q.Head

	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// CircularQueue - кольцевая очередь (Ring Buffer)
type CircularQueue struct {
	data     []int
	front    int
	rear     int
	size     int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:     make([]int, capacity),
		front:    -1,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

func (cq *CircularQueue) Enqueue(value int) bool {
	if cq.size == cq.capacity {
		return false // Очередь полна
	}

	if cq.front == -1 { // Первая вставка
		cq.front = 0
		cq.rear = 0
	} else {
		cq.rear = (cq.rear + 1) % cq.capacity
	}

	cq.data[cq.rear] = value
	cq.size++
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.size == 0 {
		return 0, false // Очередь пуста
	}

	value := cq.data[cq.front]

	if cq.front == cq.rear { // Последний элемент
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.capacity
	}

	cq.size--
	return value, true
}

func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}
