// Package meetingroom: https://leetcode.com/problems/meeting-rooms-ii/description/ (needs subscription)
// https://neetcode.io/problems/meeting-schedule-ii?list=neetcode150
//
// Given an array of meeting time interval objects consisting of start and end times
// [[start_1,end_1],[start_2,end_2],...] (start_i < end_i), find the minimum number of days required to schedule all meetings without any conflicts.
// Note: (0,8),(8,10) is not considered a conflict at 8.
// Example 1: Input: intervals = [(0,40),(5,10),(15,20)], Output: 2
// Example 2: Input: intervals = [(4,9)] Output: 1
package meetingroom

import (
	"slices"
)

type Interval struct {
	start int
	end   int
}

// Sort + two pointers
// When a new meeting starts, we need a new room; when a meeting ends, a room frees up.
// Time complexity: O(nlogn) + O(n)
// Space complexity: O(n)
func minMeetingRoomsTwoPointers(intervals []Interval) int {
	n := len(intervals)
	starts, ends := []int{}, []int{}

	for _, inter := range intervals {
		starts = append(starts, inter.start)
		ends = append(ends, inter.end)
	}

	// Sort the starts and ends
	slices.Sort(starts)
	slices.Sort(ends)

	count, max := 0, 0
	i, j := 0, 0

	for i < n && j < n {
		// A new meetings starts before the earliest meeting ends, need a new room
		if starts[i] < ends[j] {
			i++
			count++
		} else { // A meeting has finished (or finishes exactly when the new one starts), freeing a room
			j++
			count--
		}
		if count > max {
			max = count
		}
	}

	return max
}

// To understand how many conference rooms we need, think about what happens at each moment in time.
// At any given moment, the number of rooms needed equals the number of meetings happening simultaneously.
// For example: [[0, 30], [5, 10], [15, 20]], if we sort all the slots (but also tracks if it starts or ends a meeting):
// 1. If a meeting starts, then need a new room
// 2. If a meeting finishes, free up a room.
// It's more clear to show on the digram to see the maximum concurrent meetings
//
// We can use the bucket sort to sort all start and end times (1 for start, -1 for end)
// Associate the start time slot with 1, and end time slot with -1, then calculate the maximum prefix sum.
// But if the time values are large, then it will not be memory efficient, for example: [[1000000, 1000005], [1000002, 1000007]]
// Time complexity: O(n)
// Space complexity: O(n)
func minMeetingRoomsBuckets(intervals []Interval) int {
	maxEnd := 0
	for _, inter := range intervals {
		maxEnd = max(maxEnd, inter.end)
	}

	slots := make([]int, maxEnd+1)
	// We need to use ++ and --, not assign to 1 or -1 here: [[0,20], [20,40]], one room finishes and another room starts at the same time.
	for _, inter := range intervals {
		slots[inter.start]++
		slots[inter.end]--
	}

	count := 0
	sum := 0
	for _, v := range slots {
		sum += v
		count = max(count, sum)
	}

	return count
}
