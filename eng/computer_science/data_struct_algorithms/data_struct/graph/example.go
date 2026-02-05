package graph

import (
	"fmt"
)

// Example demonstrates the use of a graph
func Example() {
	// Create a graph
	g := NewGraph()

	// Add edges
	// 0 -- 1
	// |    |
	// 2 -- 3
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	fmt.Println("Graph structure:")
	g.Print()

	// Problem: Find if a path exists
	start, end := 0, 3
	exists := ValidPath(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, start, end)
	fmt.Printf("Does a path exist from %d to %d? %t\n", start, end, exists)
}

// Problem: Does a Path Exist in a Graph
// Given n vertices and an array of edges, determine if there is a valid path from source to destination.
// Uses BFS.
func ValidPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// Build an adjacency list (locally, using a slice of slices for int vertices)
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)
	visited[source] = true
	queue := []int{source}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == destination {
			return true
		}

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// Problem: Number of Provinces (Connected Components)
// isConnected[i][j] = 1 if cities i and j are connected.
// Return the number of provinces. Uses DFS.
func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}
