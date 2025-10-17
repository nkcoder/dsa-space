package main

import (
	"slices"
)

// https://leetcode.com/problems/3sum/description/
//
// Sort + two pointers: the most efficient
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
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
