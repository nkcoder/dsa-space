# Medium: https://leetcode.com/problems/group-anagrams/description/


# Count each letter in each string => as key for the hash
# Complexity: T-O(m*n), S-O(n), n is the size of the input data, m is the average length of the input str
def group_anagrams(strs: list[str]) -> list[list[str]]:
    hash: dict[str, list[str]] = {}

    for s in strs:
        count = [0] * 26  # 26 lower case letters
        for c in s:
            count[ord(c) - ord("a")] += 1

        # don't use "" to join the count because: [1,0,10,0] == [10,1,0,0]
        key = ".".join(str(x) for x in count)

        if key in hash:
            hash[key].append(s)
        else:
            hash[key] = [s]

    return list(hash.values())


# Sort each string => as the key of the hash
# Complexity: T-O(n*mlogm), S-O(n), n is the size of the input list, m is the average size of each input string
def group_anagrams2(strs: list[str]) -> list[list[str]]:
    hash: dict[str, list[str]] = {}
    for s in strs:
        sorted_s = "".join(sorted(s))
        if sorted_s in hash:
            hash[sorted_s].append(s)
        else:
            hash[sorted_s] = [s]

    return list(hash.values())

if __name__ == "__main__":
    assert group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]) == [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
    assert group_anagrams2(["eat", "tea", "tan", "ate", "nat", "bat"]) == [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
