package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"case 1", []int{2, 3, 1, 2, 4, 3}, 7, 2},
		{"case 2", []int{1, 4, 4}, 4, 1},
		{"case 3", []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minSubArrayLen(test.target, test.nums)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
