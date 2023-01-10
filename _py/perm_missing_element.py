# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

from functools import reduce

def solution(A):
    if len(A) == 0:
        return 0
    actual_len = len(A)+1
    missing = reduce(lambda a,b: a+b, range(actual_len+1)) - reduce(lambda a,b: a+b, A)
    # print(f"actual {reduce(lambda a,b: a+b, range(actual_len+1))}, missing {reduce(lambda a,b: a+b, A)}")
    return missing

if __name__ == "__main__" :
    A = [2, 3, 1, 5]
    exp = 4
    res = solution(A)
    print(f"expected {exp}, result {res}")