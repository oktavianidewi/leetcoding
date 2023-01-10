func combine(n int, k int) [][]int {
    result := [][]int{}
    backtrack(n, k, 1, []int{}, &result)
    return result
}

func backtrack(n, k, start int, comb []int, result *[][]int) {
    // base case
    if len(comb) == k {
        tmp := make([]int, k)
        copy(tmp, comb)
        *result = append(*result, tmp)
        return
    }
    
    for i := start; i <= n; i++ {
        comb = append(comb, i)
        backtrack(n, k, i+1, comb, result)
        comb = comb[:len(comb)-1]
    }
}