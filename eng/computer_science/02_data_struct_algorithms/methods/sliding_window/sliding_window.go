package sliding_window

/*
Sliding Window

What is it?
The sliding window method is an optimization technique for iterating over subarrays or substrings. Instead of recalculating data for every possible subarray from scratch, we maintain a "window" (a range of indices [left, right]) and update the result as the window shifts.

Why is it needed?
- To turn nested loops (O(N^2) or O(N*K)) into a linear pass O(N).
- To efficiently solve problems involving finding substrings or subarrays with a specific sum, length, or uniqueness.

What's the point?
- The window "slides" across the array from left to right.
- The right boundary (right) expands the window by adding new elements.
- The left boundary (left) contracts by removing old elements when the window condition is violated or the window reaches the required size.

When to use?
- Given a linear array or string.
- You need to find the "longest/shortest subarray", "count of substrings", or "average in a window of size K".
- The problem implies a contiguous range of elements.

How does it work?
1. Initialize left = 0, currentSum = 0, answer = ...
2. Loop `right` from 0 to the end:
   - Add arr[right] to the current state.
   - (For variable-sized windows) While the condition is violated (e.g., sum > target):
     - Remove arr[left] from the state.
     - Increment left.
   - (For fixed-sized windows) If right >= k - 1:
     - Update the answer.
     - Remove arr[left] before the next step and increment left.
3. Return the answer.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(N) |
| Space | O(1) / O(K)* |

*If you need to store the window state (e.g., character frequencies in a hash table).
*/
