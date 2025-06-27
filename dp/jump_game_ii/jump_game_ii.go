package main

import "fmt"

/*
 * Medium: https://leetcode.com/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150
 *
 * You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
 * Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
 * 	0 <= j <= nums[i] and
 * 	i + j < n
 * Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
 */

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 0      // the minimum jumps needed
	currentEnd := 0 // the current end of the jump range
	farthest := 0   // the farthest index that can be reached so far

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i]) // update max reachable index
		// we've finished exploring all positions reachable with the current number of jumps
		if i == currentEnd { // reached end of current jump level
			jumps++               // must make another jump
			currentEnd = farthest // new boundry is the farthest we can go
		}
	}

	return jumps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))    // Output: 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))    // Output: 2
	fmt.Println(jump([]int{2, 3, 1, 1, 4, 5})) // Output: 3
	fmt.Println(jump([]int{3}))                // Output: 0
}
