func merge(nums1 []int, m int, nums2 []int, n int)  {
    e:=len(nums1)-1
    n1:=m-1
    n2:=n-1
    for n1>=0 && n2>=0{
        if nums2[n2]>nums1[n1]{
            nums1[e]=nums2[n2]
            e--
            n2--
        } else{
            nums1[e]=nums1[n1]
            e--
            n1--
        }
    }

    for n2>=0{
        nums1[e]=nums2[n2]
        n2--
        e--
    }
    
    
}