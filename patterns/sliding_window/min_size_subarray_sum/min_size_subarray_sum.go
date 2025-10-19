// https://leetcode.com/problems/minimum-size-subarray-sum/
// Sliding window: the key is when to update the `min` result
package main

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j, min := 0, 0, 0
	curSum := nums[0]
	for i < len(nums) && j < len(nums) {
		if curSum >= target {
			if j-i+1 < min || min == 0 {
				min = j - i + 1
			}
			curSum -= nums[i]
			i++
		} else {
			if j == len(nums)-1 {
				break
			}
			j++
			curSum += nums[j]
		}
	}

	return min
}
