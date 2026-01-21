package nonoverlappingintervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{name: "test case 1", intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, want: 1},
		{name: "test case 2", intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}, want: 2},
		{name: "test case 3", intervals: [][]int{{1, 2}, {2, 3}}, want: 0},
		{name: "test case 4", intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, want: 2},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := eraseOverlapIntervals(test.intervals)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
