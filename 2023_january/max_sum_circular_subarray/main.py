from typing import List

class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        circular = nums + nums
        temp = [0]*len(circular)
        temp[0] = circular[0]
        maxSubarray = 0

        # kadane algorithm
        # [-3, -2, -3, -3, -2, -3]
        for leftIdx in range(0, len(nums)+1):
            for rightIdx in range(0, len(nums)):
                print(circular[leftIdx], circular[rightIdx])

            # maxSubarray = max(temp[i-1]+circular[i], circular[i])
            # temp[i] = max(temp[i-1]+circular[i], circular[i])
        
        print("circ ", circular)
        print("temp ", temp)
        return maxSubarray

if __name__ == "__main__":
    # nums, exp = [1,-2,3,-2], 3
    # nums, exp = [5,-3,5], 10
    nums, exp = [-3,-2,-3], -2
    x = Solution()
    res = x.maxSubarraySumCircular(nums)
    print(exp, res)