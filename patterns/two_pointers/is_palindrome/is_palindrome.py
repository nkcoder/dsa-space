# https://leetcode.com/problems/valid-palindrome/description/
def is_palindrome(s: str) -> bool:
    str_removed_non_alpha_num: list[str] = []
    for c in s:
        if c.isalnum():
            str_removed_non_alpha_num.append(c.lower())

    i, j = 0, len(str_removed_non_alpha_num) - 1
    while i < j:
        if str_removed_non_alpha_num[i] != str_removed_non_alpha_num[j]:
            return False
        i += 1
        j -= 1

    return True
