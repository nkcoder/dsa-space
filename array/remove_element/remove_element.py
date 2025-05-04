# Easy: https://leetcode.com/problems/remove-element/description/
def remove_element(nums: list[int], val: int) -> int:
    left, right = 0, len(nums) - 1
    k = 0

    while left <= right and left < len(nums) and right >= 0:
        if nums[right] == val:
            nums[right] = -val
            right = right - 1
        elif nums[left] == val:
            tmp = nums[right]
            nums[right] = -val
            nums[left] = tmp
            left = left + 1
            right = right - 1
            k = k + 1
        else:
            left = left + 1
            k = k + 1

    return k

if __name__ == "__main__":
    assert remove_element([3, 2, 2, 3], 3) == 2
    assert remove_element([0, 1, 2, 2, 3, 0, 4, 2], 2) == 5
