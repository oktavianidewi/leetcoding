func letterCasePermutation(s string) []string {
    result := []string{}
    runes := []rune(strings.ToLower(s))
    result = dfs(runes, []string{}, 0)
    return result
}


func dfs(runes []rune, result []string, i int)[]string{
    if i == len(runes) {
        sCopy := string(runes)
        result = append(result, sCopy)
        return result
    }
    
    if i < len(runes)  {
        result = dfs(runes, result, i+1)
        if runes[i] >= 'a' && runes[i] <= 'z' {
            runes[i] = runes[i] - 32
            result = dfs(runes, result, i+1)
            runes[i] = runes[i] + 32
        }
    }
    return result
}