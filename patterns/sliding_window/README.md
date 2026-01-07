## Template

```go
func slidingwindow(nums[] int) {
  state := make(map[int]int) // choose appropriate data structure
  start, total := 0, 0

  for end := start; end < len(nums); end++ {
    // extend window
    // add nums[end] to the state in O(1) time


    for /*state is invalid*/ {
      // repeatedly contract window until it is valid again
      // remove nums[start] from state in O(1) time
      start++
    }

    // Invariant: state of current window is valid here
    total = max(total, end - start + 1)
  }
}
```
```
```

## When to use it

Consider using the sliding window pattern for questions that involve searching for a continuous subarray/substring in an array or string that satisfies a certain constraint.

If you know the length of the subarray/substring you are looking for, use a fixed-length sliding window. Otherwise, use a variable-length sliding window.
Examples:

Finding the largest substring without repeating characters in a given string (variable-length).

Finding the largest substring containing a single character that can be made by replacing at most k characters in a given string (variable-length).

Finding the largest sum of a subarray of size k without duplicate elements in a given array (fixed-length).


## Reference

- [Sliding Window](https://www.hellointerview.com/learn/code/sliding-window/variable-length)
