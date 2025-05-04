# Easy: https:// leetcode.com/problems/two-sum/


# search the complement => build the nums map
def two_sum(nums: list[int], target: int) -> bool:
    prev_map = {}  # value: index
    for i, v in enumerate(nums):
        complement = target - v
        if complement in prev_map:
            return [prev_map[complement], i]
        prev_map[v] = i

    return


# initialize nums map => search the complement
def two_sum2(nums: list[int], target: int) -> list[int]:
    nums_map = {nums[i]: i for i in range(len(nums))}

    for i in range(len(nums)):
        complement = target - nums[i]
        if complement in nums_map and nums_map[complement] != i:
            return [i, nums_map[complement]]

    return [-1, -1]


# initialize nums map => sort => search
def two_sum3(nums: list[int], target: int) -> list[int]:
    # value => index, the index is a set, because we might the same element twice in the input, for example: [3, 3]
    nums_map: dict[int, set[int]] = {}

    for i in range(len(nums)):
        if nums[i] in nums_map:
            nums_map.get(nums[i]).add(i)
        else:
            nums_map[nums[i]] = {i}

    nums.sort()
    i, j = 0, len(nums) - 1

    while i <= j:
        tmp_sum = nums[i] + nums[j]
        if tmp_sum == target:
            if nums_map[nums[i]] == nums_map[nums[j]]:
                return list(nums_map[nums[i]])
            return list([nums_map[nums[i]].pop(), nums_map[nums[j]].pop()])
        elif tmp_sum > target:
            j -= 1
        else:
            i += 1
