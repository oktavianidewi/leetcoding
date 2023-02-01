from typing import List, Counter
from collections import defaultdict

class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        max = 0
        for i in range(0, len(sentences)):
            words = len(sentences[i].split(" "))
            if max < words: 
                max = words
        return max

if __name__ == "__main__":
    sentences, exp = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"], 6
    sentences, exp = ["please wait", "continue to fight", "continue to win"], 3
    x = Solution()
    res = x.mostWordsFound(sentences)
    print(exp, res)