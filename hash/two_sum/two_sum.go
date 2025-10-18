// https://leetcode.com/problems/two-sum/description/
//

package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums {
		j, ok := hash[target-v]
		if ok {
			return []int{i, j}
		} else {
			hash[v] = i
		}
	}
	return []int{}
}
