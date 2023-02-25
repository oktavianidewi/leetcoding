from typing import List
class Solution:
    def isPrefix(self, word: str, pref: str) -> bool:
        if len(pref) > len(word):
            return False
        i = 0
        while i < len(pref):
            if pref[i] != word[i]:
                return False
            i += 1
        return True

    # optimal solution: prefix always starts from beginning, we can use array of prefix length to compare
    def prefixCount(self, words: List[str], pref: str) -> int:
        counter = 0
        for word in words:
            if self.isPrefix(word, pref):
                counter += 1
        return counter

if __name__ == "__main__":
    words = ["pay","attention","practice","attend"]
    pref = "at"
    exp = 2
    # words = ["leetcode","win","loops","success"]
    # pref = "code"
    # exp = 0
    s = Solution()
    res = s.prefixCount(words, pref)
    print(exp, res)