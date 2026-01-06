package sortcolors

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Test Case 1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"Test Case 2", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sortColors(test.input)
			if !slices.Equal(test.input, test.want) {
				t.Errorf("Got: %v, want: %v\n", test.input, test.want)
			}
		})
	}
}

func TestSortColors2(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Test Case 1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"Test Case 2", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sortColors2(test.input)
			if !slices.Equal(test.input, test.want) {
				t.Errorf("Got: %v, want: %v\n", test.input, test.want)
			}
		})
	}
}
