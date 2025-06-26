package main

import "fmt"

/*
 * https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
 * Medium: Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
 */

// for each number, calculate its new position and place it there
func rotate(nums []int, k int) {
	rotated_nums := make([]int, len(nums))
	for i, v := range nums {
		j := (i + k) % len(nums)
		rotated_nums[j] = v
	}
	copy(nums, rotated_nums)
}

// reverse the array in place
func rotate2(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // [5 6 7 1 2 3 4]

	nums2 := []int{-1, -10, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2) // [3 99 -1 -10]

	nums3 := []int{1, 2}
	rotate(nums3, 3)
	fmt.Println(nums3) // [2, 1]

	nums4 := []int{-1, -1}
	rotate(nums3, 3)
	fmt.Println(nums4) // [-1, -1]
}
