class Solution:
    def maxDepth(self, s: str) -> int:
        counter = 0
        max = 0
        for i, v in enumerate(s):
            if i == 0 and v == ")":
                return -1
            
            if v == "(":
                counter += 1
            elif v == ")":
                counter -= 1
            
            if max < counter:
                max = counter
        
        # print(counter, max)
        if counter > 0:
            return -1
        
        return max

if __name__ == "__main__":
    s = "(1+(2*3)+((8)/4))+1"
    exp = 3

    # s = "(1)+((2))+(((3)))"
    # exp = 3

    # s = ")("
    # exp = -1

    s = "(()"
    exp = -1

    x = Solution()
    res = x.maxDepth(s)
    print(exp, res)