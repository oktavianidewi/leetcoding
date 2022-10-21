func min(a, b int) int {
    if a > b { return b }
    return a
}

func isAlienSorted(words []string, order string) bool {
    dic := make([]int, 26)
    for i:= 0; i < len(order); i++ {
        dic[order[i]-'a'] = i
    }
    
    for i:= 0; i < len(words); i++ {
        for j:=i+1; j < len(words); j++ {
            minLen := min(len(words[i]), len(words[j]))
            for k:=0; k < minLen; k++ {
                c1, c2 := words[i][k], words[j][k]
                if dic[c1 - 'a'] < dic[c2 - 'a'] {
                    break
                } else if dic[c1 - 'a'] > dic[c2 - 'a'] {
                    return false
                } else if k == minLen-1 && len(words[i]) > len(words[j]) {
                    return false
                }
            }
        }
    }
    return true
}