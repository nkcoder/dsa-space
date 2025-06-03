# Medium: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
# Requirement: Your solution must use only constant extra space.


# Use two pointers: left, right and compare the sum with target.
# Complexity: T-O(n), M-O(1)
def two_sum_input_sorted(numbers: list[int], target: int) -> list[int]:
    left, right = 0, len(numbers) - 1
    while left < right:
        sum = numbers[left] + numbers[right]
        if sum == target:
            return [left + 1, right + 1]
        elif sum < target:
            left += 1
        else:
            right -= 1
