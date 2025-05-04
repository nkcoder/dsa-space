def ransom_note(ransom_note: str, magazine: str) -> bool:
    m_hash, r_hash = {}, {}
    for x in magazine:
        m_hash[x] = m_hash.get(x, 0) + 1

    for y in ransom_note:
        r_hash[y] = r_hash.get(y, 0) + 1

    for (k,v) in r_hash.items():
        if m_hash.get(k, 0) < v:
            return False

    return True
