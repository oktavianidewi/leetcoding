from typing import List
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        threshold = len(nums) >> 1
        maps = {}
        for i in nums:
            if i in maps:
                maps[i] += 1
                if maps[i] > threshold:
                    return i
            else:
                maps[i] = 1
        return maps[0]

if __name__ == "__main__":
    nums, exp = [3,2,3], 3
    # nums, exp = [2,2,1,1,1,2,2], 2

    x = Solution()
    res = x.majorityElement(nums)
    print(exp, res)