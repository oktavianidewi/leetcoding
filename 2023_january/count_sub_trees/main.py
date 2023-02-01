from typing import List, Counter
from collections import defaultdict

class Solution:
    def countSubTrees(self, n: int, edges: List[List[int]], labels: str) -> List[int]:
        tree = defaultdict(list)
        for s,e in edges:
            tree[s].append(e)
            tree[e].append(s)
        
        res = [0] * n
        
        def dfs(node, par):
            count = Counter()
            for nei in tree[node]:
                if nei != par:
                    count += dfs(nei, node)
            count[labels[node]] += 1
            res[node] = count[labels[node]]
            return count
        
        dfs(0,-1)
        return res
    


if __name__ == "__main__":
    n = 7
    edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]
    labels = "abaedcd"
    s = Solution()
    res = s.countSubTrees(n, edges, labels)
    print("hasil ", res)