/*
Graph

What is it?
A Graph is an abstract data structure representing a set of vertices (nodes) and a collection of edges (links) that connect pairs of vertices. Graphs are used to model objects and the relationships between them.

Why is it needed?
- Modeling networks (social networks, internet, road maps).
- Solving routing problems (finding the shortest path).
- Dependency analysis (code compilation, task scheduling).

What's the point?
- Vertices (Nodes/Vertices) represent objects.
- Edges represent relationships between objects.
- Graphs can be directed (digraphs) or undirected, weighted (edges have weights) or unweighted.

When to use?
- When there are objects with complex relationships between them.
- Pathfinding (navigation, games).
- Connectivity analysis (who knows whom, which computers are connected).
- Topological sorting (order of execution for dependent tasks).

How does it work?
- There are two main ways to represent a graph:
  1. Adjacency Matrix: A two-dimensional array where array[i][j] indicates a connection. Efficient for dense graphs.
  2. Adjacency List: An array of lists/maps where each vertex stores a list of its neighbors. Efficient for sparse graphs.

### Complexity

| Operation | Adjacency List (O) | Adjacency Matrix (O) | Space Complexity (O) |
|:---|:---:|:---:|:---:|
| Add Vertex | O(1) | O(V²) | O(V) / O(V²) |
| Add Edge | O(1) | O(1) | O(1) |
| Check Edge | O(V) | O(1) | O(1) |
| List Neighbors | O(degree) | O(V) | O(1) |
| Storage | O(V + E) | O(V²) | — |

*V — number of vertices, E — number of edges.
**degree — number of neighbors for a specific vertex.

How to know if a problem fits Graph?
- The task is about cities and roads, computers and cables, or friends and acquaintances.
- You need to find the shortest path or the number of ways to get from A to B.
- You need to check if objects are connected.
*/

package graph

import "fmt"

// Graph represents a graph using an Adjacency List.
// This is the most versatile way to represent most algorithmic problems.
type Graph struct {
	adjList map[int][]int
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

// AddEdge adds an edge (u, v).
// For an undirected graph, the edge is added in both directions.
func (g *Graph) AddEdge(u, v int) {
	// Add v to u's adjacency list
	g.adjList[u] = append(g.adjList[u], v)
	// Add u to v's adjacency list (since the graph is undirected)
	g.adjList[v] = append(g.adjList[v], u)
}

// Print outputs the graph structure.
func (g *Graph) Print() {
	for node, neighbors := range g.adjList {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}
