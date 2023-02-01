from typing import List
class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        freq_ar=[0]*k
        freq_ar[0]=1
        sum=0
        count=0
        for no in range(len(nums)):
            sum+=nums[no]
            remainder=sum%k
            if remainder<0:
                remainder+=k
            count+=freq_ar[remainder]
            freq_ar[remainder]+=1
        return count
    # https://leetcode.com/problems/subarray-sums-divisible-by-k/solutions/3072318/prefixsum-easiest/


    # error kena TLE :(
    def _subarraysDivByK(self, nums: List[int], k: int) -> int:
        counter = 0
        temp = [0]*len(nums)
        for i in range(0, len(nums)):
            temp[i] = nums[i]

            # for prefix sum
            for j in range(i+1, len(nums)):
                temp[j] = temp[j-1]+nums[j]

            # counter
            for j in range(i, len(nums)):
                if temp[j] % k == 0 :
                    counter += 1
            temp = [0]*len(nums)
        return counter

if __name__ == "__main__":
    nums, k, exp = [4,5,0,-2,-3,1], 5, 7
    # nums, k, exp = [5], 9, 0
    x = Solution()
    res = x.subarraysDivByK(nums, k)
    print(exp, res)