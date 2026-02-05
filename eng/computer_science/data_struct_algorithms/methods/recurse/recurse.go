package recurse

/*
Recursion

What is it?
Recursion is a process in which a function calls itself as part of its execution.
It is a powerful tool for solving problems that can be broken down into identical subproblems of smaller size.

Why is it needed?
- To simplify code for complex tasks (e.g., tree and graph traversals).
- To solve problems defined inductively (factorial, sequences).
- To replace complex nested loops.

What's the point?
- Any recursive function consists of two parts:
  1. Base Case: an escape condition when the function stops calling itself. Without it, a stack overflow will occur.
  2. Recursive Step: calling the function with new arguments that bring us closer to the base case.

When to use?
- Data structures are recursive by nature (Trees, DOM, JSON).
- "Divide and Conquer" algorithms (Merge Sort, Quick Sort).
- Dynamic Programming and Backtracking.

How does it work?
With each function call, a new context is created on the call stack. When the base case is reached, the stack begins to "unwind," returning the results back.

Deep recursion can lead to a stack overflow. Some languages have tail call optimization, but standard Go does not.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | Depends on the number of calls |
| Space | O(d) |

*d — maximum recursion depth (call stack depth).
*/
