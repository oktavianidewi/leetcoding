func matrixReshape(mat [][]int, r int, c int) [][]int {
    if r*c != len(mat)*len(mat[0]) {
        return mat
    }
    return toMatrix(flatten(mat), r, c)
}

func flatten(mat [][]int) []int {
    sol := make([]int, 0, len(mat)*len(mat[0]))
    for _, row := range mat {
        for _, val := range row {
            sol = append(sol, val)
        }
    }
    return sol
}

func toMatrix(flat []int, r int, c int) [][]int {
    curr, sol := 0, make([][]int, r)
    for i := 0; i < r; i ++ {
        row := make([]int, c)
        for j := 0; j < c; j++ {
            row[j], curr = flat[curr], curr+1
        }
        sol[i] = row
    }
    return sol
}