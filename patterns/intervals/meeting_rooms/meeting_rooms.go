// Package meetingrooms: https://leetcode.com/problems/meeting-rooms/description/ (needs subscription)
// https://www.hellointerview.com/learn/code/intervals/can-attend-meetings
//
// Given an array of meeting time interval objects consisting of start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i), determine if a person could add all meetings to their schedule without any conflicts.
// Example 1: Input: intervals = [(0,30),(5,10),(15,20)], Output: false
// Example 2: Input: intervals = [(5,8),(9,15)], Output: true
package meetingrooms

import (
	"slices"
)

type Interval struct {
	start int
	end   int
}

// Sort by interval start, if the next interval start > previous interval end, then overlap
// Time complexity: O(nlogn)
// Space complexity: O(1)
func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.start - b.start
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].end > intervals[i+1].start {
			return false
		}
	}

	return true
}
