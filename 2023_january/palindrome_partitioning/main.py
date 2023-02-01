from typing import List
from collections import defaultdict
from itertools import combinations
class Solution:
    def partition(self, s: str) -> List[List[str]]:
        if len(s) == 1:
            return [[s]]
        
        possible_combinations = defaultdict()
        # generate all possible combination of words
        # isPalindrome =  true
        # aba
        # ccc
        # casac
        
        for i in range(1, len(s)+1):
            print(i)
            for x in combinations(s, i):
                print(set(x))
                
        # print(combs)
        return []

    def idPalindrome(self, word: str) -> bool: 
        return True

if __name__ == "__main__":
    s, exp = "aab", [["a","a","b"],["aa","b"]]
    # s, exp = "a", [["a"]]
    x = Solution()
    res = x.partition(s)
    print(exp, res)
