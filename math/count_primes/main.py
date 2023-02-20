class Solution:
    def countPrimes(self, n: int) -> int:
        # using sieve of erathosthenes
        if n <= 2 :
            return 0

        prime = [True] * n
        prime[0] = prime[1] = False
        for idx in range(2, n):
            if prime[idx]:
                for multiple in range (2*idx, n, idx):
                    prime[multiple] = False
        # print(prime)
        return sum(prime)

if __name__ == "__main__":

    n, exp = 10, 4
    n, exp = 0, 0
    n, exp = 1, 0
    x = Solution()
    res = x.countPrimes(n)
    print(exp, res)