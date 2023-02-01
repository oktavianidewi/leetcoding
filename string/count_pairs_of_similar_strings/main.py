from typing import List
class Solution:
    def similarPairs(self, words: List[str]) -> int:
        pairs = {}
        res = 0

        for word in words:
            # harus di-sort karena urutan string yang berbeda, bisa jadi key yang beda
            key = ''.join(sorted(set(word)))
            if key in pairs:
                pairs[key] += 1
                res += 1
            else:
                pairs[key] = 1
        # https://leetcode.com/problems/count-pairs-of-similar-strings/solutions/2923639/python3-simple-hash-map-solution/
        # find how many pairs dengan combination
        return sum(int(n*(n-1)/2) for n in pairs.values())

if __name__ == "__main__":
    words = ["aba","aabb","abcd","bac","aabc"]
    exp = 2
    words = ["aabb","ab","ba"]
    exp = 3
    # words = ["nba","cba","dba"]
    # exp = 0
    x = Solution()
    res = x.similarPairs(words)
    print(exp, res)