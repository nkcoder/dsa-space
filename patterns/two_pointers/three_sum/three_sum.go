// https://leetcode.com/problems/3sum/description/
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
//
// Time complexity: O(n^2), we perform n iterations of the outer loop, and each iteration takes O(n) time to use the two-pointer technique.
// Space complexity: O(n^2), the output array can have maximum n^2 items (for each choice of the first number, there are O(n) valid pairs); except for the output, the space complexity is O(1).

package main

import (
	"slices"
)

// Sort + two pointers: the most efficient
func threeSum(nums []int) [][]int {
	slices.Sort(nums) // the same as: sort.Ints(nums)

	result := make([][]int, 0) // the same as result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// if the first element greater than 0, early return
		if nums[i] > 0 {
			break
		}

		// skip duplicates for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		target := 0 - nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				// skip duplicate for j
				for left < right && nums[left-1] == nums[left] {
					left++
				}

				// skip duplicate for k
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
