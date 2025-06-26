package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := numSet[num]; exists {
			return true
		}
		numSet[num] = struct{}{}
	}
	return false
}

// Example usage
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums)) // true

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) // false
}
