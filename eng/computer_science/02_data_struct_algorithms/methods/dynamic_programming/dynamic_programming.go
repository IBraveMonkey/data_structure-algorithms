package dynamic_programming

/*
Dynamic Programming

What is it?
Dynamic Programming is a method for solving complex problems by breaking them down into simpler subproblems. The key feature is that the solutions to subproblems are remembered (memoization or tabulation) so that they are not recomputed when overlapping computations occur.

Why is it needed?
- To speed up problem-solving from exponential to polynomial complexity.
- To avoid repeated computations (e.g., as in naive Fibonacci recursion).
- To find a global optimum (maximum/minimum) in optimization problems.

What's the point?
- We divide the problem into small steps (subproblems).
- We save the answers (in an array, hash table, or matrix).
- We build the solution "bottom-up" (tabulation) or "top-down" (memoization) with caching.

When to use?
- There is "optimal substructure" (the solution to a large problem consists of optimal solutions to small ones).
- There are "overlapping subproblems" (the same subproblems are solved multiple times).
- Problems involving finding an extremum (max profit, min path) or the number of variants.

How does it work?
1. Define the state (parameters of the subproblem, e.g., i-th element, weight w).
2. Write the state transition equation (recurrence relation) connecting the current state with previous ones.
3. Set base cases (initial conditions).
4. Fill the table (for the iterative approach) or use recursion with a cache.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(N * M...) |
| Space | O(N * M...) |

*N, M — dimensions of the state space table.
**Space can be optimized to O(N) in some cases.
*/
