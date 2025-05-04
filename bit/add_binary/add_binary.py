def add_binary(a: str, b: str) -> str:
    i = len(a) - 1
    j = len(b) - 1
    carry = 0
    result: list[str] = []
    while i >= 0 and j >= 0:
        sum = int(a[i]) + int(b[j]) + carry
        current = str(sum)
        if sum >= 2:
            carry = 1
            current = str(sum - 2)
        else:
            carry = 0

        result.append(current)
        i -= 1
        j -= 1

    while i >= 0:
        sum = int(a[i]) + carry
        current = str(sum)
        if sum >= 2:
            carry = 1
            current = str(sum - 2)
        else:
            carry = 0
        result.append(current)
        i -= 1

    while j >= 0:
        sum = int(b[j]) + carry
        current = str(sum)
        if sum >= 2:
            carry = 1
            current = str(sum - 2)
        else:
            carry = 0
        result.append(current)
        j -= 1


    if carry == 1:
        result.append("1")

    result.reverse()
    return "".join(result)

if __name__ == "__main__":
    assert add_binary("11", "1") == "100"
    assert add_binary("1010", "1011") == "10101"
