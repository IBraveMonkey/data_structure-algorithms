# 🔗 Связный список

**Описание**: 
Связный список — это динамическая структура данных, состоящая из **узлов**. Каждый узел содержит два поля: сами данные и **ссылку (указатель)** на следующий узел в последовательности. В отличие от массива, элементы списка могут быть разбросаны по разным уголкам оперативной памяти.

- **Как это устроено внутри**: Список начинается с головы (**Head**). Чтобы найти 5-й элемент, компьютер не может просто вычислить адрес, как в массиве. Ему приходится буквально "пройтись" по цепочке: от головы к первому, от первого ко второму и так до нужного. Поэтому доступ по индексу занимает **O(n)**.
- **Аналогия**: Представьте игру-квест "Поиск сокровищ". У вас есть первая записка (голова), в ней написано, где лежит клад и где искать следующую записку. Вы не знаете, где находится финальный приз, пока не пройдете всю цепочку записок по очереди.


### Виды списков
- **Односвязный**: Каждый узел знает только о *следующем*.
- **Двусвязный**: Каждый узел знает и о *следующем*, и о *предыдущем*. Это позволяет эффективно перемещаться в обоих направлениях.
- **Кольцевой**: Последний узел ссылается на первый, замыкая цепь.


### Преимущества и недостатки
✅ **Плюсы**:
1. **Динамический размер**: Список легко растет и уменьшается без необходимости перераспределять огромные блоки памяти.
2. **Быстрая вставка/удаление**: Если у вас уже есть указатель на нужное место, вставка или удаление происходят мгновенно (O(1)), так как нужно просто перекинуть "ссылки", не сдвигая сами данные.
3. **Эффективное использование фрагментированной памяти**: Списку не нужен один цельный блок памяти, он заполняет свободные "дырки".

❌ **Минусы**:
1. **Медленный доступ**: Чтобы добраться до середины или конца, нужно перебрать все элементы с начала.
2. **Лишняя память**: На каждый элемент приходится хранить один или два указателя, что увеличивает расход памяти.
3. **Плохая кэш-дружелюбность**: Узлы разбросаны по памяти, поэтому процессор тратит больше времени на их поиск.

---


### Визуализация

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    Head[Head] --> N1[Node 1<br/>data: 10<br/>next: →]
    N1 --> N2[Node 2<br/>data: 20<br/>next: →]
    N2 --> N3[Node 3<br/>data: 30<br/>next: nil]



linkStyle default stroke:#009688,stroke-width:2px;




```


### Сложность

| Операция | Временная сложность (O) | Пространственная сложность (O) |
|:---|:---:|:---:|
| Доступ (по индексу) | O(n) | O(1) |
| Вставка (в начало) | O(1) | O(1) |
| Вставка (в конец) | O(n)* | O(1) |
| Вставка (в середину) | O(n) | O(1) |
| Удаление (из начала) | O(1) | O(1) |
| Удаление (из конца) | O(n)* | O(1) |
| Удаление (из середины) | O(n) | O(1) |
| Поиск | O(n) | O(1) |
| Хранение | — | O(n) |

> [!NOTE]
> *Для двусвязного списка вставка/удаление в конец — O(1), если хранится указатель на хвост.


---


## 💻 Реализация

```go
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
```

```javascript
class Node {
    constructor(value) {
        this.value = value;
        this.next = null;
    }
}

class LinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
        this.size = 0;
    }

    // Добавляет элемент в начало
    addToFront(value) {
        const newNode = new Node(value);
        newNode.next = this.head;
        this.head = newNode;
        if (!this.tail) this.tail = newNode;
        this.size++;
    }

    // Добавляет элемент в конец
    addToBack(value) {
        const newNode = new Node(value);
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
        } else {
            this.tail.next = newNode;
            this.tail = newNode;
        }
        this.size++;
    }

    // Вставка по индексу
    insertAt(value, index) {
        if (index < 0 || index > this.size) return false;
        if (index === 0) return this.addToFront(value);
        if (index === this.size) return this.addToBack(value);

        let current = this.head;
        for (let i = 0; i < index - 1; i++) {
            current = current.next;
        }
        const newNode = new Node(value);
        newNode.next = current.next;
        current.next = newNode;
        this.size++;
        return true;
    }

    // Удаление из начала
    removeFromFront() {
        if (!this.head) return;
        this.head = this.head.next;
        if (!this.head) this.tail = null;
        this.size--;
    }

    // Удаление из конца
    removeFromBack() {
        if (!this.head) return;
        if (this.head === this.tail) {
            this.head = null;
            this.tail = null;
            this.size--;
            return;
        }
        let current = this.head;
        while (current.next !== this.tail) {
            current = current.next;
        }
        current.next = null;
        this.tail = current;
        this.size--;
    }

    // Поиск значения
    find(value) {
        let current = this.head;
        while (current) {
            if (current.value === value) return current;
            current = current.next;
        }
        return null;
    }

    // Получение по индексу
    get(index) {
        if (index < 0 || index >= this.size) return null;
        let current = this.head;
        for (let i = 0; i < index; i++) {
            current = current.next;
        }
        return current.value;
    }

    // Разворот списка
    reverse() {
        let prev = null;
        let current = this.head;
        this.tail = this.head;
        while (current) {
            let nextTemp = current.next;
            current.next = prev;
            prev = current;
            current = nextTemp;
        }
        this.head = prev;
    }
}
```


## 🚀 Практические задачи

```go
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

// Распространенные задачи
func HasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func Middle(l *LinkedList) *Node {
	if l.Head == nil {
		return nil
	}
	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

```javascript
// Распространенные алгоритмы со списками
function hasCycle(head) {
    if (!head || !head.next) return false;
    let slow = head;
    let fast = head;
    while (fast && fast.next) {
        slow = slow.next;
        fast = fast.next.next;
        if (slow === fast) return true;
    }
    return false;
}

function findMiddle(head) {
    if (!head) return null;
    let slow = head;
    let fast = head;
    while (fast && fast.next) {
        slow = slow.next;
        fast = fast.next.next;
    }
    return slow;
}

function reverseList(head) {
    let prev = null;
    let current = head;
    while (current) {
        let nextTemp = current.next;
        current.next = prev;
        prev = current;
        current = nextTemp;
    }
    return prev;
}
```

<!-- QUIZ_START 
[
    {
        "question": "В чем основное преимущество связного списка перед обычным массивом?",
        "options": ["Мгновенный доступ к любому элементу по индексу", "Динамический размер и эффективная вставка/удаление в начало без сдвига других элементов", "Меньшее потребление памяти", "Лучшая работа с кэшем процессора"],
        "correctIndex": 1
    },
    {
        "question": "Какова временная сложность поиска элемента в односвязном списке в худшем случае?",
        "options": ["O(1)", "O(log n)", "O(n)", "O(n log n)"],
        "correctIndex": 2
    },
    {
        "question": "Что такое 'хвост' (Tail) в контексте связного списка?",
        "options": ["Первый элемент списка", "Указатель на следующий список", "Последний узел списка, поле 'next' которого обычно равно nil", "Лишняя память"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```
