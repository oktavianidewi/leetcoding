class Solution:
    def truncateSentence(self, s: str, k: int) -> str:
        return " ".join(s.split(" ")[:k])

if __name__ == "__main__":
    s = "Hello how are you Contestant"
    k = 4
    exp = "Hello how are you"
    
    x = Solution()
    res = x.truncateSentence(s, k)
    
    print(exp, res)