package main

import "testing"

func TestIsHappyHash(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"case 1: happy number", 19, true},
		{"case 2: not happy number", 2, false},
		{"case 3: not happy number", 7, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isHappyHash(test.n)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}

func TestIsHappyFastSlowPointers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"case 1: happy number", 19, true},
		{"case 2: not happy number", 2, false},
		{"case 3: not happy number", 7, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isHappyFastSlowPointers(test.n)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}
