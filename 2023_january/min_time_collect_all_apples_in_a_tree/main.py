from collections import defaultdict
from typing import List

class Solution:
    def minTime(self, n: int, edges: List[List[int]], hasApple: List[bool]) -> int:
        maps = defaultdict(list)
        for x, y in edges:
            maps[x].append(y)
            maps[y].append(x)


        def traverse(node, parent):
            res = 0
            for child in maps[node]:
                if child != parent:
                    res += traverse(child, node)
            if res or hasApple[node]:
                return res + 2
            return res
        
        return max(traverse(0, -1)-2, 0)

if __name__ == "__main__":
    # n = 7
    # edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]
    # hasApple = [False,False,True,False,True,True,False]

    
    n = 7
    edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]
    hasApple = [False,False,True,False,False,True,False] 
    # hasApple = [False,False,False,False,False,False,False]
    s = Solution()
    res = s.minTime(n, edges, hasApple)
    print(res)