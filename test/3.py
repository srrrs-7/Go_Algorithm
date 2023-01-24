def solution(A):
    n = len(A)
    if (n <= 1): return 0
    result = 0
    for i in range(n - 1):
        # if adjacent is same, result+1 and i step 2
        if (A[i] == A[i + 1]):
            result = result + 1
            i = i + 1
    if (result == n - 1):
        return result - 1

    r = 0
    for i in range(n):
        count = 0
        if (i > 0):
            if (A[i - 1] != A[i]):
                count = count + 1
            else:
                count = count - 1
        if (i < n - 1):
            if (A[i + 1] != A[i]):
                count = count + 1
            else:
                count = count - 1
        r = max(r, count)
    return result + r