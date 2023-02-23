class Solution:
    def shift(self, c: str, x: int) -> str:
        return chr(ord(c)+int(x))
    
    def replaceDigits(self, s: str) -> str:
        sarr = []
        for i in range(0, len(s)):
            sarr.append(s[i])
            if i % 2 == 1:
                sarr[i] = self.shift(s[i-1], s[i])
        return ''.join(sarr)


if __name__ == "__main__":
    s = "a1c1e1"
    exp = "abcdef"
    # s = "a1b2c3d4e"
    # exp = "abbdcfdhe"
    x = Solution()
    res = x.replaceDigits(s)
    print(exp, res)