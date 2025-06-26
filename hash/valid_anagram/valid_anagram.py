from collections import Counter


# EASY: https://leetcode.com/problems/valid-anagram/description/

def valid_anagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    count_s: dict[str, int] = {}
    count_t: dict[str, int] = {}

    for i in range(len(s)):
        count_s[s[i]] = count_s.get(s[i], 0) + 1
        count_t[t[i]] = count_t.get(t[i], 0) + 1

    return count_s == count_t


def valid_anagram2(s: str, t: str) -> bool:
    return sorted(s) == sorted(t)


def valid_anagram3(s: str, t: str) -> bool:
    return Counter(s) == Counter(t)
