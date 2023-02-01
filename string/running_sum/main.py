from typing import List
class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        res = list([0]*len(nums))
        res[0] = nums[0]
        for i in range(1, len(nums)):
            res[i] = res[i-1]+nums[i]
        return res

if __name__ == "__main__":
    nums = [1,2,3,4]
    exp = [1,3,6,10]

    nums = [3,1,2,10,1]
    exp = [3,4,6,16,17]

    x = Solution()
    res = x.runningSum(nums)
    print(exp, res)