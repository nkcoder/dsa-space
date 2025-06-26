# Easy: https://leetcode.com/problems/majority-element/description/


# Use hash to count each element, T-O(n), S-O(n)
def majority_element(nums: list[int]) -> int:
    hash = {}
    for n in nums:
        hash[n] = hash.get(n, 0) + 1

    for k, v in hash.items():
        if v > len(nums) / 2:
            return k

    return -1


# Sort the input, then the majority element is at index len/2, T-O(nlogn), O(1)
def majority_element_sort(nums: list[int]) -> int:
    nums_sorted = sorted(nums)
    index = int(len(nums_sorted) / 2)
    return nums_sorted[index]


# Moore voting:
# Initialize two variables: count and candidate. Set count to 0 and candidate to an arbitrary value.
# Iterate through the array nums:
# a. If count is 0, assign the current element as the new candidate and increment count by 1.
# b. If the current element is the same as the candidate, increment count by 1.
# c. If the current element is different from the candidate, decrement count by 1.
# After the iteration, the candidate variable will hold the majority element.
# T-O(n), S-O(1)
def majority_element_moore_voting(nums: list[int]) -> int:
    count, candidate = 0, 0
    for n in nums:
        if count == 0:
            candidate = n
            count += 1
        elif n == candidate:
            count += 1
        else:
            count -= 1
    return candidate

if __name__ == "__main__":
    assert majority_element([3, 2, 3]) == 3
    assert majority_element_sort([3, 2, 3]) == 3
    assert majority_element_moore_voting([3, 2, 3]) == 3
