from typing import List
class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        value = 0
        for op in operations:
            if op == "++X" or op == "X++":
                value += 1
            elif op == "--X" or op == "X--":
                value -= 1
        return value

if __name__ == "__main__":
    operations, exp = ["--X","X++","X++"], 1
    operations, exp = ["++X","++X","X++"], 3
    operations, exp = ["X++","++X","--X","X--"], 0
    x = Solution()
    res = x.finalValueAfterOperations(operations)
    print(res, exp)