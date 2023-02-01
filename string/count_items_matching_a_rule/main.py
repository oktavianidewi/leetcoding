from typing import List

class Solution:
    def countMatches(self, items: List[List[str]], ruleKey: str, ruleValue: str) -> int:
        # items = ["type", "color", "name"]
        res = 0
        dict_rule = {"type": 0, "color": 1, "name": 2}
        for item in items:
            if item[dict_rule[ruleKey]] == ruleValue:
                res += 1
        return res


if __name__ == "__main__":
    items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]]
    ruleKey = "color"
    ruleValue = "silver"
    exp = 1

    items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]]
    ruleKey = "type"
    ruleValue = "phone"
    exp = 2

    x = Solution()
    res = x.countMatches(items, ruleKey, ruleValue)
    print(exp, res)