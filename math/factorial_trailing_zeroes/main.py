class Solution:
    def trailingZeroes(self, n: int) -> int:
        if n == 0:
            return 0

        # x = self.factorial(n)
        x = 100
        counter = 0
        print("factorial ", x)
        while x > 0:
            if x % 10:
                print("habis " , x)
                counter += 1
            x //= 10
        return counter

    def factorial(self, n: int) -> int:
        if n == 0:
            return 1
        return n * self.factorial(n-1)

if __name__ == "__main__":
    # tricky: https://leetcode.com/problems/factorial-trailing-zeroes/solutions/522315/clear-explanation-of-the-solution-since-i-didn-t-find-an-adequate-one/?languageTags=python
    n, exp = 3, 0
    n, exp = 5, 1
    # n, exp = 0, 0
    x = Solution()
    res = x.trailingZeroes(n)
    print(exp, res)