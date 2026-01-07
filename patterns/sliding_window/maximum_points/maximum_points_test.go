package maximumpoints

import "testing"

func TestMaxScore(t *testing.T) {
	testCases := []struct {
		name       string
		cardPoints []int
		k          int
		want       int
	}{
		{"Case 1", []int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{"Case 2", []int{2, 2, 2}, 2, 4},
		{"Case 3", []int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := maxScore(test.cardPoints, test.k)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func TestMaxScore2(t *testing.T) {
	testCases := []struct {
		name       string
		cardPoints []int
		k          int
		want       int
	}{
		{"Case 1", []int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{"Case 2", []int{2, 2, 2}, 2, 4},
		{"Case 3", []int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := maxScore2(test.cardPoints, test.k)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
