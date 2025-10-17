package main

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{"return all triplets", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"handle empty triplets", []int{0, 1, 1}, [][]int{}},
		{"handle all 0 inputs", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := threeSum(test.input)
			if !sliceEqual(result, test.want) {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func sliceEqual(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if len(x[i]) != len(y[i]) {
			return false
		}
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}

	return true
}
