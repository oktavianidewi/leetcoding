from itertools import combinations, permutations
from collections import defaultdict

# not optimized
def find_sum_of_three_1(nums, target):
    def add(nums):
        return sum(nums) == target

    # x = list(permutations(nums,3))
    x = list(filter(add, list(combinations(nums,3))))
    if len(x) > 0:
        return True
    return False

# optimized
def find_sum_of_three_(nums, target):
    nums.sort()
    for curr in range(len(nums)-1):
        print(curr)
        left = 1
        right = len(nums)-1
        while left < right :
            print("hasil ", nums[curr]+nums[left]+nums[right], curr, left, right)
            sumOfThree = nums[curr]+nums[left]+nums[right]
            if sumOfThree == target:
                return True
            elif sumOfThree < target:
                left += 1
            elif sumOfThree > target:
                right -= 1
    return False

def find_sum_of_three(nums, target):
    nums.sort()
    for i in range(0, len(nums)-2):
        low = i + 1
        high = len(nums) - 1
        while low < high:
            triple = nums[i] + nums[low] + nums[high]
            if triple == target:
                return True
            elif triple < target:
                low += 1
            else:
                high -= 1
    return False

if __name__ == "__main__":
    nums = [1, -1, 0]
    target = -1
    exp = False

    # nums = [-1, 2, 1, -4, 5, -3]
    # target = 0
    # exp = True

    # nums = [1,2,3,4,5,7,8]
    # target = 20
    # exp = True

    num = [-1, 2, 1, -4, 5, -3] 
    target = 7
    exp = False

    res = find_sum_of_three(nums, target)
    print(f"expexted {exp}, result {res}")