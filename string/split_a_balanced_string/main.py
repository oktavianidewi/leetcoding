class Solution:
    def balancedStringSplit(self, s: str) -> int:
        counter = 0
        temp = 0
        for i in range(0, len(s)):
            if s[i] == "R":
                temp += 1
            else:
                temp -= 1
            
            if temp == 0:
                counter += 1

        return counter

if __name__ == "__main__":
    s, exp = "RLRRLLRLRL", 4
    # s, exp = "RLRRRLLRLL", 2
    # s, exp = "LLLLRRRR", 1
    x = Solution()
    res = x.balancedStringSplit(s)
    print(exp, res)