from typing import List
class Solution:
    def sortPeople(self, names: List[str], heights: List[int]) -> List[str]:
        name_height = dict(zip(heights, names))
        keys = list(name_height.keys())
        keys.sort(reverse=True)
        result = [name_height[i] for i in keys]
        return result

if __name__ == "__main__":
    names = ["Mary","John","Emma"]
    heights = [180,165,170]
    exp = ["Mary","Emma","John"]
    x = Solution()
    res = x.sortPeople(names, heights)
    print("exp, res ", exp, res)