from typing import List
class Solution:
    # most optimal isPalindrome checking
    """
    # reverse word and check if it is equal to word
    def is_palindrome(self, word):
        return word[::-1] == word
    """
    
    def isPalindrome(self, word: str) -> bool:
        mid = len(word) >> 1
        left = 0
        right = len(word)-1
        while left <= mid :
            if word[left] != word[right-left]:
                return False
            left += 1
        return True

    def firstPalindrome(self, words: List[str]) -> str:
        res = ""
        for word in words:
            if self.isPalindrome(word):
                res = word
                break
        return res

if __name__ == "__main__":
    words = ["abc","car","ada","racecar","cool"]
    exp = "ada"
    # words = ["notapalindrome","racecar"]
    # exp = "racecar"
    s = Solution()
    res = s.firstPalindrome(words)
    print(exp, res)