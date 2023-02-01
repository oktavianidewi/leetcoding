from typing import List
class Solution:
    # 1st
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        res = 0
        for word in words:
            counter = 0
            uniq = set(word)
            for char in allowed:
                if char in uniq:
                    counter += 1
            if len(uniq) == counter:
                res += 1
        return res

    # optimal solution
    def opt_countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        c = len(words)
        for i in words:
            for j in i:
                if j not in allowed:
                    c-=1
                    break
        return c

if __name__ == "__main__":
    allowed = "ab"
    words = ["ad","bd","aaab","baa","badab"]
    exp = 2

    # allowed = "abc"
    # words = ["a","b","c","ab","ac","bc","abc"]
    # exp = 7

    x = Solution()
    res = x.countConsistentStrings(allowed, words)
    print(exp, res)