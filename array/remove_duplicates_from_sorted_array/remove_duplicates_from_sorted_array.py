# Easy:https://leetcode.com/problems/remove-duplicates-from-sorted-array
# two-pointers


def remove_duplicates_from_sorted_array(nums: list[int]):
    i, j = 0, 1

    while i < len(nums) and j < len(nums):
        if nums[i] == nums[j]:
            j = j + 1
        else:
            nums[i + 1] = nums[j]
            i = i + 1
            j = j + 1

    return i + 1

if __name__ == "__main__":
    assert remove_duplicates_from_sorted_array([1, 1, 2]) == 2
    assert remove_duplicates_from_sorted_array([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]) == 5
