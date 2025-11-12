package main

import (
	"testing"

	"github.com/daniel/dsa-space/common"
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
			if !common.SliceArrayEqual(result, test.want) {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
