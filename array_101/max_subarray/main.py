from typing import List
class Solution:
    # kadane algorithm
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        temp = [0]*len(nums)
        temp[0] = nums[0]
        maxSubArr = nums[0]
        for i in range(1, len(nums)):
            val = max(temp[i-1]+nums[i], nums[i])
            maxSubArr = max(maxSubArr, val)
            temp[i] = val
        return maxSubArr

if __name__ == "__main__":
    # nums, exp = [-2,1,-3,4,-1,2,1,-5,4], 6
    # nums, exp = [1], 1
    nums, exp = [5,4,-1,7,8], 23
    # nums, exp = [-1,-2], -1
    x = Solution()
    res = x.maxSubArray(nums)
    print(exp, res)