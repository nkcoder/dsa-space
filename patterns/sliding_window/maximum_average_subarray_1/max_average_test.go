package main

import "testing"

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"case 1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"case 2", []int{5}, 1, 5},
		{"case 3", []int{-1}, 1, -1},
		{"case 4", []int{0, 4, 0, 3, 2}, 1, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findMaxAverage(test.nums, test.k)
			if result != test.want {
				t.Errorf("Got: %.2f, want: %.2f", result, test.want)
			}
		})
	}
}
