package main

import "testing"

func TestMaxWater(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"case 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"case 2", []int{1, 1}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxWater(test.height)
			if result != test.want {
				t.Errorf("Got %d, want %d", result, test.want)
			}
		})
	}
}
