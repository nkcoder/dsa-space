package main

import "testing"

func TestContainsNearbyDuplicatesSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"case 1", []int{1, 2, 3, 1}, 3, true},
		{"case 2", []int{1, 0, 1, 1}, 1, true},
		{"case 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := containsNearbyDuplicatesSlidingWindow(test.nums, test.k)
			if result != test.want {
				t.Errorf("Got %t, want %t", result, test.want)
			}
		})
	}
}

func TestContainsNearbyDuplicatesHash(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"case 1", []int{1, 2, 3, 1}, 3, true},
		{"case 2", []int{1, 0, 1, 1}, 1, true},
		{"case 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := containsNearbyDuplicatesHash(test.nums, test.k)
			if result != test.want {
				t.Errorf("Got %t, want %t", result, test.want)
			}
		})
	}
}

func TestContainsNearbyDuplicatesSlidingHash(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"case 1", []int{1, 2, 3, 1}, 3, true},
		{"case 2", []int{1, 0, 1, 1}, 1, true},
		{"case 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := containsNearByDuplicatesSlidingHash(test.nums, test.k)
			if result != test.want {
				t.Errorf("Got %t, want %t", result, test.want)
			}
		})
	}
}
