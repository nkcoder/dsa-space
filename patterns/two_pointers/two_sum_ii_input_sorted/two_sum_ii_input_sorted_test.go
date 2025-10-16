package main

import (
	"slices"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		target   int
		want     []int
	}{
		{"consecutive index", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"non consecutive index", []int{2, 3, 4}, 6, []int{1, 3}},
		{"negative target", []int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum2(tt.inputArr, tt.target)
			if !slices.Equal(result, tt.want) {
				t.Errorf("got: %d, want: %d", result, tt.want)
			}
		})
	}
}
