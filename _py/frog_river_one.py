# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(X, A):
    i = 0
    dict_temp = {}
    while i < len(A):
        dict_temp[A[i]] = i
        print(dict_temp)
        # harus ada x (5) distinct key in dictionary, supaya frog bisa lompat
        if len(dict_temp) == X:
            # return index
            return i
        i += 1
    return -1

if __name__ == "__main__":
    A = [1,3,1,4,4,3,5,4]
    X = 5
    exp = 6
    res = solution(X, A)
    print(f"expected {exp}, result {res}")