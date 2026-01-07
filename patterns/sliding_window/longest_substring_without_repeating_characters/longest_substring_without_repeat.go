// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package main

// Needs to careful about the edge cases which impacts the update of the window boundary.
// Check the test cases for edge cases: 1. no repeat characters at all; 2. character repeats outside of the window.
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

// Use the variable-length sliding window problem template: the invalid state is a character appears more than once
func lengthOfLongestSubstringByTemplate(s string) int {
	state := make(map[byte]int)

	start, longest := 0, 0
	for end := 0; end < len(s); end++ {
		state[s[end]]++

		for state[s[end]] > 1 {
			state[s[start]]--
			start++
		}
		longest = max(longest, end-start+1)
	}
	return longest
}

// Inspried by the sliding window problem template, can be a bit faster since we can track the index of each character in the state.
// If we find the previous index of an existing character, we can advance `start` to the previous index to avoid a loop.
func lengthOfLongestSubstringFaster(s string) int {
	state := make(map[byte]int)

	start, longest := 0, 0
	for end := 0; end < len(s); end++ {
		// Current character exists in the state, advance `start` to previous index + 1
		if lastIndex, exists := state[s[end]]; exists {
			start = max(start, lastIndex+1) // use max to ensure start always move forward (the character might occure before the current window)
		}
		state[s[end]] = end
		longest = max(longest, end-start+1)
	}
	return longest
}

// Similar to the template solution, but update the state first if the character exists, and use byte->bool for the state map
func lengthOfLongestSubstring3(s string) int {
	seen := make(map[byte]bool)

	start, longest := 0, 0
	for end := 0; end < len(s); end++ {
		for seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		seen[s[end]] = true
		longest = max(longest, end-start+1)
	}

	return longest
}
