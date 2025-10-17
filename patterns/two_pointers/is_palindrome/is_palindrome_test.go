package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"case 1", "A man, a plan, a canal: Panama", true},
		{"case 2", "race a car", false},
		{"case 3", " ", true},
		{"case 4", "0P", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindrome(test.input)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}

func TestIsPalindromeOptimized(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"case 1", "A man, a plan, a canal: Panama", true},
		{"case 2", "race a car", false},
		{"case 3", " ", true},
		{"case 4", "0P", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindromeOptimized(test.input)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}
