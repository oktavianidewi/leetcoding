class Solution:
    def minFlipsMonoIncr(self, s: str) -> int:
        counter_one = 0
        counter_zero = 0
        for num in s:
            if num == "1":
                counter_one += 1
            else:
                counter_zero += 1
            counter_zero = min(counter_one, counter_zero)
            
        return counter_zero

if __name__ == "__main__":
    # s, exp = "00110", 1
    s, exp = "010110", 2
    s, exp = "00011000", 2

    # punya ku pseudocodenya mirip ini https://www.youtube.com/watch?v=LvfuqUx_qw8
    # tapi gagal di test case "0101100011"
    s, exp = "0101100011", 3
    
    x = Solution()
    res = x.minFlipsMonoIncr(s)
    print(exp, res)