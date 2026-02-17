# Recurse

**Description**: 
Recursion is one of the most elegant and, at the same time, tricky techniques in programming. it occurs when a function calls itself to solve a smaller subproblem.

- **How it works internally**: Every time a function calls itself, a new "layer" — a **stack frame** — is created in RAM. It stores the local variables of that specific call.
  - *Base Case*: This is the "emergency brake" of recursion. It is the condition under which the function stops calling itself and begins returning a result. Without it, the program will keep spawning calls until a `Stack Overflow` occurs.
- **Analogy**: Imagine a Matryoshka (Russian nesting doll). To get to the smallest doll (the result), you must open the largest one, then the one inside, and so on. You cannot reach the center without passing through all the outer layers. Each doll represents a new function call.


### Pros and Cons
✅ **Pros**:
1. **Code Clarity**: Complex algorithms (like tree traversal or Quick Sort) can be written very concisely and clearly.
2. **Natural Fit**: Many problems in mathematics and nature are recursive by essence (fractals, family trees, file systems).

❌ **Cons**:
1. **Memory Consumption**: Every call consumes stack space. If the recursion depth is too high, memory may run out.
2. **Infinite Loop Risk**: An error in the base case turns recursion into an infinite loop that quickly crashes the program.
3. **Overhead**: A function call always involves additional CPU resource costs to create the context.

---


## 🚀 Examples


### Task 1: Factorial
A classic example: $n! = n \times (n-1) \times \dots \times 1$.

```go
func Factorial(n int) int {
    if n <= 1 { return 1 }
    return n * Factorial(n-1)
}
```

```javascript
function factorial(n) {
  if (n <= 1) return 1;
  return n * factorial(n - 1);
}
```


### Task 2: Binary Search (Recursive)
Searching for an element in a sorted array by repeatedly dividing the search interval in half.

```go
func BinarySearchRecursive(nums []int, target, left, right int) int {
    if left > right { return -1 }
    
    mid := left + (right - left) / 2
    if nums[mid] == target { return mid }
    
    if nums[mid] > target {
        return BinarySearchRecursive(nums, target, left, mid - 1)
    }
    return BinarySearchRecursive(nums, target, mid + 1, right)
}
```

```javascript
function binarySearchRecursive(nums, target, left, right) {
  if (left > right) return -1;

  const mid = Math.floor(left + (right - left) / 2);
  if (nums[mid] === target) return mid;

  if (nums[mid] > target) {
    return binarySearchRecursive(nums, target, left, mid - 1);
  }
  return binarySearchRecursive(nums, target, mid + 1, right);
}
```

<!-- QUIZ_START 
[
    {
        "question": "What is a 'stack frame' in the context of recursion?",
        "options": ["A physical frame around the computer screen", "A layer in memory created for each function call to store local variables", "A way to speed up network requests", "A data structure for storing large images"],
        "correctIndex": 1
    },
    {
        "question": "What is the purpose of the 'Base Case' in a recursive function?",
        "options": ["To start the recursion", "To act as an 'emergency brake' and prevent infinite calls", "To double the speed of the algorithm", "To encrypt the return value"],
        "correctIndex": 1
    },
    {
        "question": "What happens if a recursive function lacks a proper base case?",
        "options": ["The code won't compile", "It will result in a 'Stack Overflow' or an infinite loop", "It will run faster", "It will use less memory"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```

