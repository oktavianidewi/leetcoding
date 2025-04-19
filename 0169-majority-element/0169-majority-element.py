class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        x = 0
        
        temp = dict()
        for v in nums:
            if v in temp:
                temp[v] += 1
            else:
                temp[v] = 1
        
        mid = len(nums)/2
        for k, v in temp.items():
            if v > mid:
                x = k
                break
        
        return x