package longestsubstringreplacement

import "testing"

func TestCharacterReplace(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"Test case 1", "ABAB", 2, 4},
		{"Test case 2", "AABABBA", 1, 4},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := characterReplacement(test.s, test.k)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
