
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i = 0
        j = 1
        while j <= len(nums) - 1:
            if nums[i] == nums[j]:
                j += 1
            else:
                temp = i + 1
                nums[temp] = nums[j]
                j += 1
                i += 1
                
        return i+1