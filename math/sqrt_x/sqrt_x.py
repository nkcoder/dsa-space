def my_sqrt(x: int) -> int:
    i = 0
    while i <= (x+1) // 2 and i * i <= x:
        i += 1

    return i - 1
