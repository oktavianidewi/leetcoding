from typing import List
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        result = []
        for i in range(len(intervals)):
            if newInterval[1] < intervals[i][0]:
                result.append(newInterval)
                return result + intervals[i:]
            elif newInterval[0] > intervals[i][1]:
                result.append(intervals[i])
            else:
                newInterval = [min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])]
                # print(newInterval)
        result.append(newInterval)
        # print(result)
        return result

if __name__ == "__main__":
    intervals = [[1,3],[6,9]]
    newInterval = [2,5]
    exp = [[1,5],[6,9]]

    intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]
    newInterval = [4,8]
    exp = [[1,2],[3,10],[12,16]]

    x = Solution()
    res = x.insert(intervals, newInterval)
    print(exp, res)