// Package movezeros: https://leetcode.com/problems/move-zeroes/description/
// We need two pointers: one points to the next non zero element (which is the faster pointer), and the other (slow pointer) points to the next position that the non-zero element should be placed.
package movezeros

func moveZeroes(nums []int) {
	// slow pointer
	indexForTheNextNonZero := 0

	// fast pointer
	for i := range nums {
		if nums[i] != 0 {
			nums[indexForTheNextNonZero], nums[i] = nums[i], nums[indexForTheNextNonZero]
			indexForTheNextNonZero++
		}
	}
}
