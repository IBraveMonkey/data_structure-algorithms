package queue

import "fmt"

// Alias for LinkedListQueue to be used in examples as "Queue"
type Queue = LinkedListQueue

// Example демонстрирует использование очереди на различных примерах
func Example() {
	// Создаем новую очередь
	queue := &Queue{}

	// Добавляем элементы
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Очередь после добавления 1, 2, 3:")

	// Смотрим первый элемент
	if val, ok := queue.Front(); ok {
		fmt.Printf("Первый элемент: %d\n", val)
	}

	// Удаляем элементы
	if val, ok := queue.Dequeue(); ok {
		fmt.Printf("Удаленный элемент: %d\n", val)
	}

	// Проверяем пуста ли очередь
	fmt.Printf("Очередь пуста? %t\n", queue.IsEmpty())

	// Пример с банком
	customers := []Customer{{1, 0, 5}, {2, 1, 3}}
	waitingTimes := SimulateBankQueue(customers, 1)
	fmt.Printf("Время ожидания клиентов: %v\n", waitingTimes)
}

// Задача 1: Проверка подпоследовательности с использованием очереди
// Проверить, является ли строка a подпоследовательностью строки b
// Пример: a = "abc", b = "aebdc" => true
func IsSubsequenceWithQueue(a, b string) bool {
	// Создаем очередь и добавляем туда символы из строки a
	queue := &ArrayQueue{}

	for _, char := range a {
		queue.Push(int(char))
	}

	// Проходим по строке b и удаляем совпадающие символы из очереди
	for _, char := range b {
		if !queue.IsEmpty() {
			front, _ := queue.Peek()
			if front == int(char) {
				queue.Pop()
			}
		}
	}

	// Если очередь пуста, значит все символы из a были найдены в b в правильном порядке
	return queue.IsEmpty()
}

// Задача 2: BFS (Breadth-First Search) для обхода графа
// Возвращает порядок посещения вершин
func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := &LinkedListQueue{}
	result := []int{}

	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		nodeInterface, _ := queue.Dequeue()
		node := nodeInterface.(int)
		result = append(result, node)

		// Добавляем соседей в очередь
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}

	return result
}

// Задача 3: Симуляция очереди в банке
// Моделирует обслуживание клиентов в порядке поступления
type Customer struct {
	ID          int
	ArrivalTime int
	ServiceTime int
}

func SimulateBankQueue(customers []Customer, tellers int) []int {
	queue := &LinkedListQueue{}
	serviceTimes := make([]int, tellers) // Время освобождения каждого кассира
	waitingTimes := make([]int, len(customers))

	for i, customer := range customers {
		// Добавляем клиента в очередь
		queue.Enqueue(customer)

		// Находим кассира, который освободится раньше
		minTeller := 0
		for j := 1; j < tellers; j++ {
			if serviceTimes[j] < serviceTimes[minTeller] {
				minTeller = j
			}
		}

		// Клиент обслуживается
		startTime := maxInt(serviceTimes[minTeller], customer.ArrivalTime)
		waitingTimes[i] = startTime - customer.ArrivalTime
		serviceTimes[minTeller] = startTime + customer.ServiceTime
	}

	return waitingTimes
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Задача 4: Счетчик недавних запросов (Recent Counter)
// Реализовать класс RecentCounter, который считает количество запросов за последние 3000 мс.
type RecentCounter struct {
	queue *Queue
}

func NewRecentCounter() *RecentCounter {
	return &RecentCounter{
		queue: &Queue{},
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue.Enqueue(t)

	// Удаляем все пинги, которые старше 3000 мс
	for !rc.queue.IsEmpty() {
		front, _ := rc.queue.Front()
		if front.(int) < t-3000 {
			rc.queue.Dequeue()
		} else {
			break
		}
	}

	return rc.queue.Size
}

// Задача 5: Количество островов
// Дана сетка m x n, где '1' - суша, '0' - вода. Вернуть количество островов.
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				// BFS чтобы пометить всю сушу этого острова как посещенную
				q := &Queue{}
				q.Enqueue([2]int{i, j})
				visited[i][j] = true

				for !q.IsEmpty() {
					posInterface, _ := q.Dequeue()
					pos := posInterface.([2]int)
					x, y := pos[0], pos[1]

					// Проверяем соседей
					for _, dir := range directions {
						nx, ny := x+dir[0], y+dir[1]

						if nx >= 0 && nx < m && ny >= 0 && ny < n &&
							grid[nx][ny] == '1' && !visited[nx][ny] {
							visited[nx][ny] = true
							q.Enqueue([2]int{nx, ny})
						}
					}
				}
			}
		}
	}

	return count
}

// Задача 6: Открыть замок
// У вас есть замок с 4 колесами (0-9). Начальная позиция "0000".
// Даны список тупиков и целевая комбинация. Вернуть минимальное количество поворотов.
func OpenLock(deadends []string, target string) int {
	deadendSet := make(map[string]bool)
	for _, de := range deadends {
		deadendSet[de] = true
	}

	if deadendSet["0000"] {
		return -1
	}

	if target == "0000" {
		return 0
	}

	q := &Queue{}
	q.Enqueue("0000")
	visited := map[string]bool{"0000": true}

	steps := 0

	// Используем nil как маркер уровня или просто считаем размер
	// В данном случае считаем по уровням (BFS)

	for !q.IsEmpty() {
		size := q.Size // Используем поле Size напрямую
		steps++

		for i := 0; i < size; i++ {
			val, _ := q.Dequeue()
			current := val.(string)

			// Генерируем все возможные следующие состояния
			for j := 0; j < 4; j++ {
				for d := -1; d <= 1; d += 2 { // d = -1 или 1
					next := []byte(current)
					digit := int(next[j] - '0')
					newDigit := (digit + d + 10) % 10 // Обработка перехода 9->0 и 0->9
					next[j] = byte('0' + newDigit)
					nextStr := string(next)

					if nextStr == target {
						return steps
					}

					if !visited[nextStr] && !deadendSet[nextStr] {
						visited[nextStr] = true
						q.Enqueue(nextStr)
					}
				}
			}
		}
	}

	return -1
}
