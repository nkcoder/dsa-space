// https://leetcode.com/problems/valid-palindrome/description/
//

// It works, but runtime is not the best because unicode conversions are expensive.
package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		ri, rj := rune(s[i]), rune(s[j])
		if !unicode.IsLetter(ri) && !unicode.IsNumber(ri) {
			i++
			continue
		}

		if !unicode.IsLetter(rj) && !unicode.IsNumber(rj) {
			j--
			continue
		}

		if unicode.ToLower(ri) != unicode.ToLower(rj) {
			return false
		}

		i++
		j--
	}
	return true
}

func isPalindromeOptimized(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphaNum(s[i]) {
			i++
		}

		for i < j && !isAlphaNum(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
