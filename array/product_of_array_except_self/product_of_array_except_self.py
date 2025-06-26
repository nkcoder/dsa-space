# Medium: https://leetcode.com/problems/product-of-array-except-self/description/


# Calculate the left product and right product for each index and save to an array
# Complexity: T-O(n), S-O(n)
def product_of_array_except_self(nums: list[int]) -> list[int]:
    # calculate the left products
    left_products = [1] * len(nums)
    for i in range(1, len(nums)):
        left_products[i] = left_products[i - 1] * nums[i - 1]

    # calculate the right products
    right_products = [1] * len(nums)
    for i in range(len(nums) - 2, -1, -1):
        right_products[i] = right_products[i + 1] * nums[i + 1]

    result = [left_products[i] * right_products[i] for i in range(len(nums))]
    return result


# Calculate the left product and right index in the fly
# Complexity: T-O(n), S-O(n) (the output array does not count as extra space for space complexity analysis)
def product_of_array_except_self2(nums: list[int]) -> list[int]:
    result = [1] * len(nums)

    left = 1
    for i in range(len(nums)):
        result[i] = left
        left = left * nums[i]

    right = 1
    for i in range(len(nums) - 1, -1, -1):
        result[i] = result[i] * right
        right = right * nums[i]

    return result

if __name__ == "__main__":
    assert product_of_array_except_self([1, 2, 3, 4]) == [24, 12, 8, 6]
    assert product_of_array_except_self2([1, 2, 3, 4]) == [24, 12, 8, 6]
