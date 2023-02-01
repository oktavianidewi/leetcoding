from typing import List
class Solution:
    def cellsInRange(self, s: str) -> List[str]:
        result = []
        # in range == x < stop value 
        for x in range( ord(s[0]), (ord(s[3])+1) ):
            for y in range(ord(s[1]), (ord(s[4])+1)):
                result.append(chr(x)+chr(y))
        return result

if __name__ == "__main__":
    # s, exp = "K1:L2", ["K1","K2","L1","L2"]
    s, exp = "A1:F1", ["A1","B1","C1","D1","E1","F1"]
    x = Solution()
    res = x.cellsInRange(s)
    print(exp, res)