# LeetCode Interview Patterns Guide

## 1. Two Pointers

### What is it?
Two pointers is a technique where you use two indices to traverse a data structure, typically an array or string. The pointers can move:
- Toward each other (opposite ends)
- Same direction (one slow, one fast)
- From different starting points

```
Opposite Direction:
[1, 2, 3, 4, 5, 6]
 ↑              ↑
left          right

Same Direction:
[1, 2, 3, 4, 5, 6]
 ↑  ↑
slow fast
```

### When to use?
- Array is sorted (or can be sorted)
- Need to find pairs/triplets with a target sum
- Remove duplicates in-place
- Partition problems
- Palindrome checks

### Sample Problems
- **Easy**: Valid Palindrome (125), Remove Duplicates from Sorted Array (26)
- **Medium**: 3Sum (15), Container With Most Water (11), Sort Colors (75)

### Template Code (Go)

```go
// Opposite direction - finding pairs
func twoSumSorted(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    
    for left < right {
        sum := nums[left] + nums[right]
        
        if sum == target {
            return []int{left, right}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    
    return []int{-1, -1}
}

// Same direction - remove duplicates
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    
    return slow + 1
}
```

---

## 2. Sliding Window

### What is it?
A technique to process subarrays/substrings by maintaining a window that slides through the data structure. The window expands to explore new elements and contracts to maintain constraints.

```
Fixed Size Window:
[1, 2, 3, 4, 5, 6], k=3
 [------]           window at position 0
    [------]        window slides right
       [------]     window slides right

Dynamic Size Window:
[1, 2, 3, 4, 5, 6], sum <= 10
 [------]           expand until constraint violated
    [--]            contract to satisfy constraint
```

### When to use?
- Contiguous subarray/substring problems
- Finding maximum/minimum sum of subarray of size K
- Longest/shortest substring with certain properties
- Problems with "window" in description
- When you need to optimize from O(n²) to O(n)

### Sample Problems
- **Easy**: Maximum Average Subarray I (643), Contains Duplicate II (219)
- **Medium**: Longest Substring Without Repeating Characters (3), Minimum Size Subarray Sum (209), Longest Repeating Character Replacement (424)

### Template Code (Go)

```go
// Fixed size window
func maxSumSubarray(nums []int, k int) int {
    windowSum, maxSum := 0, 0
    
    // Build first window
    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }
    maxSum = windowSum
    
    // Slide window
    for i := k; i < len(nums); i++ {
        windowSum = windowSum - nums[i-k] + nums[i]
        maxSum = max(maxSum, windowSum)
    }
    
    return maxSum
}

// Dynamic size window
func longestSubstringKDistinct(s string, k int) int {
    left, maxLen := 0, 0
    charCount := make(map[byte]int)
    
    for right := 0; right < len(s); right++ {
        // Expand window
        charCount[s[right]]++
        
        // Contract window if constraint violated
        for len(charCount) > k {
            charCount[s[left]]--
            if charCount[s[left]] == 0 {
                delete(charCount, s[left])
            }
            left++
        }
        
        maxLen = max(maxLen, right-left+1)
    }
    
    return maxLen
}
```

---

## 3. Fast & Slow Pointers (Floyd's Cycle Detection)

### What is it?
Two pointers moving at different speeds through a linked list. The slow pointer moves one step at a time, while the fast pointer moves two steps.

```
Cycle Detection:
1 → 2 → 3 → 4
    ↑       ↓
    6 ← 5 ←─┘

slow: 1 → 2 → 3 → 4 → 5 → 6 → 3
fast: 1 → 3 → 5 → 3 (cycle detected when slow == fast)
```

### When to use?
- Detect cycles in linked lists
- Find middle of linked list
- Find nth node from end
- Check if linked list is palindrome
- Find start of cycle

### Sample Problems
- **Easy**: Linked List Cycle (141), Middle of the Linked List (876), Happy Number (202)
- **Medium**: Linked List Cycle II (142), Find the Duplicate Number (287), Reorder List (143)

### Template Code (Go)

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

// Detect cycle
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if slow == fast {
            return true
        }
    }
    
    return false
}

// Find middle
func findMiddle(head *ListNode) *ListNode {
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow
}

// Find cycle start
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    // Phase 1: Detect cycle
    slow, fast := head, head
    hasCycle := false
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if slow == fast {
            hasCycle = true
            break
        }
    }
    
    if !hasCycle {
        return nil
    }
    
    // Phase 2: Find cycle start
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}
```

---

## 4. Tree DFS (Depth-First Search)

### What is it?
Explores a tree by going as deep as possible before backtracking. Can be implemented recursively or with a stack. Three main traversal types:
- **Preorder**: Root → Left → Right
- **Inorder**: Left → Root → Right
- **Postorder**: Left → Right → Root

```
        1
       / \
      2   3
     / \
    4   5

Preorder:  1, 2, 4, 5, 3
Inorder:   4, 2, 5, 1, 3
Postorder: 4, 5, 2, 3, 1
```

### When to use?
- Tree traversal problems
- Path sum problems
- Validate BST
- Find height/depth
- Serialize/deserialize tree
- Count nodes
- Check if trees are same/symmetric

### Sample Problems
- **Easy**: Maximum Depth of Binary Tree (104), Invert Binary Tree (226), Path Sum (112), Same Tree (100)
- **Medium**: Validate Binary Search Tree (98), Binary Tree Right Side View (199), Lowest Common Ancestor (236), Diameter of Binary Tree (543)

### Template Code (Go)

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Recursive DFS - most common
func dfsRecursive(root *TreeNode) {
    if root == nil {
        return
    }
    
    // Preorder: process root first
    // process(root)
    
    dfsRecursive(root.Left)
    
    // Inorder: process root in middle
    // process(root)
    
    dfsRecursive(root.Right)
    
    // Postorder: process root last
    // process(root)
}

// Path sum example
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    
    newSum := targetSum - root.Val
    return hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum)
}

// Max depth example
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    
    return 1 + max(leftDepth, rightDepth)
}

// Iterative DFS with stack
func dfsIterative(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    result := []int{}
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        result = append(result, node.Val)
        
        // Push right first so left is processed first
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
    
    return result
}
```

---

## 5. Tree BFS (Breadth-First Search / Level Order)

### What is it?
Explores a tree level by level using a queue. Processes all nodes at the current level before moving to the next level.

```
        1
       / \
      2   3
     / \   \
    4   5   6

Level 0: [1]
Level 1: [2, 3]
Level 2: [4, 5, 6]

BFS Order: 1, 2, 3, 4, 5, 6
```

### When to use?
- Level order traversal
- Find minimum depth
- Connect nodes at same level
- Zigzag level order
- Binary tree right side view
- Average of levels
- Any "level by level" problem

### Sample Problems
- **Easy**: Binary Tree Level Order Traversal (102), Average of Levels (637), Minimum Depth of Binary Tree (111)
- **Medium**: Binary Tree Zigzag Level Order Traversal (103), Binary Tree Right Side View (199), Populating Next Right Pointers (116)

### Template Code (Go)

```go
// Basic BFS
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    result := [][]int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []int{}
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            currentLevel = append(currentLevel, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, currentLevel)
    }
    
    return result
}

// Right side view
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    result := []int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            // Last node in level
            if i == levelSize-1 {
                result = append(result, node.Val)
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return result
}

// Zigzag level order
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    result := [][]int{}
    queue := []*TreeNode{root}
    leftToRight := true
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            index := i
            if !leftToRight {
                index = levelSize - 1 - i
            }
            currentLevel[index] = node.Val
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, currentLevel)
        leftToRight = !leftToRight
    }
    
    return result
}
```

---

## 6. Modified Binary Search

### What is it?
Binary search on sorted arrays or search spaces where you can eliminate half the possibilities in each iteration. Works on both explicit arrays and abstract search spaces.

```
Searching for target = 7
[1, 3, 5, 7, 9, 11, 13]
 l        m          r   mid=7, found!

Rotated array: [4, 5, 6, 7, 0, 1, 2]
                l     m        r   which half is sorted?
```

### When to use?
- Searching in sorted array
- Finding element in rotated sorted array
- Finding peak element
- Search in 2D matrix
- Finding minimum in rotated array
- Any "find X in O(log n)" problem

### Sample Problems
- **Easy**: Binary Search (704), Search Insert Position (35), First Bad Version (278)
- **Medium**: Search in Rotated Sorted Array (33), Find Peak Element (162), Search a 2D Matrix (74), Find Minimum in Rotated Sorted Array (153)

### Template Code (Go)

```go
// Standard binary search
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

// Find first occurrence
func findFirst(nums []int, target int) int {
    left, right := 0, len(nums)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            result = mid
            right = mid - 1 // Continue searching left
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}

// Search in rotated sorted array
func searchRotated(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        }
        
        // Determine which half is sorted
        if nums[left] <= nums[mid] {
            // Left half is sorted
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right half is sorted
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}
```

---

## 7. Top K Elements (Heap)

### What is it?
Uses a heap (priority queue) to efficiently find the K largest or smallest elements. A min-heap of size K keeps the K largest elements, while a max-heap of size K keeps the K smallest.

```
Find top 3 largest: [3, 2, 1, 5, 6, 4]

Min-heap of size 3:
Add 3: [3]
Add 2: [2, 3]
Add 1: [1, 2, 3]
Add 5: [2, 3, 5] (removed 1)
Add 6: [3, 5, 6] (removed 2)
Add 4: [4, 5, 6] (removed 3)

Result: [4, 5, 6]
```

### When to use?
- Find K largest/smallest elements
- K closest points
- Top K frequent elements
- Kth largest element
- Merge K sorted lists
- Problems mentioning "top K", "largest K", "smallest K"

### Sample Problems
- **Easy**: Kth Largest Element in a Stream (703), Last Stone Weight (1046)
- **Medium**: Kth Largest Element in an Array (215), Top K Frequent Elements (347), K Closest Points to Origin (973), Kth Smallest Element in a Sorted Matrix (378)

### Template Code (Go)

```go
import "container/heap"

// Min heap for top K largest
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Find K largest elements
func findKLargest(nums []int, k int) []int {
    h := &MinHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = (*h)[i]
    }
    
    return result
}

// Top K frequent elements
func topKFrequent(nums []int, k int) []int {
    // Count frequencies
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }
    
    // Use heap to find top K
    type pair struct {
        num   int
        count int
    }
    
    h := make([]pair, 0)
    for num, count := range freq {
        h = append(h, pair{num, count})
        if len(h) > k {
            // Remove element with min frequency
            minIdx := 0
            for i := 1; i < len(h); i++ {
                if h[i].count < h[minIdx].count {
                    minIdx = i
                }
            }
            h = append(h[:minIdx], h[minIdx+1:]...)
        }
    }
    
    result := make([]int, len(h))
    for i, p := range h {
        result[i] = p.num
    }
    
    return result
}
```

---

## 8. Dynamic Programming

### What is it?
Breaks down complex problems into simpler subproblems, stores results to avoid recomputation. Two main approaches:
- **Top-down**: Recursion + Memoization
- **Bottom-up**: Iterative with table

```
Fibonacci example:
fib(5) = fib(4) + fib(3)
       = [fib(3) + fib(2)] + [fib(2) + fib(1)]
       
Without DP: Recalculates fib(3), fib(2) multiple times
With DP: Stores each result, calculates once
```

### When to use?
- Optimization problems (min/max)
- Counting problems (how many ways)
- Problems with overlapping subproblems
- Keywords: "maximum", "minimum", "longest", "count ways", "can you"
- Fibonacci-like patterns
- Decision at each step affects future

### Sample Problems
- **Easy**: Climbing Stairs (70), Min Cost Climbing Stairs (746), Best Time to Buy and Sell Stock (121)
- **Medium**: Coin Change (322), Longest Increasing Subsequence (300), House Robber (198), Unique Paths (62), Longest Common Subsequence (1143)

### Template Code (Go)

```go
// 1D DP - Climbing stairs
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2
    
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    
    return dp[n]
}

// 1D DP with space optimization
func climbStairsOptimized(n int) int {
    if n <= 2 {
        return n
    }
    
    prev2, prev1 := 1, 2
    
    for i := 3; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}

// 2D DP - Unique paths
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    
    // Initialize first row and column
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    
    // Fill the table
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    
    return dp[m-1][n-1]
}

// Top-down DP with memoization
func coinChange(coins []int, amount int) int {
    memo := make(map[int]int)
    
    var dp func(remaining int) int
    dp = func(remaining int) int {
        if remaining == 0 {
            return 0
        }
        if remaining < 0 {
            return -1
        }
        
        if val, exists := memo[remaining]; exists {
            return val
        }
        
        minCoins := int(1e9)
        for _, coin := range coins {
            result := dp(remaining - coin)
            if result >= 0 {
                minCoins = min(minCoins, result+1)
            }
        }
        
        if minCoins == int(1e9) {
            memo[remaining] = -1
        } else {
            memo[remaining] = minCoins
        }
        
        return memo[remaining]
    }
    
    return dp(amount)
}
```

---

## 9. Backtracking

### What is it?
Explores all possible solutions by building candidates incrementally and abandoning candidates ("backtracking") when they cannot lead to a valid solution. Think of it as intelligent brute force with pruning.

```
Generate all subsets of [1,2,3]:

                    []
           /         |         \
         [1]        [2]        [3]
        /   \        |
    [1,2]  [1,3]   [2,3]
      |
   [1,2,3]

At each node: decide to include or exclude element
```

### When to use?
- Generate all combinations/permutations
- Solve puzzles (Sudoku, N-Queens)
- Find all solutions to a constraint problem
- Subset/subarray generation
- Keywords: "all possible", "generate all", "find all combinations"
- When you need to try all possibilities with constraints

### Sample Problems
- **Easy**: Binary Tree Paths (257)
- **Medium**: Combinations (77), Permutations (46), Subsets (78), Letter Combinations of a Phone Number (17), Combination Sum (39), Generate Parentheses (22)

### Template Code (Go)

```go
// Subsets pattern
func subsets(nums []int) [][]int {
    result := [][]int{}
    current := []int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {
        // Add current subset
        temp := make([]int, len(current))
        copy(temp, current)
        result = append(result, temp)
        
        for i := start; i < len(nums); i++ {
            // Include nums[i]
            current = append(current, nums[i])
            backtrack(i + 1)
            // Exclude nums[i] (backtrack)
            current = current[:len(current)-1]
        }
    }
    
    backtrack(0)
    return result
}

// Combinations pattern
func combine(n int, k int) [][]int {
    result := [][]int{}
    current := []int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {
        // Base case: combination is complete
        if len(current) == k {
            temp := make([]int, k)
            copy(temp, current)
            result = append(result, temp)
            return
        }
        
        // Pruning: not enough elements left
        for i := start; i <= n-(k-len(current))+1; i++ {
            current = append(current, i)
            backtrack(i + 1)
            current = current[:len(current)-1]
        }
    }
    
    backtrack(1)
    return result
}

// Permutations pattern
func permute(nums []int) [][]int {
    result := [][]int{}
    used := make([]bool, len(nums))
    current := []int{}
    
    var backtrack func()
    backtrack = func() {
        if len(current) == len(nums) {
            temp := make([]int, len(nums))
            copy(temp, current)
            result = append(result, temp)
            return
        }
        
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            
            current = append(current, nums[i])
            used[i] = true
            backtrack()
            current = current[:len(current)-1]
            used[i] = false
        }
    }
    
    backtrack()
    return result
}

// Combination sum (can reuse elements)
func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    current := []int{}
    
    var backtrack func(start, remaining int)
    backtrack = func(start, remaining int) {
        if remaining == 0 {
            temp := make([]int, len(current))
            copy(temp, current)
            result = append(result, temp)
            return
        }
        
        if remaining < 0 {
            return
        }
        
        for i := start; i < len(candidates); i++ {
            current = append(current, candidates[i])
            backtrack(i, remaining-candidates[i]) // i, not i+1 (can reuse)
            current = current[:len(current)-1]
        }
    }
    
    backtrack(0, target)
    return result
}
```

---

## 10. Prefix Sum

### What is it?
Preprocesses an array to store cumulative sums, allowing constant-time range sum queries. The prefix sum at index i represents the sum of all elements from index 0 to i.

```
Original array: [1, 2, 3, 4, 5]
Prefix sum:     [1, 3, 6, 10, 15]

Sum of range [1,3] = prefix[3] - prefix[0] = 10 - 1 = 9
Which is: 2 + 3 + 4 = 9
```

### When to use?
- Multiple range sum queries
- Subarray sum equals K
- Contiguous subarray problems
- 2D matrix sum queries
- Problems asking for "sum of range" or "subarray with sum X"

### Sample Problems
- **Easy**: Running Sum of 1d Array (1480), Find Pivot Index (724)
- **Medium**: Subarray Sum Equals K (560), Contiguous Array (525), Product of Array Except Self (238), Range Sum Query 2D (304)

### Template Code (Go)

```go
// Basic prefix sum
func prefixSum(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    prefix[0] = nums[0]
    
    for i := 1; i < n; i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    
    return prefix
}

// Range sum query
func rangeSum(prefix []int, left, right int) int {
    if left == 0 {
        return prefix[right]
    }
    return prefix[right] - prefix[left-1]
}

// Subarray sum equals K
func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0
    prefixMap := make(map[int]int)
    prefixMap[0] = 1 // Important: empty subarray
    
    for _, num := range nums {
        prefixSum += num
        
        // Check if there exists a prefix sum such that
        // prefixSum - previousSum = k
        if freq, exists := prefixMap[prefixSum-k]; exists {
            count += freq
        }
        
        prefixMap[prefixSum]++
    }
    
    return count
}

// Product except self using prefix concept
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    
    // Left products
    result[0] = 1
    for i := 1; i < n; i++ {
        result[i] = result[i-1] * nums[i-1]
    }
    
    // Right products
    rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= rightProduct
        rightProduct *= nums[i]
    }
    
    return result
}

// 2D prefix sum
type NumMatrix struct {
    prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{}
    }
    
    m, n := len(matrix), len(matrix[0])
    prefix := make([][]int, m+1)
    for i := range prefix {
        prefix[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            prefix[i][j] = matrix[i-1][j-1] +
                          prefix[i-1][j] +
                          prefix[i][j-1] -
                          prefix[i-1][j-1]
        }
    }
    
    return NumMatrix{prefix: prefix}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
    row1++; col1++; row2++; col2++
    return nm.prefix[row2][col2] -
           nm.prefix[row1-1][col2] -
           nm.prefix[row2][col1-1] +
           nm.prefix[row1-1][col1-1]
}
```

---

## Pattern Selection Guide

When facing a problem, ask yourself:

1. **Is the array sorted?** → Consider Binary Search or Two Pointers
2. **Need to find a contiguous subarray/substring?** → Consider Sliding Window
3. **Working with linked lists?** → Consider Fast & Slow Pointers
4. **Tree or graph traversal?** → Consider BFS or DFS
5. **Need top K elements?** → Consider Heap
6. **Multiple range queries?** → Consider Prefix Sum
7. **All possible combinations?** → Consider Backtracking
8. **Overlapping subproblems?** → Consider Dynamic Programming
9. **Two arrays to compare?** → Consider Two Pointers or DP
10. **Optimization problem (min/max)?** → Consider DP or Greedy

## Time Complexity Quick Reference

- Two Pointers: O(n)
- Sliding Window: O(n)
- Fast & Slow Pointers: O(n)
- Tree DFS: O(n)
- Tree BFS: O(n)
- Binary Search: O(log n)
- Heap Operations: O(log k) per operation
- Dynamic Programming: O(n²) to O(n³) typically
- Backtracking: O(2ⁿ) or O(n!) typically
- Prefix Sum: O(n) preprocessing, O(1) query

## Helper Functions

```go
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
```
