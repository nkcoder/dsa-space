package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{"case 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"caes 2", []int{2, 3, 4}, 6, []int{0, 2}},
		{"case 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := twoSum(test.input, test.target)
			if slices.Equal(result, test.want) {
				t.Errorf("Got %v, want: %v", result, test.want)
			}
		})
	}
}
