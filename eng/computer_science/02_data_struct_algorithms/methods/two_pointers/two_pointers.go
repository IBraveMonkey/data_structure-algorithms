package two_pointers

/*
Two Pointers

What is it?
The two pointers method is an algorithmic technique that uses two indices (pointers) moving through a data structure (array, string, list) at different speeds or in different directions.

Why is it needed?
- Time optimization: often leads to an O(N) solution instead of an O(N^2) nested loop.
- Space optimization: allows solving problems "in-place" with O(1) additional memory.

What's the point?
We narrow the search space or process data from two sides or at different speeds, based on data properties (e.g., being sorted).

When to use?
1. Left-Right Pointers:
   - Sorted array (finding a pair with a sum).
   - Palindrome check.
   - Reversing an array.
2. Slow-Fast Pointers (Hare-Tortoise):
   - Cycle detection in a Linked List.
   - Removing duplicates from a sorted array.
   - Finding the middle of a list.
3. Merging two sorted arrays.

How does it work?
- Left-Right: Place pointers at the beginning and end, then move them towards each other based on a condition.
- Slow-Fast: Both start at the beginning. The fast pointer moves 1 or 2 steps, and the slow one moves 1. Use the speed difference.

### Complexity

| Metric | Complexity (O) |
|:---|:---:|
| Time | O(N) |
| Space | O(1) |
*/
