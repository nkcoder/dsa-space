// Package intervals: Interval problems typically involve sorting the given intervals, and then processing each interval in sorted each order.
package intervals

import "slices"

type Interval struct {
	start int
	end   int
}

// Sorting by start time.
// Sorting intervals by their start times makes it easy to merge two intervals that are overlapping.
// 1. Detect if any overlapping in the intervals
func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}
	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.start - b.start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].start {
			return false
		}
	}

	return true
}

// 2. Merge intervals
// To merge an interval into a previous interval, we set the end time of the previous interval to be the max of either end time.
func mergeIntervals(intervals []Interval) []Interval {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.start - b.start
	})

	merged := make([]Interval, 0)

	for _, interval := range intervals {
		if len(merged) == 0 || interval.start > merged[len(merged)-1].end {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1].end = max(merged[len(merged)-1].end, interval.end)
		}
	}

	return merged
}

// Sort by end time.
// Find the maximum number of non-overlapping intervals in a given list of intervals
// If we sort by start time, we risk adding an interval that starts early but ends late, which will block us from adding other intervals until that interval ends.
// This is the same question as: find the minimum number of intervals to remove to eliminate any overlap.
func nonOverlappingIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.end - b.end
	})

	count := 1
	end := intervals[0].end
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start > end {
			end = intervals[i].end
			count++
		}
	}

	return count
}
