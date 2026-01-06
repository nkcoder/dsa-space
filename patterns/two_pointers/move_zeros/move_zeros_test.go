package movezeros

import (
	"fmt"
	"slices"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Test case 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Test case 1", []int{0}, []int{0}},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			moveZeroes(test.input)
			if !slices.Equal(test.input, test.want) {
				fmt.Printf("Got: %v, want: %v", test.input, test.want)
			}
		})
	}
}
