// Package validtrianglenumber: https://leetcode.com/problems/valid-triangle-number/
//
// Sort the input, then we have all the elements a <= b <= c, we only need to check that a + b > c for a triangle.
// We can fix on the right side of the input, then we can use two pointers to check all the pairs.
// **Warning**: we cannot fix on the left side (smallest side) as it might miss pairs, for example: [2, 2, 3, 4], the triangle [2, 2, 3] is missing
package validtrianglenumber

import (
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	count := 0

	for i := len(nums) - 1; i > 1; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
