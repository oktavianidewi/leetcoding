class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        if ch not in word: 
            return word
        
        stack = []
        last = 0
        for i, v in enumerate(word):
            if v != ch:
                stack.append(v)
            else:
                stack.append(v)
                last = i+1
                break
        return ''.join(stack[::-1])+word[last:]

if __name__ == "__main__":
    word = "abcdefd"
    ch = "d"
    exp = "dcbaefd"

    # word = "xyxzxe"
    # ch = "z"
    # exp = "zxyxxe"

    s = Solution()
    res = s.reversePrefix(word, ch)
    print(exp, res)