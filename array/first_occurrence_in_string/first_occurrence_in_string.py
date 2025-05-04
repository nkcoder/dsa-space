# Easy: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description
# Use substring, T-O(n), S-O(1)


def first_occurrence_in_string(haystack: str, needle: str) -> int:
    for i in range(0, len(haystack)):
        if haystack[i : i + len(needle)] == needle:
            return i
    return -1

if __name__ == "__main__":
    assert first_occurrence_in_string("hello", "ll") == 2
    assert first_occurrence_in_string("aaaaa", "bba") == -1
