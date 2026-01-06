// Package sortcolors: https://leetcode.com/problems/sort-colors/description/
package sortcolors

import "slices"

// Solution 1, we re-use the solution in move_zeros: first we move all 2 to the back of the array; then find the new array without 2, and move all 1 to the back of the array.
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

// Solution 2: use 3 pointers, left, i and right, then:
// 1. All elements before left are 0
// 2. All elements between left and i-1 are 1
// 3. All elements between i and right are unsorted (not processed)
// 4. All elements after right are 2
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
