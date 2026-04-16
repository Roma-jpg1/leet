func sortArray(nums []int) []int {
    if len(nums)<=1{
        return nums
    }
    
    QuickSort(nums, 0,len(nums)-1)

    return nums
}

func QuickSort(nums []int, l,r int){

    for l<r{
        med := nums[l+(r-l)/2]

        i, j, k := l, l, r
		for j <= k {
			if nums[j] < med {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			} else if nums[j] > med {
				nums[j], nums[k] = nums[k], nums[j]
				k--
			} else {
				j++
			}
		}

		if i-l < r-k {
			QuickSort(nums, l, i-1)
			l = k + 1
		} else {
			QuickSort(nums, k+1, r)
			r = i - 1
		}

    }
    
}