# Medium: https://leetcode.com/problems/container-with-most-water/description/


# Brute force: O(n^2) time, O(1) space
# Time Limit Exceeded
def container_with_most_water0(height: list[int]) -> int:
    most_water = 0
    for i, n in enumerate(height):
        for j in range(i + 1, len(height)):
            water = (j - i) * min(height[j], n)
            if abs(water) > most_water:
                most_water = abs(water)

    return most_water


# Two pointers: O(n) time, O(1) space
def container_with_most_water(height: list[int]) -> int:
    most_water = 0
    i, j = 0, len(height) - 1
    while i < j:
        water = (j - i) * min(height[j], height[i])
        if water >= most_water:
            most_water = water
        if height[i] < height[j]:
            i += 1
        else:
            j -= 1

    return most_water

