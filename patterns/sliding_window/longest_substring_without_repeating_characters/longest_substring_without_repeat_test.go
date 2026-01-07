package main

import "testing"

var tests = []struct {
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

func TestLengthOfLongetSubstring(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(test.s)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func TestLengthOfLongetSubstring3(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstring3(test.s)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func TestLengthOfLongetSubstringByTemplate(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstringByTemplate(test.s)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}

func TestLengthOfLongetSubstringFaster(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstringFaster(test.s)
			if result != test.want {
				t.Errorf("Got: %d, want: %d", result, test.want)
			}
		})
	}
}
