# Medium: https://leetcode.com/problems/3sum/description/


# Option 1: sort => for each number => two sum (target - number)
# Complexity: O(n^2) time, O(n) space
# nums.sort() might be more efficient than sorted(nums) since it's in-place
def three_sum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    result: list[list[int]] = []

    for i, n in enumerate(nums):
        # skip duplicates in the same position
        if i > 0 and n == nums[i - 1]:
            continue

        # two sum with sorted input
        i, j = i + 1, len(nums) - 1
        while i < j:
            sum = n + nums[i] + nums[j]
            if sum == 0:
                result.append([n, nums[i], nums[j]])
                i += 1
                j -= 1
                while i < j and nums[i] == nums[i - 1]:
                    i += 1
                while i < j and nums[j] == nums[j + 1]:
                    j -= 1
            elif sum < 0:
                i += 1
            else:
                j -= 1

    return result
