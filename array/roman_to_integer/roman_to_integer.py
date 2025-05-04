# Easy: https://leetcode.com/problems/roman-to-integer/description/


## brute force, step by step
def roman_to_integer(s: str) -> int:
    hash = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    i, result = 0, 0
    while i < len(s):
        if i == len(s) - 1:
            result += hash[s[i]]
            i += 1
        elif s[i] == "I" and (s[i + 1] == "V" or s[i + 1] == "X"):
            result += hash[s[i + 1]] - hash[s[i]]
            i += 2
        elif s[i] == "X" and (s[i + 1] == "L" or s[i + 1] == "C"):
            result += hash[s[i + 1]] - hash[s[i]]
            i += 2
        elif s[i] == "C" and (s[i + 1] == "D" or s[i + 1] == "M"):
            result += hash[s[i + 1]] - hash[s[i]]
            i += 2
        else:
            result += hash[s[i]]
            i += 1

    return result


# use an extra set for substraction
def roman_to_integer2(s: str) -> int:
    hash = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    subtraction_set = set(["IV", "IX", "XL", "XC", "CD", "CM"])

    i, result = 0, 0
    while i < len(s):
        if s[i : i + 2] in subtraction_set:
            result += hash[s[i + 1]] - hash[s[i]]
            i += 2
        else:
            result += hash[s[i]]
            i += 1

    return result

if __name__ == "__main__":
    assert roman_to_integer("III") == 3
    assert roman_to_integer2("III") == 3
