# Medium: https://leetcode.com/problems/top-k-frequent-elements/description/


# use hash for count => use list for reversed count, no need to sort
# similar to bucket sort
# complexity: T-O(n), S-O(n)
from collections import defaultdict
def top_k_frequent(nums: list[int], k: int) -> list[int]:
    num_count: dict[int, int] = {}
    for n in nums:
        num_count[n] = num_count.get(n, 0) + 1

    frequency: list[list[int]] = [[] for _ in range(len(nums) + 1)]
    for n, c in num_count.items():
        frequency[c].append(n)

    result: list[int] = []
    for i in range(len(frequency) - 1, 0, -1):
        if frequency[i]:
            result.extend(frequency[i])
        if len(result) >= k:
            return result[:k]
    return result[:k]


# use hash for count => reverse count => sort
# complexity: T-O(nlogn), S-O(n)
def top_k_frequent2(nums: list[int], k: int) -> list[int]:
    num_count: dict[int, int] = {}

    for n in nums:
        num_count[n] = num_count.get(n, 0) + 1

    num_count_reversed: defaultdict[int, list[int]] = defaultdict(list)
    for n, c in num_count.items():
        num_count_reversed[c].append(n)

    count_sorted = sorted(list(num_count_reversed.keys()), reverse=True)

    result: list[int] = []
    i = 0
    while len(result) < k and i < len(count_sorted):
        data: list[int] = num_count_reversed[count_sorted[i]]
        result.extend(data)
        i += 1

    return result[:k]
