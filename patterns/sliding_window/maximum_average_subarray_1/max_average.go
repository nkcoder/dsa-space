package main

// https://leetcode.com/problems/maximum-average-subarray-i/description/

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for m := range k {
		max += nums[m]
	}
	preSum := max

	i, j := 1, k
	for j < len(nums) {
		currentSum := preSum + nums[j] - nums[i-1]

		if currentSum > max {
			max = currentSum
		}
		preSum = currentSum

		i++
		j++
	}

	return float64(max) / float64(k)
}
