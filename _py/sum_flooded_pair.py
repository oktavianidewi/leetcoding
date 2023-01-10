# class Solution:
#     def sumOfFlooredPairs(self, nums: List[int]) -> int:
#         sumP = 0 #To store the value of Sum of floor values
#         for i in nums: #Traverse every element in nums
#             for j in nums: #Traverse every element in nums
#                 sumP += (j//i) #Simply do floor division and add the number to sumP
#         return sumP % (10**9 +7)#return the sumof the pairs


# : List[int]
class Solution:
    def sumOfFlooredPairs(self, nums) -> int:
        maxi = max(nums) + 1
        dic = {}
        prefix=[0]*maxi
        print(">> prefix ", prefix)
        
        # utilize the frequency map
        for i in nums:
            if i in dic:
                dic[i] += 1
            else:
                dic[i] = 1
        print(">> dic ", dic)
        
        for i in range(1,maxi):
            print(">> isi i ", i)
            if i not in dic:
                # jika i tidak ditemukan dalam dic, maka prefix[1] = prefix[0]; prefix[2] = prefix[1]
                prefix[i] = prefix[i-1]
            else:
                print(">> isi prefix[i-1] ", prefix[i-1]+dic[i])
                prefix[i] = prefix[i-1]+dic[i]
        print(prefix,dic)
        sumP = 0
        for i in set(nums):
            print("= i", i)
            for j in range(i,maxi,i):
                print("=== j", j)
                sumP += dic[i]*(prefix[-1]-prefix[j-1])
                print(sumP,end = " ")
        return sumP % (10**9 +7)

def main():
    # nums = [2,5,9]
    nums = [1, 2, 3, 4, 5]
    x = Solution()
    x.sumOfFlooredPairs(nums=nums)

if __name__ == "__main__":
    main()