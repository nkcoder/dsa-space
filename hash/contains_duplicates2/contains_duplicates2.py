from collections import defaultdict


def containsNearbyDuplicate(nums: list[int], k: int) -> bool:
    num_indices: defaultdict[int, list[int]] = defaultdict(list)
    for i, n in enumerate(nums):
        num_indices[n].append(i)

    for value in num_indices.values():
        if len(value) == 1:
            continue

        for i in range(len(value) - 1):
            if value[i + 1] - value[i] <= k:
                return True

    return False


if __name__ == "__main__":
    assert containsNearbyDuplicate([1, 2, 3, 1], 3)
    assert containsNearbyDuplicate([1, 0, 1, 1], 1)
    assert not containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2)
