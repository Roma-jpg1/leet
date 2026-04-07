func twoSum(nums []int, target int) []int {
    n := make(map[int]int)

    for i,num := range nums{
        need:=target-num

        if idx, ok :=n[need]; ok{
            return []int{idx, i}
        }
        n[num]=i

    }
    return nil
    
}