from typing import List
from itertools import permutations
class Solution:
    def findSubsequences(self, nums: List[int]) -> List[List[int]]:
        subsequences = set()
        for num in nums:
            new_subsequences = set()
            new_subsequences.add((num,))
            print("new ", new_subsequences)
            print("before ", subsequences)

            for s in subsequences:
                print("s ", s[-1])
                if num >= s[-1]:
                    new_subsequences.add(s+(num,))
            subsequences |= new_subsequences
            print(subsequences)
        return [s for s in subsequences if len(s) > 1]

if __name__ == "__main__":
    nums = [4,6,7,7]
    exp = [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

    nums = [4,4,3,2,1]
    exp = [[4,4]]

    x = Solution()
    res = x.findSubsequences(nums)
    print(exp, res)