class Solution:
    # optimized
    """
    from collections import defaultdict
    class Solution:
        def countPoints(self, rings: str) -> int:
            count=0
            an=collections.defaultdict(list)
            for x in range(1,len(rings),2):
                an[rings[x]].append(rings[x-1])
            for x in an:
                if len(set(an[x]))==3:
                    count+=1
            return count
    """

    # not optimized
    def countPoints(self, rings: str) -> int:
        if len(rings) < 3:
            return 0
        
        rods_rings = dict()
        for i, v in enumerate(rings):
            if i % 2 == 1:
                temp = []
                if rings[i] not in rods_rings:
                    rods_rings[rings[i]] = set(rings[i-1])
                else:
                    rods_rings[rings[i]].add(rings[i-1])

        # check rods_rings
        # print("rods_rings ", rods_rings)
        count = 0
        for i in rods_rings:
            x = list(rods_rings[i])
            x.sort()
            st = ''.join(x)
            if st == "BGR":
                count += 1
        return count




if __name__ == "__main__":
    # rings = "B0B6G0R6R0R6G9"
    # exp = 1
    rings = "B0R0G0R9R0B0G0"
    exp = 1
    # rings = "G4"
    # exp = 0
    # rings = "G7G8G0"
    # exp = 0

    # rings = "G3R3R7B7R5B1G8G4B3G6"
    # exp = 1

    x = Solution()
    res = x.countPoints(rings)
    print(exp, res)
    