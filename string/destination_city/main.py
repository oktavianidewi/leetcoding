from typing import List
class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        cities = dict()
        for path in paths:
            cities[path[0]] = path[1]
        res = ""
        for path in paths:
            if path[1] not in cities:
                res = path[1]
                break
        return res
        

if __name__ == "__main__":
    paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
    exp = "Sao Paulo"
    paths = [["B","C"],["D","B"],["C","A"]]
    exp = "A"
    paths = [["A","Z"]]
    exp = "Z"
    s = Solution()
    res = s.destCity(paths)
    print(exp, res)