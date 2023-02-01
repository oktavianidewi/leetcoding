class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        pass

if __name__ == "__main__":
    allowed = "ab"
    words = ["ad","bd","aaab","baa","badab"]
    exp = 2

    allowed = "abc"
    words = ["a","b","c","ab","ac","bc","abc"]
    exp = 7

    allowed = "cad"
    words = ["cc","acd","b","ba","bac","bad","ac","d"]
    exp = 4

    x = Solution()
    res = x.countConsistentStrings(allowed, words)
    print(exp, res)