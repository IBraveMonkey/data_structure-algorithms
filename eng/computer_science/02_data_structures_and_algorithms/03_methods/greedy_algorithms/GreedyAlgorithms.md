# GreedyAlgorithms

**Description**: 
Greedy Algorithms are a strategy of "living for today." At each step, the algorithm chooses the best available option right now, hoping that the sum of these "best" individual steps will lead to the ideal global outcome.

- **How it works internally**: The algorithm never looks back or reconsiders its past decisions. It takes the locally optimal choice (for example, the largest coin for change or the shortest immediate path) and moves on to the next subproblem.
- **Analogy**: Imagine you are allowed to enter a pastry shop and pick 5 treats. A greedy strategy is to immediately grab the biggest cake, then the biggest eclair, and so forth. This works if your goal is just to maximize weight, but it might fail if you want to assemble the most balanced set of flavors.


### Pros and Cons
✅ **Pros**:
1. **Simplicity and Speed**: These algorithms are very easy to conceptualize and code. They run extremely fast (usually O(n log n) due to sorting).
2. **Quality Approximations**: Even if greediness doesn't provide the perfect solution, it often yields a result that is "good enough" for practical use.

❌ **Cons**:
1. **Lack of Global Vision**: The main pitfall of "greediness" is that it can lead to a dead end. Local success at the start can turn into a disaster by the end.
2. **Limited Applicability**: Before use, one must mathematically prove that the greedy strategy actually works for that specific problem.

---


## 🚀 Examples


### Task 1: Coin Change (Greedy Approach)
Given an integer `amount` and an array of coin `coins` denominations, find the number of coins needed for change by always picking the largest denomination first.
> [!NOTE]
> The greedy approach for coin change only works perfectly for certain denomination sets (like standard US or Russian coins).

```go
func CoinChangeGreedy(coins []int, amount int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(coins)))
    count := 0
    for _, coin := range coins {
        for amount >= coin {
            amount -= coin
            count++
        }
    }
    if amount == 0 { return count }
    return -1
}
```

```javascript
function coinChangeGreedy(coins, amount) {
  coins.sort((a, b) => b - a);
  let count = 0;
  for (const coin of coins) {
    while (amount >= coin) {
      amount -= coin;
      count++;
    }
  }
  return amount === 0 ? count : -1;
}
```


### Task 2: Interval Scheduling
Given a set of processes (intervals) with start and end times, select the maximum number of non-overlapping processes.

```go
type Interval struct {
    Start, End int
}

func MaxNonOverlapping(intervals []Interval) int {
    if len(intervals) == 0 { return 0 }
    
    // Sort by end time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].End < intervals[j].End
    })

    count := 1
    lastEnd := intervals[0].End
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start >= lastEnd {
            count++
            lastEnd = intervals[i].End
        }
    }
    return count
}
```
```javascript
function maxNonOverlapping(intervals) {
  if (intervals.length === 0) return 0;
  
  // Sort by end time
  intervals.sort((a, b) => a[1] - b[1]);
  
  let count = 1;
  let lastEnd = intervals[0][1];
  
  for (let i = 1; i < intervals.length; i++) {
    if (intervals[i][0] >= lastEnd) {
      count++;
      lastEnd = intervals[i][1];
    }
  }
  return count;
}
```

<!-- QUIZ_START 
[
    {
        "question": "What is the defining characteristic of a Greedy Algorithm?",
        "options": ["It reconsiders past decisions to find the best current one", "It makes the locally optimal choice at each step without looking back", "It uses deep recursion to explore all possibilities", "It always sorts data in ascending order"],
        "correctIndex": 1
    },
    {
        "question": "In which scenario might a Greedy strategy fail to find the global optimal solution?",
        "options": ["The Coin Change problem with standard denominations (1, 2, 5)", "The Activity Selection problem", "The Knapsack problem (0/1 version)", "Huffman coding"],
        "correctIndex": 2
    },
    {
        "question": "What is the 'price' often paid for the simplicity and speed of Greedy algorithms?",
        "options": ["High memory usage", "Extremely long code", "Possible lack of global vision, leading to sub-optimal solutions", "Incompatibility with Go"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

