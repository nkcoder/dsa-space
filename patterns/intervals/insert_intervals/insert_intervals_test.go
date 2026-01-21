package insertintervals

import (
	"testing"

	"github.com/daniel/dsa-space/common"
)

func TestInsertIntervals(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "Test case 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{name: "Test case 2", intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := insertIntervals(test.intervals, test.newInterval)
			if !common.SliceArrayEqual(result, test.want) {
				t.Errorf("Got: %v, want: %v", result, test.want)
			}
		})
	}
}
