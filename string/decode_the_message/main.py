from collections import defaultdict, Counter
class Solution:
    # without built-in
    def decodeMessage(self, key: str, message: str) -> str:
        # cara bikin dictionary di python
        dict_key = {" ": " "}
        alphabet = "abcdefghijklmnopqrstuvwxyz"

        z = 0
        for i, j in enumerate(key):
            if j not in dict_key:
                dict_key[j] = alphabet[z]
                z += 1

        # parse message
        decoded = ""
        for char in message:
            decoded += dict_key[char]
            
        return decoded

if __name__ == "__main__":
    key = "the quick brown fox jumps over the lazy dog"
    message = "vkbs bs t suepuv"
    exp = "this is a secret"

    key = "eljuxhpwnyrdgtqkviszcfmabo"
    message = "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
    exp = "the five boxing wizards jump quickly"
    
    x = Solution()
    res = x.decodeMessage(key, message)
    print(exp, res)