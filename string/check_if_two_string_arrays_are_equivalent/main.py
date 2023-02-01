from typing import List
class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        return "".join(word1) == "".join(word2)

if __name__ == "__main__":
    word1 = ["ab", "c"]
    word2 = ["a", "bc"]
    exp = True

    # word1 = ["a", "cb"]
    # word2 = ["ab", "c"]
    # exp = False

    x = Solution()
    res = x.arrayStringsAreEqual(word1, word2)
    print(exp, res)
