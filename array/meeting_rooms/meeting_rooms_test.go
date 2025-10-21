package main

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      bool
	}{
		{"case 1", []Interval{{0, 40}, {5, 10}, {15, 20}}, false},
		{"case 2", []Interval{{4, 9}}, true},
		{"case 3", []Interval{{1000001, 1000010}, {1000008, 1000030}, {1000040, 1000050}}, false},
		{"case 4", []Interval{{0, 30}, {5, 10}, {15, 20}}, false},
		{"case 5", []Interval{{5, 8}, {9, 15}}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := canAttendMeetings(test.intervals)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}
