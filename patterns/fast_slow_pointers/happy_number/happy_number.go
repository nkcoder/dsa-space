// https://leetcode.com/problems/happy-number/description/
package main

import "fmt"

func isHappyHash(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}

	return n == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func isHappyFastSlowPointers(n int) bool {
	slow, fast := n, n
	for slow != 1 && fast != 1 {
		fmt.Println("slow", slow, "fast", fast)
		slow = getNext(slow)
		fast = getNext(getNext(fast))
		if slow == fast {
			return false
		}
	}

	return true
}
