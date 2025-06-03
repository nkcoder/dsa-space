package org.nkcoder.array_hash

/** Easy: https://leetcode.com/problems/two-sum/
  *   - brute force: T-O(n2), S-O(n)
  *   - map and sort, then two pointers: T-O(nlogn), S-O(n)
  *   - search by (target - value), T-O(nlogn), S-O(n)
  */
object TwoSum:
  def twoSum(nums: Array[Int], target: Int): Array[Int] =
    val numsMap = nums.zipWithIndex.toMap
    val index =
      nums.indices.indexWhere(i => numsMap.get(target - nums(i)).exists(_ != i))
    if index == -1 then Array(-1, -1)
    else Array(index, numsMap(target - nums(index)))
