// https://leetcode.com/problems/contains-duplicate-ii/description/
//
// Given an integer array nums and an integer k, return true if
// there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
//
// Example 1: Input: nums = [1,2,3,1], k = 3, Output: true
// Example 2: Input: nums = [1,0,1,1], k = 1, Output: true
// Example 3: Input: nums = [1,2,3,1,2,3], k = 2, Output: false
package main

// Use sliding window: for each left index, keep a max window of (left + k) elements.
// almost timeout: 979ms, Beats 5.05%
func containsNearbyDuplicatesSlidingWindow(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j-i <= k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// Use hash: index map to check if the same element has occurred, if yes, check the index.
// Still not fast enough: 60ms, Beats 18.95%
func containsNearbyDuplicatesHash(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, v := range nums {
		index, ok := indexMap[v]
		if ok && abs(index-i) <= k {
			return true
		}

		indexMap[v] = i
	}

	return false
}

// math.Abs only works for float64
func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

// Use sliding window with hash set: maintain a set of at most k elements
// 21ms, Beats 96.30%, maintain only revelant members in the window
func containsNearByDuplicatesSlidingHash(nums []int, k int) bool {
	// Pre allocate map with expected size to avoid rehashing
	window := make(map[int]bool, k+1)
	for i, v := range nums {
		// Check if current member already exists in the window
		if window[v] {
			return true
		}

		// Add current member to the window
		window[v] = true

		// Remove elements that is outside of the window (more than k distance)
		if i >= k {
			delete(window, nums[i-k])
		}
	}

	return false
}
