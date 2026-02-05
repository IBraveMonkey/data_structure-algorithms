package greedy_algorithms

/*
Greedy Algorithms

What is it?
A greedy algorithm is an approach to solving optimization problems where, at each step, a locally optimal choice is made with the hope that these steps will lead to a globally optimal solution.

Why is it needed?
- Fast solution of optimization problems (finding minimum/maximum).
- When a brute force or DP approach is too slow, and a greedy approach provides an exact (or good approximate) solution.

What's the point?
- We don't look into the future and don't backtrack (unlike DP and Backtracking).
- We make a choice "here and now" based on the current situation.

When to use?
- When the problem has the "greedy choice property": a local optimum leads to a global one.
- The problem has "optimal substructure" (like DP).
- Examples: shortest path algorithms (Dijkstra), Huffman coding, constructing a spanning tree (Prim, Kruskal), scheduling problems.

How does it work?
1. Sort the data (greediness is often based on order: the cheapest, the shortest, the earliest, etc.).
2. Iterate and choose an element if it improves the solution and is permissible.
3. Repeat until the end.

Greediness does not work everywhere. A simple example of failure is coin change with arbitrary denominations (e.g., change for 6 with coins 4, 3, 1: greedy (4+1+1) gives 3 coins, while the optimum (3+3) gives 2 coins).

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(N log N)* / O(N)** |
| Space | O(1) / O(N) |

*O(N log N) if preliminary sorting is required.
**O(N) if the data is already sorted or sorting is not needed.
*/
