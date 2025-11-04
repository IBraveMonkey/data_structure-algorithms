package queue

import (
	"fmt"
)

type ArrayQueue struct {
	Data []int
}

func (queue *ArrayQueue) Push(data int) {
	queue.Data = append(queue.Data, data)
}

func (queue *ArrayQueue) Pop() int {
	if len(queue.Data) <= 0 {
		return 0
	}

	firstElem := queue.Data[0]
	queue.Data = queue.Data[1:]
	return firstElem
}

// Очередь на LinkedList
type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

// Добавление элемента в очередь(в конец)
func (q *Queue) Enqueue(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil, // У последнего узла всегда next = nil
	}

	// Если очередь пуста, новый узел становится и головой и хвостом
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
		return
	}

	// Иначе добавляем новый узел в конец(tail)
	q.Tail.Next = newNode
	q.Tail = newNode
}

// Удаление элемента из очереди(из начала)
func (q *Queue) Dequeue() (int, bool) {
	if q.Head == nil {
		return 0, false
	}

	value := q.Head.Value
	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	return value, true
}

func (q *Queue) Peek() (int, bool) {
	if q.Head == nil {
		return 0, false
	}

	return q.Head.Value, true
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue) Print() {
	current := q.Head

	for current != nil {
		fmt.Println(current.Value, "->")
		current = current.Next
	}

	fmt.Println("nil")
}

// Является ли одна строка исходной для другой строки a = abc, b - adbec => true

type StringQueue struct {
	Data []string
}

func (q *StringQueue) Push(val string) {
	q.Data = append(q.Data, val)
}

func (q *StringQueue) Pop() string {
	if len(q.Data) <= 0 {
		return ""
	}

	firstElem := q.Data[0]
	q.Data = q.Data[1:]

	return firstElem
}

func (q *StringQueue) Peek() string {
	if len(q.Data) <= 0 {
		return ""
	}

	return q.Data[0]
}

func (q *StringQueue) Size() int {
	return len(q.Data)
}

// Через Очередь
func IsSubsequence(a, b string) bool {
	queue := &StringQueue{}

	for _, val := range a {
		queue.Push(string(val))
	}

	for _, val := range b {
		firstEl := queue.Peek()

		if firstEl == string(val) {
			queue.Pop()
		}
	}

	return queue.Size() == 0
}

// Через два указателя
func IsSubsequencePointer(a, b string) bool {
	i := 0
	j := 0

	// abc      asbqcd
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	return i == len(a)
}
