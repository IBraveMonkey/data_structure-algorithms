# Кэш LRU

**Описание**: 
LRU Cache (Least Recently Used) — это одна из самых популярных стратегий кэширования, которая помогает программам "не забывать" важное и вовремя избавляться от ненужного. Ее главный принцип: "Если мы долго чем-то не пользовались, скорее всего, оно нам не понадобится и в будущем".

- **Как это устроено внутри**: LRU — это тандем двух структур. **Хеш-таблица** обеспечивает мгновенный поиск данных по ключу (O(1)), а **двусвязный список** хранит элементы в порядке их использования. 
  - Когда вы читаете или записываете элемент, он перемещается в "голову" списка (самый свежий).
  - Элементы, к которым давно не обращались, постепенно "сползают" к хвосту.
  - Как только кэш заполняется, элемент из самого хвоста (самый "старый") удаляется навсегда.
- **Аналогия**: Представьте ваш рабочий стол. Самые нужные документы лежат прямо перед вами. Те, что вы достали вчера, лежат чуть дальше. Если стол заполнится полностью, вы возьмете ту бумажку, которую не трогали дольше всех, и выбросите ее в корзину, чтобы освободить место для новой.


### Преимущества и недостатки
✅ **Плюсы**:
1. **Фиксированная память**: Вы точно знаете, сколько памяти занимает кэш, и он никогда не разрастется сверх меры.
2. **Высокая скорость**: И поиск, и вставка, и удаление работают за **O(1)**.
3. **Предсказуемость**: Отлично работает в ситуациях, когда одни и те же данные запрашиваются часто (принцип локальности).

❌ **Минусы**:
1. **Сложность реализации**: Нужно аккуратно управлять связями в списке и мапой одновременно.
2. **Расход памяти на служебные структуры**: Для хранения указателей (Next/Prev) в списке требуется больше памяти, чем для обычного массива.

---


### Сложность

| Операция | Временная сложность (O) | Пространственная сложность (O) |
|:---|:---:|:---:|
| Get (получение) | O(1) | O(1) |
| Put (вставка/обновление) | O(1) | O(1) |
| Удаление | O(1) | O(1) |
| Хранение | — | O(n) |

\*O(1) достигается за счёт комбинации хеш-таблицы (для поиска) и двусвязного списка (для поддержания порядка).

Как понять, что задача подходит под LRU Cache?
- Нужно ограничить размер кэша
- Нужно автоматически удалять "старые" данные
- Важна частота использования элементов

LRU Cache:

Реализация LRU кэша обычно сочетает:
- Двусвязный список для отслеживания порядка использования
- Хеш-таблицу (map) для быстрого доступа к элементам
- Синхронизацию при многопоточном доступе
- Управление памятью для эффективной работы

Примеры задач с использованием LRU Cache:
Смотрите файл example.go

Node - узел двусвязного списка для LRU кэша
LRUCache - реализация LRU кэша с использованием двусвязного списка и хеш-таблицы
Constructor - создает новый LRU кэш с заданной вместимостью
Создаем виртуальные голову и хвост для упрощения операций
remove - удаляет узел из двусвязного списка
addToHead - добавляет узел в начало списка (как недавно использованный)
Get - возвращает значение по ключу, помечая элемент как недавно использованный
Перемещаем узел в начало списка (как недавно использованный)
Put - добавляет или обновляет элемент в кэше
Обновляем существующий элемент
Перемещаем в начало как недавно использованный
Проверяем, не превышена ли вместимость
Удаляем наименее недавно использованный элемент (в конце списка)
Создаем новый узел
Добавляем в начало списка и в хеш-таблицу


## 💻 Реализация

```go
package lru

// Node - узел двусвязного списка для LRU кэша
type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

// LRUCache - реализация LRU кэша с использованием двусвязного списка и хеш-таблицы
type LRUCache struct {
	capacity int
	cache    map[int]*Node // Ключ -> Узел
	head     *Node         // Виртуальная голова (самый недавно использованный)
	tail     *Node         // Виртуальный хвост (наименее недавно использованный)
}

// New создает новый LRU кэш с заданной вместимостью
func New(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{}, // Виртуальная голова
		tail:     &Node{}, // Виртуальный хвост
	}
	// Связываем виртуальные узлы
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

// remove - удаляет узел из двусвязного списка
func (lru *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// addToHead - добавляет узел в начало списка (как недавно использованный)
func (lru *LRUCache) addToHead(node *Node) {
	node.Next = lru.head.Next
	node.Prev = lru.head
	lru.head.Next.Prev = node
	lru.head.Next = node
}

// Get - возвращает значение по ключу, помечая элемент как недавно использованный
func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if !exists {
		return -1
	}
	
	// Перемещаем узел в начало списка (как недавно использованный)
	lru.remove(node)
	lru.addToHead(node)
	
	return node.Value
}

// Put - добавляет или обновляет элемент в кэше
func (lru *LRUCache) Put(key int, value int) {
	node, exists := lru.cache[key]
	
	if exists {
		// Обновляем существующий элемент
		node.Value = value
		// Перемещаем в начало как недавно использованный
		lru.remove(node)
		lru.addToHead(node)
	} else {
		// Создаем новый узел
		newNode := &Node{Key: key, Value: value}
		
		// Проверяем, не превышена ли вместимость
		if len(lru.cache) >= lru.capacity {
			// Удаляем наименее недавно использованный элемент (в конце списка)
			lruNode := lru.tail.Prev
			lru.remove(lruNode)
			delete(lru.cache, lruNode.Key)
		}
		
		// Добавляем в начало списка и в хеш-таблицу
		lru.addToHead(newNode)
		lru.cache[key] = newNode
	}
}

// Size возвращает текущий размер кэша
func (lru *LRUCache) Size() int {
	return len(lru.cache)
}

// Clear очищает весь кэш
func (lru *LRUCache) Clear() {
	lru.cache = make(map[int]*Node)
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
}

```

```javascript
// В JS Map сохраняет порядок вставки.
// Чтобы реализовать LRU, мы можем удалять и снова вставлять элемент,
// чтобы он переместился в конец (как самый свежий).
class LRUCache {
  constructor(capacity) {
    this.capacity = capacity;
    this.cache = new Map();
  }

  get(key) {
    if (!this.cache.has(key)) return -1;

    const value = this.cache.get(key);
    // Перемещаем в конец (делаем самым свежим)
    this.cache.delete(key);
    this.cache.set(key, value);
    return value;
  }

  put(key, value) {
    if (this.cache.has(key)) {
      // Если ключ уже существует, удаляем его
      this.cache.delete(key);
    } else if (this.cache.size >= this.capacity) {
      // Удаляем первое добавленное (самое старое)
      const firstKey = this.cache.keys().next().value;
      this.cache.delete(firstKey);
    }
    // Добавляем элемент (он будет в конце)
    this.cache.set(key, value);
  }

  size() {
    return this.cache.size;
  }

  clear() {
    this.cache.clear();
  }
}

// Альтернативная реализация с doubly-linked list для более явного контроля
class Node {
  constructor(key, value) {
    this.key = key;
    this.value = value;
    this.prev = null;
    this.next = null;
  }
}

class LRUCacheWithList {
  constructor(capacity) {
    this.capacity = capacity;
    this.cache = new Map(); // key -> Node
    this.head = new Node(0, 0); // Виртуальная голова
    this.tail = new Node(0, 0); // Виртуальный хвост
    this.head.next = this.tail;
    this.tail.prev = this.head;
  }

  _remove(node) {
    node.prev.next = node.next;
    node.next.prev = node.prev;
  }

  _addToHead(node) {
    node.next = this.head.next;
    node.prev = this.head;
    this.head.next.prev = node;
    this.head.next = node;
  }

  get(key) {
    if (!this.cache.has(key)) return -1;
    
    const node = this.cache.get(key);
    // Перемещаем в начало (как недавно использованный)
    this._remove(node);
    this._addToHead(node);
    
    return node.value;
  }

  put(key, value) {
    if (this.cache.has(key)) {
      const node = this.cache.get(key);
      node.value = value;
      this._remove(node);
      this._addToHead(node);
    } else {
      const newNode = new Node(key, value);
      
      if (this.cache.size >= this.capacity) {
        // Удаляем наименее недавно использованный (у хвоста)
        const lruNode = this.tail.prev;
        this._remove(lruNode);
        this.cache.delete(lruNode.key);
      }
      
      this._addToHead(newNode);
      this.cache.set(key, newNode);
    }
  }

  size() {
    return this.cache.size;
  }

  clear() {
    this.cache.clear();
    this.head.next = this.tail;
    this.tail.prev = this.head;
  }
}

// Пример использования
const lru = new LRUCache(2);
lru.put(1, 1);
lru.put(2, 2);
console.log(lru.get(1));     // 1
lru.put(3, 3);               // Вытесняет ключ 2
console.log(lru.get(2));     // -1 (не найдено)

```


## 🚀 Практические задачи
```go
package lru

import "fmt"

// Example демонстрирует использование LRU кэша
func Example() {
	// Создаем LRU кэш с вместимостью 2
	cache := New(2)

	// Добавляем элементы
	cache.Put(1, 1) // кэш: [1=1]
	fmt.Println("Put(1, 1)")

	cache.Put(2, 2) // кэш: [2=2, 1=1]
	fmt.Println("Put(2, 2)")

	// Получаем элемент
	val := cache.Get(1)
	fmt.Printf("Get(1): %d (ожидается 1)\n", val) // кэш: [1=1, 2=2]

	// Добавляем еще один элемент (приведет к удалению наименее используемого - 2)
	cache.Put(3, 3) // кэш: [3=3, 1=1]
	fmt.Println("Put(3, 3) - вытесняет ключ 2")

	val = cache.Get(2)
	fmt.Printf("Get(2): %d (ожидается -1)\n", val) // 2 был удален

	cache.Put(4, 4) // кэш: [4=4, 3=3] - вытесняет 1
	fmt.Println("Put(4, 4) - вытесняет ключ 1")

	val = cache.Get(1)
	fmt.Printf("Get(1): %d (ожидается -1)\n", val)

	val = cache.Get(3)
	fmt.Printf("Get(3): %d (ожидается 3)\n", val)

	val = cache.Get(4)
	fmt.Printf("Get(4): %d (ожидается 4)\n", val)
}

<!-- QUIZ_START 
[
    {
        "question": "На каком главном принципе основана стратегия вытеснения в LRU кэше?",
        "options": ["Удаляется самый старый по времени добавления элемент", "Удаляется элемент, к которому дольше всего не обращались (Least Recently Used)", "Удаляется самый большой по размеру элемент", "Удаляется случайный элемент"],
        "correctIndex": 1
    },
    {
        "question": "Комбинация каких двух структур данных позволяет реализовать LRU кэш с эффективностью O(1) для основных операций?",
        "options": ["Стек и Очередь", "Хэш-таблица и Двусвязный список", "Массив и Дерево поиска", "Граф и Куча"],
        "correctIndex": 1
    },
    {
        "question": "Что происходит с элементом в LRU кэше при обращении к нему через Get()?",
        "options": ["Он удаляется из кэша", "Он перемещается в начало (голову) списка как недавно использованный", "Его значение обнуляется", "Ничего не происходит"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```
