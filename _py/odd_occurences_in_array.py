# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A):
    dictvar = {}
    for i in range (0, len(A), 1):
        if A[i] in dictvar:
            dictvar[A[i]] += 1
        else: 
            dictvar[A[i]] = 1

    for i in dictvar:
        if dictvar[i]%2 > 0:
            return i
    return 0

if __name__ == "__main__":
    A = [9,3,9,3,9,7,9]
    exp = 7
    res = solution(A)
    print(f"expect {exp}, result {res}")