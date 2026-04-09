func largestAltitude(gain []int) int {
    alt:=0
    cur:=0

    for i:=0;i<len(gain);i++{
        cur+=gain[i]
        if cur>alt{
            alt=cur
        }
    }
    return alt
}