class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        pass

if __name__ == "__main__":
    str1 = "ABCABC"
    str2 = "ABC"
    exp = "ABC"

    str1 = "ABABAB"
    str2 = "ABAB"
    exp = "AB"

    str1 = "LEET"
    str2 = "CODE"
    exp = ""

    x = Solution()
    res = x.gcdOfStrings(str1, str2)
    print(exp, res)