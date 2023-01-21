def solution(A):
    start = 0
    end = start + 1
    longest = 0

    A.sort()
    while end <= len(A):
        tmp = A[start:end]
        if amplitude(tmp) < 2:
            end += 1
            if len(tmp) > longest:
                longest = len(tmp)
        else:
            start = end
            end = start + 1
    return longest

def amplitude(tmp):
    if len(tmp) < 2:
        return 0
    else:
        return max(tmp) - min(tmp)