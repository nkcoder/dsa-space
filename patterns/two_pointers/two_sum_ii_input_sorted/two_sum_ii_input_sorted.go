// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// The input array is sorted, and we need to find pair for sum, using two pointers on the input.

package main

func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return []int{}
}
