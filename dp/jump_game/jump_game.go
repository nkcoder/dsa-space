package main

import "fmt"

/*
 * Medium: https://leetcode.com/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
 *
 * You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
 * Return true if you can reach the last index, or false otherwise.
 */

func canJump(nums []int) bool {
	// the farthest index we can possibly reach
	maxReach := 0
	for i := range nums {
		// if I can't even get to position i, how can I possibly reach the end?
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0}))             // true
	fmt.Println(canJump([]int{1, 2, 3}))       // true
	fmt.Println(canJump([]int{0, 2, 3}))       // false
}
