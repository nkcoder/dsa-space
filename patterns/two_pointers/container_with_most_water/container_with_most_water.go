package main

// https://leetcode.com/problems/container-with-most-water/description/

// Two pointers
// The area is: the minimum of the heights * the width (distance between hieghts). The key point here is how to move the pointers.
// No matter we move which pointer, the width decreases. The only possible to get a bigger container area is to find a bigger height.
// So the condition to move the pointer is: always find the bigger height between the left and right pointer.
func maxWater(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		width := right - left
		currentMax := min(height[left], height[right]) * width
		if currentMax > max {
			max = currentMax
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
