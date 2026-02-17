# DynamicProgramming

**Description**: 
Dynamic Programming (DP) is the pinnacle of algorithmic skill. It is a method for solving complex problems that transforms exponential chaos into ordered efficiency. The motto of DP is: "Never compute the same thing twice."

- **How it works internally**: DP breaks a problem down into subproblems that overlap. Instead of solving them repeatedly, we record the answer in a "notebook" (a table or an array).
  - *Memoization (Top-Down)*: We solve recursively, but before computing, we check — is the answer already in the cache?
  - *Tabulation (Bottom-Up)*: We fill a table starting from the smallest subproblems up to the global goal (iteratively).
- **Analogy**: Imagine you are building a wall. To lay the 10th row of bricks, you don't need to re-calculate how you built the first nine — they are already there, serving as a foundation. You simply build upon the result of the previous step.


### Pros and Cons
✅ **Pros**:
1. **Incredible Speed**: Allows solving in seconds problems that naive algorithms would spend years on.
2. **Optimality**: Guarantees finding the absolute best (minimum or maximum) solution.

❌ **Cons**:
1. **Memory Overhead**: To remember all the answers, DP often requires creating large tables (arrays or matrices).
2. **Conceptual Difficulty**: The hardest part of DP is identifying the "state" and the "transition formula" (how the current step depends on previous ones). This requires strong abstract thinking skills.

---


## 🚀 Examples


### Task 1: Climbing Stairs / Fibonacci
Find the number of ways to climb to the $N$-th step if you can hop 1 or 2 steps at a time. This is a classic DP problem with the same logic as the Fibonacci sequence.

```go
func ClimbStairs(n int) int {
    if n <= 2 { return n }
    dp := make([]int, n+1)
    dp[1], dp[2] = 1, 2
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

```javascript
function climbStairs(n) {
  if (n <= 2) return n;
  let prev1 = 1, prev2 = 2; // Space optimization
  for (let i = 3; i <= n; i++) {
    const current = prev1 + prev2;
    prev1 = prev2;
    prev2 = current;
  }
  return prev2;
}
```


### Task 2: Coin Change
Given an integer `amount` and an array of coin `coins` denominations, find the minimum number of coins needed for change.

```go
func CoinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ { dp[i] = amount + 1 }
    dp[0] = 0
    
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            if dp[i-coin]+1 < dp[i] {
                dp[i] = dp[i-coin] + 1
            }
        }
    }
    if dp[amount] > amount { return -1 }
    return dp[amount]
}
```

```javascript
function coinChange(coins, amount) {
  const dp = new Array(amount + 1).fill(amount + 1);
  dp[0] = 0;

  for (const coin of coins) {
    for (let i = coin; i <= amount; i++) {
      dp[i] = Math.min(dp[i], dp[i - coin] + 1);
    }
  }
  return dp[amount] > amount ? -1 : dp[amount];
}
```

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

<!-- QUIZ_START 
[
    {
        "question": "What is the primary motto of Dynamic Programming?",
        "options": ["Divide and conquer", "Never compute the same thing twice", "Live for today", "Step back and retry"],
        "correctIndex": 1
    },
    {
        "question": "What is the main difference between Memoization and Tabulation?",
        "options": ["Memoization is Bottom-Up, while Tabulation is Top-Down", "Memoization is Top-Down (recursive with cache), while Tabulation is Bottom-Up (iterative table-filling)", "Memoization uses more CPU, while Tabulation uses more RAM", "There is no difference"],
        "correctIndex": 1
    },
    {
        "question": "Why is 'identifying the state and transition formula' considered the hardest part of DP?",
        "options": ["It requires knowing all Go keywords", "It involves writing the longest code", "It requires complex abstract thinking to define how steps depend on each other", "It needs a very fast internet connection"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

