// Package sortcolors: https://leetcode.com/problems/sort-colors/description/
package sortcolors

import "slices"

// For this solution, we re-use the solution in move_zeros: first we move all 2 to the back of the array; then find the new array without 2, and move all 1 to the back of the array.
func sortColors(nums []int) {
	move(nums, 2)
	indexOfLast2 := slices.Index(nums, 2)
	if indexOfLast2 == -1 {
		indexOfLast2 = len(nums)
	}

	move(nums[:indexOfLast2], 1)
}

func move(nums []int, value int) {
	nextNonZero := 0

	for i := range nums {
		if nums[i] != value {
			nums[nextNonZero], nums[i] = nums[i], nums[nextNonZero]
			nextNonZero++
		}
	}
}

func sortColors2(nums []int) {
	left, right := 0, len(nums)-1
	i := 0

	for i <= right {
		switch nums[i] {
		case 0:
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		case 2:
			nums[right], nums[i] = nums[i], nums[right]
			right--
		default:
			i++
		}
	}
}
