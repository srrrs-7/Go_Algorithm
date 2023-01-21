import itertools

def solution(N):
    cnt = 0
    for i in set(itertools.permutations(str(N))):
        perm = "".join(i)
        if str(int(perm)) == perm and int(perm) >= 0:
            cnt += 1
    return cnt