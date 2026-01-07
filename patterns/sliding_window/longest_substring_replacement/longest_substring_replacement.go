// Package longestsubstringreplacement: https://leetcode.com/problems/longest-repeating-character-replacement/description/
// https://www.hellointerview.com/learn/code/sliding-window/longest-repeating-character-replacement
package longestsubstringreplacement

// We can still use the variable length sliding window problem template, just need to find the invalid status.
// Actually we need to find the longest repetive substring (you can have k characters that is different), for example:
// (AABABBA, 1): we're lookng for longest repetive A (or B) that can have 1 B (or A) inside.
// The invalid status is: (end-start+1) - maxFreq > k
func characterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	start, longest, maxFreq := 0, 0, 0

	for end := 0; end < len(s); end++ {
		charCount[s[end]]++
		maxFreq = max(maxFreq, charCount[s[end]])

		if end-start+1-maxFreq > k {
			charCount[s[start]]--
			start++
		}

		longest = max(longest, end-start+1)
	}

	return longest
}
