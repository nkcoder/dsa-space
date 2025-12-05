// Package findtheduplicatenumber finds the duplicate number in an slice
package findtheduplicatenumber

// https://leetcode.com/problems/find-the-duplicate-number/
// You must solve the problem without modifying the input array and using only constant extra space

/**
* We need to think the array as a linked list (with loop). The value of the array element means it points to the next element.
* For example: 3, 1, 4, 2, 3, linked list: 3 -> 2 -> 4 -> 3 -> 2
 */
func findDuplicate(nums []int) int {

	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// Potentially might be faster
func findDuplicate2(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// this is not: slow = nums[0]
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
