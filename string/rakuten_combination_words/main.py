from typing import List, Counter
from collections import defaultdict
from itertools import permutations

class Solution:
    def wordCombination(self, sentence: str) -> int:
        y = permutations(sentence)
        words = [''.join(i) for i in y]

        res = defaultdict()
        for word in words:
            if self.isValid(word):
                res[word] = 1
        return len(res)

    def isValid(self, word: str) -> bool:
        vowels = {"A": True, "I": True, "U": True, "E": True, "O": True}

        if word[0] in vowels:
            return False

        for i in range (0, len(word)-1):
            # even index should contain consonant
            if i % 2 == 0 and word[i] in vowels:
                return False
            if i % 2 == 1 and word[i] not in vowels:
                return False
        return True

if __name__ == "__main__":
    s, exp = "BAR", 2
    # s, exp = "AABB", 1
    # s, exp = "AABCY", 6
    x = Solution()
    res = x.wordCombination(s)
    print(exp, res)
  