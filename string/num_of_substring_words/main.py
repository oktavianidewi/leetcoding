from typing import List
class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:
        count = 0
        for i in patterns:
            if i in word:
                count += 1
        return count

if __name__ == "__main__":
    patterns = ["a","abc","bc","d"]
    word = "abc"
    exp = 3
    patterns = ["a","b","c"]
    word = "aaaaabbbbb"
    exp = 2
    s = Solution()
    res = s.numOfStrings(patterns, word)
    print(exp, res)