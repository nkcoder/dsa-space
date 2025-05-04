def is_happy(n: int) -> bool:
    sum_set = set()
    sum = get_square_sum(n)
    while sum != 1:
        sum_set.add(sum)
        new_sum = get_square_sum(sum)
        if new_sum in sum_set:
            return False
        sum = new_sum

    return True


def get_square_sum(n: int) -> int:
    s = str(n)
    sum = 0
    for c in s:
        sum += int(c) * int(c)
    return sum
