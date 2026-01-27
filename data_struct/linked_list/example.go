package linked_list

import "fmt"

// Example демонстрирует использование связного списка с различными примерами
func Example() {
	// Создаем новый связный список
	list := &LinkedList{}

	// Добавляем элементы
	list.AddToBack(1)
	list.AddToBack(2)
	list.AddToBack(3)
	list.AddToBack(4)
	list.AddToBack(5)

	fmt.Printf("Исходный список: ")
	list.Print()

	// Поиск элемента
	node := list.Find(3)
	if node != nil {
		fmt.Printf("Найден узел со значением: %d\n", node.Value)
	} else {
		fmt.Println("Узел не найден")
	}

	// Получение элемента по индексу
	value, err := list.Get(2)
	if err == nil {
		fmt.Printf("Значение по индексу 2: %d\n", value)
	}

	// Разворот списка
	list.Reverse()
	fmt.Printf("Развернутый список: ")
	list.Print()

	// Поиск середины
	middle := Middle(list)
	if middle != nil {
		fmt.Printf("Средний элемент: %d\n", middle.Value)
	}
}

// Задача 1: Проверка цикла в списке (Алгоритм Флойда)
// Проверяет, есть ли цикл в списке.
func HasCycle(l *LinkedList) bool {
	if l.Head == nil || l.Head.Next == nil {
		return false
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// Задача 2: Поиск среднего элемента (Алгоритм двух указателей)
// Находит средний элемент списка.
func Middle(l *LinkedList) *Node {
	if l.Head == nil {
		return nil
	}

	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Задача 3: Удалить N-й узел с конца списка
// Дан список, удалить n-й узел с конца и вернуть голову.
func RemoveNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{Value: 0, Next: head}
	fast := dummy
	slow := dummy

	// Сдвигаем fast на n+1 шагов вперед
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Двигаем оба указателя, пока fast не достигнет конца
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Удаляем n-й узел с конца
	slow.Next = slow.Next.Next

	return dummy.Next
}

// Задача 4: Объединить два отсортированных списка
// Объединяет два отсортированных связных списка в один отсортированный.
func MergeTwoLists(list1 *Node, list2 *Node) *Node {
	dummy := &Node{}
	current := dummy

	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Value <= p2.Value {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}

	// Добавляем оставшиеся узлы
	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}

	return dummy.Next
}

// Задача 5: Палиндром связного списка
// Проверить, является ли односвязный список палиндромом.
func IsPalindromeList(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Находим середину списка
	slow, fast := head, head
	var prev *Node

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Разрезаем список
	// Note: в оригинале prev could be nil if list is small but loop condition handles it
	if prev != nil {
		prev.Next = nil // actually splitting helpful for reverse, though not strictly needed depending on reverse logic
	}

	// Разворачиваем вторую половину (начиная со slow)
	// В оригинале код разворота был inline, лучше вынести или оставить как есть
	var secondHalf *Node
	curr := slow
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = secondHalf
		secondHalf = curr
		curr = nextTemp
	}

	// Сравниваем две половины
	firstHalf := head
	p2 := secondHalf

	// Важно: если список нечетной длины, secondHalf может быть длиннее или короче в зависимости от реализации
	// Здесь slow был серединой.
	// Если 1->2->3->2->1. fast stops at 1. slow at 3.
	// prev at 2. split at 2->nil. first: 1->2. second: 3->2->1 reversed -> 1->2->3.
	// Comparison needs care.
	// Let's rely on standard logic: compare until one ends.

	for firstHalf != nil && p2 != nil {
		if firstHalf.Value != p2.Value {
			return false
		}
		firstHalf = firstHalf.Next
		p2 = p2.Next
	}

	return true
}
