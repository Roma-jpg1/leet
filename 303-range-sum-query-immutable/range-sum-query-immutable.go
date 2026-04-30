type NumArray struct {
    numsarr []int     
}


func Constructor(nums []int) NumArray {
    return NumArray {
        numsarr: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    r:=0

    for i:=left; i<=right;i++{
        r+=this.numsarr[i]
    }    
    return r
    
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */