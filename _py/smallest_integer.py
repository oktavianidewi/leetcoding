# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

# 55%
def solution_dewi(A):
    # Implement your solution here
    checker = {}
    max, min = A[0], A[0]
    for i in range(0, len(A), 1):
        if A[i] > 0:
            checker[A[i]] = True
            if max <= A[i]:
                max = A[i]
            if min > A[i]:
                min = A[i]
    # print(checker, min, max)
    
    if len(checker) > 0:
        j = min
        while j <= max :
            if j not in checker:
                break
            j+=1
    else: 
        j = 1
    return j

# pigeon hole principle : https://pbedn.github.io/post/2018-05-18-codility-demo-test/
def solution(A):
    N = len(A)
    count = [0]*(N+1)
    # [0,0,0,0,0,0,0]
    for i in range(N):
        # 6 >= 1 > 0
        if N >= A[i] > 0:
            count[A[i]] += 1
    # print(count)
    for x in range(1, len(count)):
        if count[x] == 0:
            return x
    return N+1

if __name__ == "__main__":
    # A, exp = [1, 3, 6, 4, 1, 2], 5
    A, exp = [1, 2, 3], 4
    # A, exp = [-1, -3], 1
    # A, exp = [2], 1
    res = solution(A)
    print(f"expected {exp}, result {res}")