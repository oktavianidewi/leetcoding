class Solution:
    def maxDepth(self, s: str) -> int:
        pass

if __name__ == "__main__":
    s = "(1+(2*3)+((8)/4))+1"
    exp = 3

    s = "(1)+((2))+(((3)))"
    exp = 3

    x = Solution()
    x.maxDepth(s)