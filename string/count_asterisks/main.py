class Solution:
    def countAsterisks(self, s: str) -> int:
        bars = s.split("|")

        countAsterisk = 0
        for i in range(0, len(bars), 2):
            for char in bars[i]:
                if char == '*':
                    countAsterisk += 1
        return countAsterisk
    
    # optimized
    def _countAsterisks(self, s: str) -> int:
        count_on = ans = 0
        for char in s:
            if char == '|':
                count_on += 1
            else:
                if count_on % 2 == 0 and char == '*':
                    ans += 1
        return ans


if __name__ == "__main__":
    s, exp = "l|*e*et|c**o|*de|", 2
    # s, exp = "iamprogrammer", 0
    # s, exp = "yo|uar|e**|b|e***au|tifu|l", 5
    x = Solution()
    res = x.countAsterisks(s)
    print(res, exp)