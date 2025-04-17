class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        # with 2 pointers
        fast = 0
        slow = len(nums) - 1
        while fast <= slow:
            if nums[fast] == val: 
                nums[fast] = '_'
                temp = nums[fast]
                nums[fast] = nums[slow]
                nums[slow] = temp
                nums.pop(slow)
                slow -= 1
            else:
                fast += 1