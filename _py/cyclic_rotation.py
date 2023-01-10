# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A, K):
    
    if len(A) == 0:
        return A
    if K % len(A) == 0: 
        return A
    
    K = K % len(A)
    
    temp = [0]*len(A)
    for i in range (0, len(A), 1):
        print(f"index{(i+K)%len(A)}")
        temp[(i+K)%len(A)] = A[i]
    return temp

if __name__ == "__main__" : 
    # A = [3, 8, 9, 7, 6]
    # K = 3
    # exp = [9, 7, 6, 3, 8]

    # A = [1, 2, 3, 4]
    # K = 4
    # exp = [1, 2, 3, 4]

    # # error
    # A = [1, 2]
    # K = 1
    # exp = [2, 1]

    A = []
    K = 3
    exp = [6, 7, 1, 2, 3, 4, 5]

    res = solution(A, K)
    print(f"expected {exp}, result {res}")