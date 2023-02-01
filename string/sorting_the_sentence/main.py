class Solution:
    def sortSentence(self, s: str) -> str:
        res = [""]*len(s)
        words = s.split(" ")
        for j in words:
            ix = int(j[len(j)-1]) - 1
            res[ix] = j[0:len(j)-1]
        return " ".join(res).lstrip().rstrip()

if __name__ == "__main__":
    s = "is2 sentence4 This1 a3"
    exp = "This is a sentence"

    # s = "Myself2 Me1 I4 and3"
    # exp = "Me Myself and I"

    x = Solution()
    res = x.sortSentence(s)
    print(exp, res)