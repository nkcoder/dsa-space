package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"caes 1", "abc", "ahbgdc", true},
		{"case 2", "axc", "ahbgdc", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSubsequence(test.s, test.t)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}
