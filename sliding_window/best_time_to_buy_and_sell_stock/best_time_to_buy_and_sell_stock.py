# Easy: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/


# Use two pointers (sliding window)
# Complexity: T-O(n), S-O(1)
def best_time_to_buy_and_sell_stock(prices: list[int]) -> int:
    left, max_profit = 0, 0
    for right in range(1, len(prices)):
        if prices[left] > prices[right]:
            left = right
        else:
            if prices[right] - prices[left] > max_profit:
                max_profit = prices[right] - prices[left]

    return max_profit


# Brute force: double loop, will get `Time Limit Exceeded` error
# Complexity: T-O(n^2), S-O(1)
def best_time_to_buy_and_sell_stock_exceed_time(prices: list[int]) -> int:
    max_profit = 0
    for i in range(len(prices)):
        for j in range(i, len(prices)):
            profit = prices[j] - prices[i]
            if profit > max_profit:
                max_profit = profit

    return max_profit
