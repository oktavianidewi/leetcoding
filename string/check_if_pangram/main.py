class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        # every letter of english alphabet appears at least once
        english = "abcdefghijklmnopqrstuvwxyz"
        map_char = {" ": " "}
        for char in sentence:
            map_char[char] = 1
        
        counter = 0
        for i in range(0, len(english)):
            if english[i] in map_char:
                counter += map_char[english[i]]
        
        return counter == len(english)


if __name__ == "__main__":
    sentence = "thequickbrownfoxjumpsoverthelazydog"
    exp = True

    sentence = "leetcode"
    exp = False

    x = Solution()
    res = x.checkIfPangram(sentence)
    print(exp, res)