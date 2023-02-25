class Solution:
    def squareIsWhite(self, coordinates: str) -> bool:
        x = ord(coordinates[0])-ord('a')
        y = int(coordinates[1])
        # square is white = true; square is black = false
        # if y is odds, black is on the even x
        # if y is even, black is on the odd x
        if (y % 2 == 1 and x % 2 == 0) or (y % 2 == 0 and x % 2 == 1):
            return False
        return True            

if __name__ == "__main__":
    coordinates = "a1"
    exp = False

    # coordinates = "h3"
    # exp = True

    # coordinates = "c7"
    # exp = False

    s = Solution()
    res = s.squareIsWhite(coordinates)
    print(exp, res)
