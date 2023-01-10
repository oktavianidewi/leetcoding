# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(X, Y, D):
    # # time complexity = o(n^2)
    # jump = 0
    # for i in range(X, Y, D):
    #     jump += 1
    
    # time complexity o(1)
    jump = (Y - X) // D
    remainder =  (Y - X) % D
    if remainder > 0 :
        jump += 1
    return jump

if __name__ == "__main__":
    X = 10
    Y = 85
    D = 30
    exp = 3
    res = solution(X, Y, D)
    print(f"expect {exp}, result {res}")