package main

import "testing"

func TestLengthOfLongetSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"case 1", "abcabcbb", 3},
		{"case 2", "bbb", 1},
		{"case 3", "pwwkew", 3},
		{"case 4", " ", 1},
		{"case 5", "au", 2},
		{"case 6", "abba", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(test.s)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
