func hammingWeight(n int) int {
    a:=strconv.FormatInt(int64(n), 2) 
    return strings.Count(a, "1")
    
}