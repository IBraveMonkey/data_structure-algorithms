/*
Stack (Стек)

Что это такое?
Стек — это упорядоченная коллекция элементов, работающая по принципу LIFO (Last In, First Out — "последним пришёл, первым ушёл").
Представьте стопку тарелок: вы можете положить новую тарелку только наверх и взять тарелку только сверху.

Зачем это нужно?
- Для отслеживания порядка выполнения операций (например, вызовы функций, рекурсия).
- Для отмены действий (Undo).
- Для парсинга выражений и проверки синтаксиса (например, скобок).

В чём смысл?
- Доступ осуществляется только к одному элементу — верхнему.
- Строгий порядок добавления и удаления.

Когда использовать?
- Когда нужно обработать элементы в обратном порядке их поступления.
- В алгоритмах обхода (DFS).
- Для проверки сбалансированности скобок.

Как работает?
- Push: Добавить элемент на вершину стека.
- Pop: Удалить и вернуть элемент с вершины стека.
- Peek (Top): Посмотреть на верхний элемент без удаления.

### Сложность

| Операция | Временная сложность (O) | Пространственная сложность (O) |
|:---|:---:|:---:|
| Push (вставка) | O(1) | O(1) |
| Pop (удаление) | O(1) | O(1) |
| Peek (просмотр) | O(1) | O(1) |
| Поиск | O(n) | O(1) |
| Хранение | — | O(n) |

Как понять, что задача подходит под Stack?
- Нужно найти "пару" для элемента (например, закрывающую скобку).
- Нужно обрабатывать данные в обратном порядке.
- Задача связана с вложенными структурами.
*/

package stack

// ArrayStack - Реализация стека на основе слайса (более производительная в Go за счет локальности кэша)
type ArrayStack struct {
	Data []interface{}
}

// Push - добавляет элемент в стек
func (stack *ArrayStack) Push(data interface{}) {
	stack.Data = append(stack.Data, data)
}

// Pop - удаляет и возвращает верхний элемент стека
func (stack *ArrayStack) Pop() (interface{}, bool) {
	if len(stack.Data) == 0 {
		return nil, false
	}

	lastIndex := len(stack.Data) - 1
	lastElem := stack.Data[lastIndex]

	stack.Data = stack.Data[0:lastIndex]
	return lastElem, true
}

// Peek - возвращает верхний элемент без удаления
func (stack *ArrayStack) Peek() (interface{}, bool) {
	if len(stack.Data) == 0 {
		return nil, false
	}

	lastIndex := len(stack.Data) - 1
	return stack.Data[lastIndex], true
}

// IsEmpty - проверяет, пуст ли стек
func (stack *ArrayStack) IsEmpty() bool {
	return len(stack.Data) == 0
}

// Size - возвращает размер стека
func (stack *ArrayStack) Size() int {
	return len(stack.Data)
}

// Node - узел для связной реализации стека
type Node struct {
	Value interface{}
	Next  *Node
}

// Stack - Реализация стека на основе связного списка (менее эффективна по памяти из-за указателей)
type Stack struct {
	Top     *Node
	SizeVal int
}

// Push - добавляет элемент в стек
func (s *Stack) Push(value interface{}) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
	s.SizeVal++
}

// Pop - удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	removedValue := s.Top.Value
	s.Top = s.Top.Next
	s.SizeVal--
	return removedValue, true
}

// Peek - возвращает верхний элемент без удаления
func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s.Top.Value, true
}

// IsEmpty - проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// Size - возвращает размер стека
func (s *Stack) Size() int {
	return s.SizeVal
}
