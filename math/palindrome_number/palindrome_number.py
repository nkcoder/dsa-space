def is_palindrome(x: int) -> bool:
    x_str = str(x)
    i, j = 0, len(x_str) - 1
    while i < j:
        if x_str[i] != x_str[j]:
            return False
        i += 1
        j -= 1

    return True

def is_palindrome_2(x: int) -> bool:
    if x < 0:
        return False

    x_copy = x
    palindrome = 0
    while x_copy != 0:
        t = x_copy % 10
        palindrome = palindrome * 10 + t
        x_copy = x_copy // 10

    return x == palindrome
