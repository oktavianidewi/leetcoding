# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(S):
    # Implement your solution here
    vocal_dict = {"A":True, "I":True, "U":True, "E":True, "O":True}
    vocal_counter = 0
    distinct_dict = {}
    for x in range(0, len(S)) :
        if S[x] in distinct_dict: 
            distinct_dict[S[x]] += 1
        else:
            distinct_dict[S[x]] = 1

        if S[x] in vocal_dict:
            vocal_counter += 1
    consonant_counter = len(S) - vocal_counter
    x = factorial(len(S))//(factorial(vocal_counter)*factorial(consonant_counter))-len(distinct_dict)
    return x

def factorial(x):
    if x == 1:
        return 1
    else:
        return (x * factorial(x-1))

if __name__ == "__main__":
    # S = "BAR"
    S = "AABCY"
    res = solution(S)
    print(res)