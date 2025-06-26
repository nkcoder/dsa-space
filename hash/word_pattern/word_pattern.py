def word_pattern(pattern: str, s: str) -> bool:
    words = s.split(" ")
    length = len(words)

    if len(pattern) != length:
        return False

    mapping = {}

    for i in range(length):
        p = pattern[i]
        if p in mapping:
            if mapping[p] != words[i]:
                return False
        else:
            if words[i] in mapping.values():
                return False
            mapping[p] = words[i]

    return True
