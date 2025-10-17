package main

// https://leetcode.com/problems/is-subsequence/description/

// Two pointers: can be on the same input, or can have one pointer at each input
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
