// Package fruitesintobuckets: https://leetcode.com/problems/fruit-into-baskets/description/
package fruitesintobuckets

func totalFruits(fruits []int) int {
	state := make(map[int]int)
	start, total := 0, 0

	for end := start; end < len(fruits); end++ {
		state[fruits[end]]++

		for len(state) > 2 {
			state[fruits[start]]--
			if state[fruits[start]] == 0 {
				delete(state, fruits[start])
			}
			start++
		}

		total = max(total, end-start+1)
	}

	return total
}
