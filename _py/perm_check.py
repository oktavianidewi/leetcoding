# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A):
    # Implement your solution here
    perm_checker = {}
    i = 0
    max = A[0]
    while i < len(A):
        if A[i] in perm_checker:
            # if multiple number exist 
            return 0
        else:
            perm_checker[A[i]] = 1
        
        if max <= A[i]:
            max = A[i]

        i += 1
    
    if len(perm_checker) != max:
        return 0
    
    return 1

if __name__ == "__main__":
    A = [4,1,3,2]
    exp = 1

    A = [4,1,3]
    exp = 0
    
    res = solution(A)
    print(f"expected {exp}, result {res}")