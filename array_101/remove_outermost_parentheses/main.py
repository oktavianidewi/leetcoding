class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        stack = []
        result = []
        counter = 0
        for i in s:
            if i == "(":
                counter -= 1
            elif i == ")":
                counter += 1
            stack.append(i)

            if counter == 0:
                result += stack[1:-1]
                stack = []
        
        return ''.join(result)


if __name__ == "__main__":
    # s = "(()())(())"
    # exp = "()()()"

    # s = "(()())(())(()(()))"
    # exp = "()()()()(())"

    s = "()()"
    exp = ""

    x = Solution()
    res = x.removeOuterParentheses(s)
    print(exp, res)