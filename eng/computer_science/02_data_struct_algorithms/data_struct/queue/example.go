package queue

import "fmt"

// Alias for LinkedListQueue to be used in examples as "Queue"
type Queue = LinkedListQueue

// Example demonstrates the use of a queue with various examples
func Example() {
	// Create a new queue
	queue := &Queue{}

	// Add elements
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue after adding 1, 2, 3:")

	// View the first element
	if val, ok := queue.Front(); ok {
		fmt.Printf("First element: %d\n", val)
	}

	// Remove elements
	if val, ok := queue.Dequeue(); ok {
		fmt.Printf("Removed element: %d\n", val)
	}

	// Check if the queue is empty
	fmt.Printf("Is queue empty? %t\n", queue.IsEmpty())

	// Example with a bank
	customers := []Customer{{1, 0, 5}, {2, 1, 3}}
	waitingTimes := SimulateBankQueue(customers, 1)
	fmt.Printf("Client waiting times: %v\n", waitingTimes)
}

// Problem 1: Subsequence Check using a Queue
// Check if string a is a subsequence of string b.
// Example: a = "abc", b = "aebdc" => true
func IsSubsequenceWithQueue(a, b string) bool {
	// Create a queue and add characters from string a to it
	queue := &ArrayQueue{}

	for _, char := range a {
		queue.Push(int(char))
	}

	// Traverse string b and remove matching characters from the queue
	for _, char := range b {
		if !queue.IsEmpty() {
			front, _ := queue.Peek()
			if front == int(char) {
				queue.Pop()
			}
		}
	}

	// If the queue is empty, all characters from a were found in b in the correct order
	return queue.IsEmpty()
}

// Problem 2: BFS (Breadth-First Search) for Graph Traversal
// Returns the order in which vertices are visited.
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

		// Add neighbors to the queue
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}

	return result
}

// Problem 3: Bank Queue Simulation
// Models serving customers in the order they arrive
type Customer struct {
	ID          int
	ArrivalTime int
	ServiceTime int
}

func SimulateBankQueue(customers []Customer, tellers int) []int {
	queue := &LinkedListQueue{}
	serviceTimes := make([]int, tellers) // Time each teller becomes free
	waitingTimes := make([]int, len(customers))

	for i, customer := range customers {
		// Add the customer to the queue
		queue.Enqueue(customer)

		// Find the teller who will be free soonest
		minTeller := 0
		for j := 1; j < tellers; j++ {
			if serviceTimes[j] < serviceTimes[minTeller] {
				minTeller = j
			}
		}

		// Customer is served
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

// Problem 4: Recent Counter
// Implement the RecentCounter class, which counts the number of recent requests within the last 3000 ms.
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

	// Remove all pings that are older than 3000 ms
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

// Problem 5: Number of Islands
// Given an m x n grid where '1' represents land and '0' represents water, return the number of islands.
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
				// BFS to mark all land of this island as visited
				q := &Queue{}
				q.Enqueue([2]int{i, j})
				visited[i][j] = true

				for !q.IsEmpty() {
					posInterface, _ := q.Dequeue()
					pos := posInterface.([2]int)
					x, y := pos[0], pos[1]

					// Check neighbors
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

// Problem 6: Open the Lock
// You have a lock with 4 wheels (0-9). The starting position is "0000".
// Given a list of deadends and a target combination, return the minimum number of turns.
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

	for !q.IsEmpty() {
		size := q.Size
		steps++

		for i := 0; i < size; i++ {
			val, _ := q.Dequeue()
			current := val.(string)

			// Generate all possible next states
			for j := 0; j < 4; j++ {
				for d := -1; d <= 1; d += 2 { // d = -1 or 1
					next := []byte(current)
					digit := int(next[j] - '0')
					newDigit := (digit + d + 10) % 10 // Handle 9->0 and 0->9 wrap-around
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
