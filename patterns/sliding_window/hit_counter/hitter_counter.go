// Design a Hit Counter that counts the number of hits received in the past 5 minutes (i.e., past 300 seconds). The system should provide two functions:

// hit(timestamp): Record a hit that occurs at the given timestamp (in seconds).
// getHits(timestamp): Return the number of hits in the past 5 minutes from the given timestamp (i.e., count all hits that occurred from timestamp-299 to timestamp, inclusive).
// Assume that:
//
// Timestamps are given in non-decreasing chronological order (each new call’s timestamp is >= previous call’s timestamp).
// The earliest timestamp starts at 1.
// It’s possible to have multiple hits at the same timestamp (multiple calls to hit(t) in the same second).
//
//Example 2:
// Operations: hit(100) (called 3 times), getHits(100), getHits(101), getHits(400)
//
// Three hits at timestamp 100.
// getHits(100) counts hits in [100-299, 100] = [−199, 100], which includes all three hits at 100. Output: 3.
// getHits(101) counts hits in [101-299, 101] = [−198, 101], still includes the three hits at 100 (they occurred within the last 5 minutes). Output: 3.
// getHits(400) counts hits in [400-299, 400] = [101, 400]. The hits at 100 are now more than 5 minutes old (100 is outside [101,400]), so they are not counted. Output: 0.

// Use a queue (list in container) to maintain only the hits from the last 300 seconds. The queue naturally preserves the order of hits.
package main

import (
	"container/list"
	"fmt"
)

type HitCounter struct {
	hits *list.List
}

func NewHitCounter() *HitCounter {
	return &HitCounter{hits: list.New()}
}

func (hc *HitCounter) hit(timestamp int) {
	hc.hits.PushBack(timestamp)
}

func (hc *HitCounter) getHits(timestamp int) int {
	startTimestamp := max(timestamp-300+1, 0)
	fmt.Println("startTimestamp", startTimestamp, "hits", hc.hits.Len())
	for hc.hits.Len() != 0 {
		front := hc.hits.Front()
		if front.Value.(int) < startTimestamp {
			hc.hits.Remove(front)
		} else {
			break
		}
	}

	return hc.hits.Len()
}
