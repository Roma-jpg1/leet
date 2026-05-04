func reverseBits(n int) int {
    var a int
    for i:=0; i<32; i++{
        a |= (n&1)<<(31-i)
        n>>=1
    }
    return a
}