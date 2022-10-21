func sum(arr []int) int {
    ret := 0
    for _, elem := range arr {
        ret += elem
    }
    return ret
}

func sumOddLengthSubarrays(arr []int) int {
    ret := 0
    for currentLen := 1; currentLen <= len(arr); currentLen += 2 {
        for i := range arr {
            if i + currentLen > len(arr) {
                continue
            }
            ret += sum(arr[i:i + currentLen])
        }
    }
    return ret
}