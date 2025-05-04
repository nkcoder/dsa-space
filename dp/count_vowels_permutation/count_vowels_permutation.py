# https://leetcode.com/problems/count-vowels-permutation/description/
# rules:
## a -> e
## e -> a, i
## i -> a, e, o, u
## o -> i, u
## u -> a


def count_vowels_permutation_hash(n: int) -> int:
    mod = 10 ** 9 + 7
    dp = { v: 1 for v in "aeiou"}

    for _ in range(2, n + 1):
        new_dp = {}
        new_dp['a'] = (dp['e'] + dp['u'] + dp['i']) % mod
        new_dp['e'] = (dp['a'] + dp['i']) % mod
        new_dp['i'] = (dp['e'] + dp['o']) % mod
        new_dp['o'] = dp['i']
        new_dp['u'] = (dp['o'] + dp['i']) % mod
        dp = new_dp

    return sum(dp.values()) % mod


def count_vowels_permutation_array(n: int) -> int:
    mod = 10 ** 9 + 7
    # initialize dp arrays for the current and previous states
    prev = [1] * 5  # base case for n = 1: ['a', 'e', 'i', 'o', 'u']

    for _ in range(2, n + 1):
        cur = [0] * 5
        cur[0] = (prev[1] + prev[2] + prev[4]) % mod # 'a' can come from 'e', 'i', 'u'
        cur[1] = (prev[0] + prev[2]) % mod # 'e' can come from 'a', 'i'
        cur[2] = (prev[1] + prev[3]) % mod # 'i' can come from 'e', 'o'
        cur[3] = (prev[2]) % mod # 'o' can come ffrom 'i'
        cur[4] = (prev[2] + prev[3]) % mod # 'u' can come from 'i', 'o'
        prev = cur

    return sum(prev) % mod

if __name__ == "__main__":
    assert count_vowels_permutation_hash(1) == 5
    assert count_vowels_permutation_array(1) == 5
