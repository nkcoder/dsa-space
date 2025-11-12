package main

import (
	"slices"
)

// Hash + sort: count the nums, sort by the count, then return the top k elements.
func topKFrequentHashAndSort(nums []int, k int) []int {
	numToCount := make(map[int]int)
	for _, v := range nums {
		numToCount[v]++
	}

	countToNums := make(map[int][]int)
	for num, count := range numToCount {
		countToNums[count] = append(countToNums[count], num)
	}

	counts := []int{}
	for i := range countToNums {
		counts = append(counts, i)
	}

	slices.SortFunc(counts, func(x, y int) int {
		return y - x
	})

	result := []int{}
	i := 0
	for i < len(counts) && len(result) < k {
		result = append(result, countToNums[counts[i]]...)
		i++
	}

	return result
}

// We need to sort by the count, use bucket sort (array index) if the count is not too large.
func topKFrequentBucketSort(nums []int, k int) []int {
	numToCount := make(map[int]int)
	for _, v := range nums {
		numToCount[v]++
	}

	frequency := make([][]int, len(nums)+1)
	for n, c := range numToCount {
		frequency[c] = append(frequency[c], n)
	}

	result := []int{}
	for i := len(frequency) - 1; i > 0; i-- {
		result = append(result, frequency[i]...)
		if len(result) >= k {
			return result
		}
	}

	return result
}
