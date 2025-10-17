package main

import (
	"fmt"
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
			if !slicesEqual(result, test.want) {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func slicesEqual(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}

	sliceToStr := func(s []int) string {
		return fmt.Sprint(s)
	}

	count := make(map[string]int)
	for _, xm := range x {
		count[sliceToStr(xm)]++
	}

	for _, ym := range y {
		key := sliceToStr(ym)
		count[key]--
		if count[key] < 0 {
			return false
		}
	}

	return true
}
