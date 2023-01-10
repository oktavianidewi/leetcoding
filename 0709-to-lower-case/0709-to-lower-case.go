func toLowerCase(str string) string {
    var result strings.Builder
    
    for _, char := range str{
        if char >= 'A' && char <= 'Z'{
            result.WriteRune(char - 'A' + 'a')
        }else{
            result.WriteRune(char)
        }
    }
    return result.String()
}