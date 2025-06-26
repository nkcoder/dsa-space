def plus_one(digits: list[int]) -> list[int]:
    next = 1
    for i in range(len(digits)-1, -1, -1):
        plus_one = digits[i] + next
        if plus_one >= 10:
            next = 1
            plus_one = plus_one - 10
        else:
            next = 0
        digits[i] = plus_one

    if next == 1:
        digits.insert(0, 1)

    return digits
