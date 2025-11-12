package main

import (
	"testing"

	"github.com/daniel/dsa-space/common"
)

func TestTopKFrequentHashAndSort(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"Case 1, multiple elements", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"Case 2, one element", []int{1}, 1, []int{1}},
		{"Case 3, another multiple elements", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{"Case 4, complex example", []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10, []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11}},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := topKFrequentHashAndSort(test.nums, test.k)
			if !common.SliceEqual(result, test.want) {
				t.Errorf("Got: %v, want: %v", result, test.want)
			}
		})
	}
}

func TestTopKFrequentBucketSort(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"Case 1, multiple elements", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"Case 2, one element", []int{1}, 1, []int{1}},
		{"Case 3, another multiple elements", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{"Case 4, complex example", []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10, []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11}},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := topKFrequentBucketSort(test.nums, test.k)
			if !common.SliceEqual(result, test.want) {
				t.Errorf("Got: %v, want: %v", result, test.want)
			}
		})
	}
}
