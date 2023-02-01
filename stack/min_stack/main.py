class MinStack:

    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.stack.append(val)

        if len(self.min_stack) == 0:
            self.min_stack.append(val)
        else:
            self.min_stack.append( min(self.min_stack[-1], val) )

    # def show(self) -> None:
    #     print("stack ", self.stack)
    #     print("min_stack ", self.min_stack)
    #     print("\n")

    def pop(self) -> None:
        self.stack.pop()
        self.min_stack.pop()

    def top(self) -> int:
        return self.stack[len(self.stack)-1]
        

    def getMin(self) -> int:
        return self.min_stack[-1]
        
if __name__ == "__main__":
    # ["MinStack","push","push","push","getMin","pop","top","getMin"]
    # [[],[-2],[0],[-3],[],[],[],[]]
    # Output
    # [null,null,null,null,-3,null,0,-2]

    minStack = MinStack()
    minStack.push(-2)
    print(minStack.show())
    
    minStack.push(0)
    print(minStack.show())

    minStack.push(-3)
    print(minStack.show())

    # return -3
    x = minStack.getMin()
    print("getMin ", x)

    print("pop")
    minStack.pop()
    minStack.show()

    # return 0
    x = minStack.top()
    print("top ", x)

    # return -2
    x = minStack.getMin()
    print("getMin ",x)

