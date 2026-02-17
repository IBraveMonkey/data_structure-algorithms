# Backtracking

**Description**: 
Backtracking is an algorithmic strategy based on "trial and error." It is used to solve problems where you need to explore many possibilities to find one or all solutions that fit specific criteria.

- **How it works internally**: The algorithm builds a solution step by step. If at any stage it becomes clear that the current path leads to a dead end (does not satisfy the rules), we "step back" (backtrack) to the previous state and try a different option. This is implemented via recursion: we dive deep into a decision tree and "surface" back up upon failure.
- **Analogy**: Imagine you are in a labyrinth. You reach a fork and take the left path. If you hit a wall, you return to the last fork and try the right path. If that also hits a wall, you go back even further. You continue this until you find the exit or explore all possibilities.


### Pros and Cons
✅ **Pros**:
1. **Versatility**: It can solve complex combinatorial problems like Sudoku, the N-Queens problem, or generating all possible subsets of a set.
2. **Efficiency through Pruning**: By "pruning" branches of the decision tree that obviously won't lead to a solution, it avoids checking millions of unnecessary variants compared to a brute-force approach.

❌ **Cons**:
1. **Staggering Complexity**: In the worst case, the runtime grows exponentially (O(2^n) or O(n!)). On large datasets, the algorithm could run indefinitely.
2. **Memory Overhead**: Due to deep recursion, there is always a risk of reaching the stack depth limit (Stack Overflow).

---


## 🚀 Examples


### Task 1: Generate all permutations
Find all possible ways to arrange $N$ numbers.

````carousel
```go
func Permutations(nums []int) [][]int {
    var result [][]int
    backtrack(nums, 0, &result)
    return result
}

func backtrack(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        *result = append(*result, temp)
        return
    }

    for i := start; i < len(nums); i++ {
        nums[start], nums[i] = nums[i], nums[start]
        backtrack(nums, start+1, result)
        nums[start], nums[i] = nums[i], nums[start]
    }
}
```
```javascript
function permutations(nums) {
  const result = [];
  
  const backtrack = (start) => {
    if (start === nums.length) {
      result.push([...nums]);
      return;
    }
    
    for (let i = start; i < nums.length; i++) {
      [nums[start], nums[i]] = [nums[i], nums[start]];
      backtrack(start + 1);
      [nums[start], nums[i]] = [nums[i], nums[start]];
    }
  };

  backtrack(0);
  return result;
}
```

### Task 2: N-Queens
A classic problem: place $N$ queens on an $N \times N$ chessboard so that no two queens threaten each other.

```go
func SolveNQueens(n int) [][]string {
    var result [][]string
    board := make([][]string, n)
    for i := range board {
        board[i] = make([]string, n)
        for j := range board[i] { board[i][j] = "." }
    }
    backtrack(board, 0, n, &result)
    return result
}

func backtrack(board [][]string, row, n int, result *[][]string) {
    if row == n {
        *result = append(*result, construct(board))
        return
    }
    for col := 0; col < n; col++ {
        if isValid(board, row, col, n) {
            board[row][col] = "Q"
            backtrack(board, row+1, n, result)
            board[row][col] = "."
        }
    }
}

func isValid(board [][]string, row, col, n int) bool {
    for i := 0; i < row; i++ {
        if board[i][col] == "Q" { return false }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == "Q" { return false }
    }
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == "Q" { return false }
    }
    return true
}

func construct(board [][]string) []string {
    var res []string
    for _, row := range board {
        line := ""
        for _, char := range row { line += char }
        res = append(res, line)
    }
    return res
}
```
```javascript
function solveNQueens(n) {
  const result = [];
  const board = Array.from({ length: n }, () => Array(n).fill('.'));

  const isValid = (row, col) => {
    for (let i = 0; i < row; i++) {
      if (board[i][col] === 'Q') return false;
    }
    for (let i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
      if (board[i][j] === 'Q') return false;
    }
    for (let i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
      if (board[i][j] === 'Q') return false;
    }
    return true;
  };

  const backtrack = (row) => {
    if (row === n) {
      result.push(board.map(r => r.join('')));
      return;
    }
    for (let col = 0; col < n; col++) {
      if (isValid(row, col)) {
        board[row][col] = 'Q';
        backtrack(row + 1);
        board[row][col] = '.';
      }
    }
  };

  backtrack(0);
  return result;
}
```

<!-- QUIZ_START 
[
    {
        "question": "What is the core principle of the Backtracking algorithm?",
        "options": ["Always move forward until the end", "Trial and error: step back and try another path when a dead end is reached", "Calculate all possible paths simultaneously", "Pick the largest immediate value"],
        "correctIndex": 1
    },
    {
        "question": "What is 'pruning' in the context of Backtracking?",
        "options": ["Deleting the entire code and starting over", "Cutting off branches of the decision tree that cannot lead to a solution", "Sorting the input data", "Increasing the recursion depth limit"],
        "correctIndex": 1
    },
    {
        "question": "Which of the following problems is a classic example of Backtracking?",
        "options": ["Sorting a hand of cards", "Finding the middle of a linked list", "Solving the N-Queens problem or a Sudoku puzzle", "Binary search in a sorted array"],
        "correctIndex": 2
    }
]
QUIZ_END -->

