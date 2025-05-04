# Easy: https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150


# Use two pointer, from the end to the start
# Complexity: O(m+n)-time, O(1)-space
def merge_sorted_array(nums1: list[int], m: int, nums2: list[int], n: int) -> None:
    i, j, k = m - 1, n - 1, (m + n - 1)
    while i >= 0 and j >= 0:
        if nums1[i] >= nums2[j]:
            nums1[k] = nums1[i]
            i = i - 1
        else:
            nums1[k] = nums2[j]
            j = j - 1
        k = k - 1

    if i == -1:
        while j >= 0 and k >= 0:
            nums1[k] = nums2[j]
            j = j - 1
            k = k - 1

if __name__ == "__main__":
    nums1 = [1, 2, 3, 0, 0, 0]
    m = 3
    nums2 = [2, 5, 6]
    n = 3
