func lengthOfLongestSubstring(s string) int {
    start := 0
    ans := 0
    dict := map[rune]int{}
    for end, char := range(s){
        if _, ok := dict[char];!ok{
            dict[char] = end
        } else{
            if end - start > ans{
                ans = end - start
            }
            idx := dict[char]
            for start <= idx{
                tchar := s[start]
                delete(dict, rune(tchar));
                start += 1
            }
            dict[char] = end
        }
    }
    if len(s) - start > ans{
        ans = len(s) - start
    }
    return ans
}