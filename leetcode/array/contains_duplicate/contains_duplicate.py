# Easy: https://leetcode.com/problems/contains-duplicate

def contains_duplicate(nums: list[int]) -> bool:
    hashset = set()
    for n in nums:
        if n in hashset:
            return True
        hashset.add(n)
    return False

if __name__ == "__main__":
    assert not contains_duplicate([1, 2, 3, 4])
    assert contains_duplicate([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1])
