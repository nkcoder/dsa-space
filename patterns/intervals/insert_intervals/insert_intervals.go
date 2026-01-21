// Package insertintervals  https://leetcode.com/problems/insert-interval/description/
package insertintervals

// 3 steps
func insertIntervals(intervals [][]int, newInterval []int) [][]int {
	merged := make([][]int, 0)
	i, n := 0, len(intervals)

	// 1. find all intervals that ends before the new interval starts, push to merged directly
	for i < n && intervals[i][1] < newInterval[0] {
		merged = append(merged, intervals[i])
		i++
	}

	// 2. find all overlaps and merge into a single interval
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	merged = append(merged, newInterval)

	// 3. find all intervals that starts after the new interval ends, push to merged directly
	for i < n {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}
