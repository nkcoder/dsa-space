package main

import "testing"

func TestMinMeetingRoomsTwoPointers(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      int
	}{
		{"case 1", []Interval{{0, 40}, {5, 10}, {15, 20}}, 2},
		{"case 2", []Interval{{4, 9}}, 1},
		{"case 3", []Interval{{1000001, 1000010}, {1000008, 1000030}, {1000040, 1000050}}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minMeetingRoomsTwoPointers(test.intervals)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func TestMinMeetingRoomsBits(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      int
	}{
		{"case 1", []Interval{{0, 40}, {5, 10}, {15, 20}}, 2},
		{"case 2", []Interval{{4, 9}}, 1},
		{"case 3", []Interval{{1000001, 1000010}, {1000008, 1000030}, {1000040, 1000050}}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minMeetingRoomsBits(test.intervals)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
