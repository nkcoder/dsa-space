// Package nonoverlappingintervals: https://leetcode.com/problems/non-overlapping-intervals/submissions/1891971580/
// https://www.hellointerview.com/learn/code/intervals/non-overlapping-intervals
package nonoverlappingintervals

import "slices"

// This is the same problem as: get the longest non-overlapping intervals. The minimum number of intervals need to move: length of intervals - longest non-overlapping intervals.
// Sort by end time, then count the longest non-overlapping intervals
func eraseOverlapIntervals(intervals [][]int) int {
	nonOverlappingCount := 1

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[1] - b[1]
	})

	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			nonOverlappingCount++
			end = intervals[i][1]
		}
	}
	return len(intervals) - nonOverlappingCount
}
