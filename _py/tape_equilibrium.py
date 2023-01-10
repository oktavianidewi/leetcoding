# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

from functools import reduce

def solution(A):
    # Implement your solution here
    a = 0
    b = reduce(lambda a, b: a+b, A)
    minval = -1
    for i in A:
        a+=i
        b-=i
        diff = -1*(a-b) if a - b < 0 else (a-b)
        if minval < 0 or minval > diff:
            minval = diff
    return minval


if __name__ == "__main__":
    # wrong case:
    # The following issues have been detected: wrong answers.
    # For example, for the input [-1000, 1000] the solution returned a wrong answer (got 0 expected 2000).
    
    A = [3, 1, 2, 4, 3]
    exp = 1
    res = solution(A)
    print(f"expected {exp}, result {res}")