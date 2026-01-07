package fruitesintobuckets

import (
	"testing"
)

func TestTotalFruits(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test case 1", []int{1, 2, 1}, 3},
		{"Test case 2", []int{0, 1, 2, 2}, 3},
		{"Test case 3", []int{1, 2, 3, 2, 2}, 4},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := totalFruits(test.input)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
