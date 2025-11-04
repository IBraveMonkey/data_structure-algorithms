package stack

// Реализация на массиве
type ArrayStack struct {
	Data []int
}

func (stack *ArrayStack) Push(data int) {
	stack.Data = append(stack.Data, data)
}

func (stack *ArrayStack) Pop() int {
	if len(stack.Data) < 0 {
		return 0
	}

	lastIndex := len(stack.Data) - 1
	lastElem := stack.Data[lastIndex]

	stack.Data = stack.Data[0:lastIndex]
	return lastElem
}

// Реализация на стэке
type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	removedValue := s.Top.Value
	s.Top = s.Top.Next
	return removedValue, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.Top.Value, true
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// Является ли строка палиндромом

type StackString struct {
	Data []string
}

func (q *StackString) Push(el string) {
	q.Data = append(q.Data, el)
}

func (q *StackString) Pop() string {
	if len(q.Data) == 0 {
		return ""
	}

	lastIdx := len(q.Data) - 1
	lastEl := q.Data[lastIdx]

	q.Data = q.Data[0:lastIdx]

	return lastEl
}

func (q *StackString) Size() int {
	return len(q.Data)
}

func (q *StackString) Peek() (string, bool) {
	if len(q.Data) == 0 {
		return "", false
	}

	lastIdx := len(q.Data) - 1
	lastEl := q.Data[lastIdx]

	return lastEl, true
}

// При помощи stack
func IsPalindrome(a string) bool {
	q := StackString{}

	for _, val := range a {
		q.Push(string(val))
	}

	// ["a", "b", "c", "c", "b", "a"]

	for _, val := range a {
		if string(val) != q.Pop() {
			return false
		}

	}

	return true
}

// При помощи 2 указателей
func IsPalindromePointer(a string) bool {
	left := 0
	right := len(a) - 1

	for left < right {
		if a[left] != a[right] {
			return false
		}

		left++
		right--
	}

	return true
}
