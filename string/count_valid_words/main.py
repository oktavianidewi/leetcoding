class Solution:
    def countValidWords(self, sentence: str) -> int:
        words = sentence.split(" ")
        valid = 0
        for i in words:
            if self.isInvalidWords(i):
                print(i, self.isInvalidWords(i))
                valid+=1
        return valid
    
    def isInvalidWords(self, word: str) -> bool:
        valChar = 0
        for char in word:
            if (char >= "a" and char <= "z"):
                valChar += 1
            else: 
                valChar -= 1


        return valChar == len(word)

if __name__ == "__main__":
    # sentence, exp = "cat and dog", 3
    sentence, exp = "!this  1-s b8d!", 0
    # sentence, exp = "alice and  bob are playing stone-game10", 5
    x = Solution()
    res = x.countValidWords(sentence)
    print(exp, res)