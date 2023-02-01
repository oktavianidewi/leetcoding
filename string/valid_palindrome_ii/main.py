from collections import defaultdict
class Solution:
    def validPalindrome(self, s: str) -> bool:
        map = defaultdict()
        for char in s:
            if char in map:
                map[char] += 1
            else: 
                map[char] = 1
        # print(map)
        counter_one = 0
        for n in map:
            if map[n] == 1:
                counter_one += 1
        if counter_one == 1 or counter_one == 2:
            return True
        return False

if __name__ == "__main__":
    s, exp = "aba", True
    s, exp = "abca", True
    s, exp = "abc", False
    x = Solution()
    res = x.validPalindrome(s)
    print(exp, res)
