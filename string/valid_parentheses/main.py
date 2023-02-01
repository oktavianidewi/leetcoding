# https://www.geeksforgeeks.org/stack-in-python/

class Solution:
    # error di test case yang memperhatikan order
    def isValid(self, s: str) -> bool:
        a_counter = 0
        b_counter = 0
        c_counter = 0
        for i in s:
            if i == "(":
                a_counter += 1
            elif i == ")":
                a_counter -= 1
            
            if i == "[":
                b_counter += 1
            elif i == "]":
                b_counter -= 1

            if i == "{":
                c_counter += 1
            elif i == "}":
                c_counter -= 1

        return a_counter == 0 and b_counter == 0 and c_counter == 0

if __name__ == "__main__":
    s, exp = "()", True
    # s, exp = "()[]{}", True
    # s, exp = "(}", False
    x = Solution()
    res = x.isValid(s)
    print(exp, res)