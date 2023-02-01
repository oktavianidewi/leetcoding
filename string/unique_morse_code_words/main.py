from typing import List

class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        if len(words) == 1:
            return 1

        alphabet = [chr(x) for x in range(97, 123)]
        morse = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        morse_alphabet = dict(zip(alphabet, morse))

        res = set()
        for word in words:
            word_in_morse = ""
            for char in word:
                word_in_morse += morse_alphabet[char]
            res.add(word_in_morse)
        return len(res)


if __name__ == "__main__":
    words, exp = ["gin","zen","gig","msg"], 2
    # words, exp = ["a"], 1
    x = Solution()
    res = x.uniqueMorseRepresentations(words)
    print(exp, res)