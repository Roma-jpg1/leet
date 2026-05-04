func findTheDifference(s string, t string) byte {
    var r byte
    for i:=0; i<len(s); i++{
        r^=s[i]
        r^=t[i]
        
    }
    r^=t[len(t)-1]
    return r
}