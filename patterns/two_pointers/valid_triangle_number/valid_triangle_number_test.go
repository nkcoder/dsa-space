package validtrianglenumber

import "testing"

func TestValidTriangleNumber(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test case 1", []int{2, 2, 3, 4}, 3},
		{"Test case 2", []int{4, 2, 3, 4}, 4},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := triangleNumber(test.input)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
