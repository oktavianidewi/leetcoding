from typing import List
class Solution:
    def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
        event1start = self.convertStrDate(event1[0])
        event1end = self.convertStrDate(event1[1])
        event2start = self.convertStrDate(event2[0])
        event2end = self.convertStrDate(event2[1])
        
        if event2start > event1end or event1start > event2end: 
            return False
        else:
            return True

    # optimal solution
    # class Solution:
    # def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
    #     return max(event1[0], event2[0]) <= min(event1[1], event2[1])

    def convertStrDate(self, date: str) -> int:
        hour, min = date.split(":")
        return (int(hour)+(int(min)/60))


if __name__ == "__main__":
    event1 = ["01:15","02:00"]
    event2 = ["02:00","03:00"]
    exp = True

    # event1 = ["01:00","02:00"]
    # event2 = ["01:20","03:00"]
    # exp = True

    # event1 = ["10:00","11:00"]
    # event2 = ["14:00","15:00"]
    # exp = False

    # event1 = ["14:13","22:08"]
    # event2 = ["02:40","08:08"]
    # exp = False

    x = Solution()
    res = x.haveConflict(event1, event2)
    print(exp, res)