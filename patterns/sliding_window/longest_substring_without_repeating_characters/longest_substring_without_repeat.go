// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// Sliding window: needs to careful about the edge cases which impacts the update of the window boundary.
// Check the test cases for edge cases: 1. no repeat characters at all; 2. character repeats outside of the window.
package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := make(map[byte]int)

	left, right, max := 0, 0, 1
	for right < len(s) {
		index, ok := seen[s[right]]
		// the character is repeat IN THE WINDOW
		if ok && index >= left {
			if right-left > max {
				max = right - left
			}
			left = index + 1
		}
		seen[s[right]] = right
		right++

		// handles no repeat characters
		if right == len(s) && right-left > max {
			max = right - left
		}
	}

	return max
}
