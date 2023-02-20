from typing import List
class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        i = 0
        self.helper(0, flowerbed, n)
        print("hasilnya n ", n)
    
    def helper(self, i, flowerbed, n):
        if i == len(flowerbed):
            return n
        
        if flowerbed[i] == 0:
            # replace 0 -> 1
            flowerbed[i] = 1

        print("HALO ", i)
        if flowerbed[i] == 1:
            if i+2 < len(flowerbed):
                self.helper(i+2, flowerbed, n-1)
            else:
                return n
        else:
            self.helper(i+1, flowerbed, n)

        

if __name__ == "__main__":
    flowerbed, n, exp = [1,0,0,0,1], 1, True
    # flowerbed, n, exp = [1,0,0,0,1], 2, False
    x = Solution()
    res = x.canPlaceFlowers(flowerbed, n)
    print(exp, res)