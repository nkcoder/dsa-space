package main

// https://leetcode.com/problems/search-insert-position/
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	searchInsert([]int{1, 3, 5, 6}, 5) // returns 2
	searchInsert([]int{1, 3, 5, 6}, 2) // returns 1
	searchInsert([]int{1, 3, 5, 6}, 7) // returns 4
	searchInsert([]int{1, 3, 5, 6}, 0) // returns 0
}
