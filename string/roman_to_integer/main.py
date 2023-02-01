class Solution:
    def romanToInt(self, s: str) -> int:
        map_roman = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M":1000
        }
        res = 0
        left = 0
        while left < len(s)-1:
            if map_roman[s[left]] < map_roman[s[left+1]]:
                res -= map_roman[s[left]]
            else :
                res += map_roman[s[left]]
            left += 1
        # add last value
        res += map_roman[s[-1]]
        return res

    # optimized 
    def _romanToInt(self, s: str) -> int:
        hashmap = {"I":1, "V":5, "X":10,"L":50,"C":100,"D":500,"M":1000}
        total = 0
        for i in range(len(s)):
            print(i)
            if i + 1 < len(s) and hashmap[s[i]] < hashmap[s[i + 1]]:
                total -= hashmap[s[i]]
            else:
                total += hashmap[s[i]]
        return total

if __name__ == "__main__":
    s, exp = "III", 3
    s, exp = "LVIII", 58
    s, exp = "MCMXCIV", 1994
    x = Solution()
    res = x.romanToInt(s)
    res = x._romanToInt(s)
    print(exp, res)