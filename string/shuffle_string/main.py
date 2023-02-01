from typing import List
class Solution:
    # using python built-in str function are better = more optimized solution
    def restoreString(self, s: str, indices: List[int]) -> str:
        d = dict(zip(indices,s))
        # print("zip result ", tuple(zip(indices, s)))  -> hasil zip: ((4, 'c'), (5, 'o'), (6, 'd'), (7, 'e'), (0, 'l'), (2, 'e'), (1, 'e'), (3, 't'))
        print("dict ", [d[i] for i in sorted(d)] )

        e = list(zip(indices,s)) #[(4, 'c'), (5, 'o'), (6, 'd'), (7, 'e'), (0, 'l'), (2, 'e'), (1, 'e'), (3, 't')]
        print("list ", e)
        
        return ''.join([d[i] for i in sorted(d)])

    # not optimized 
    def restoreString1(self, s: str, indices: List[int]) -> str:
        res = [0]*len(indices)
        for i in range (0, len(res), 1):
            print("i ", i)
            for j in range (0, len(indices), 1):
                if indices[j] == i:
                    res[i] = s[j]
                    break
        return "".join(res)

if __name__ == "__main__" :
    s = "codeleet"
    indices = [4,5,6,7,0,2,1,3]
    exp = "leetcode"

    # s = "abc"
    # indices = [0,1,2]
    # exp = "abc"

    x = Solution()
    res = x.restoreString(s, indices)
    print(exp, res)