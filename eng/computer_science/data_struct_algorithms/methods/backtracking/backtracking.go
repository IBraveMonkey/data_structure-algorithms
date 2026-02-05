package backtracking

/*
Backtracking

What is it?
Backtracking is a method for finding a solution by building candidates step by step and abandoning a candidate ("backtracking") as soon as it becomes clear that this candidate cannot be part of a correct solution. In other words, it is a systematic search through options, cutting off those that clearly do not fit.

Why is it needed?
- To solve combinatorial problems (permutations, combinations).
- For problems where you need to find ALL possible solutions (e.g., N-Queens, Sudoku).
- When a direct search is too expensive, and you need to prune branches of the search tree.

What's the point?
- It's like walking through a maze: go forward as long as you can. If you hit a dead end, go back to the fork and choose another path.
- It uses recursion to go deep into the decision tree.

When to use?
- Problems involving finding all solutions.
- Puzzles (Sudoku, Crosswords).
- Placement problems (N-Queens).
- Generating all possible passwords, subsets.

How does it work?
1. Choose a candidate solution (e.g., place a queen on a square).
2. Check if this is permissible (e.g., is it under attack by others).
3. If permissible, move to the next step (place the next queen).
4. If not permissible or a dead end is reached, undo the choice (remove the queen) and try the next option.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(b^d) / O(N!) |
| Space | O(d) |

*b — branching factor, d — decision tree depth.
**Complexity depends heavily on the specific problem and the efficiency of branch pruning.
*/
