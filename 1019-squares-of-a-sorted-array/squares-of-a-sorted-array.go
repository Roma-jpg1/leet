func sortedSquares(nums []int) []int {
    res:= make([]int, len(nums))
    i:=0
    j:=len(nums)-1
    k:=len(nums)-1

    for i<=j{
        if nums[i]*nums[i]>nums[j]*nums[j]{
            res[k]=nums[i]*nums[i]
            i++
            k--
        }else {
            res[k]=nums[j]*nums[j]
            j--
            k--
        }
    }
    return res
    
}