package findtheduplicatenumber

import "testing"

func TestFindDuplicate(t *testing.T) {

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"Case 1 ", []int{1, 3, 4, 2, 2}, 2},
		{"Case 2", []int{3, 1, 3, 4, 2}, 3},
		{"Case 3", []int{3, 3, 3, 3, 3}, 3},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := findDuplicate(test.input)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}

			result2 := findDuplicate2(test.input)
			if result2 != test.want {
				t.Errorf("Got: %d, want: %d", result2, test.want)
			}

		})
	}
}
