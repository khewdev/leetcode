func binSearch(nums []int, target int, leftBias bool) int {
    l, r := 0, len(nums)-1
    i := -1
    
    for l <= r {
        m := (l+r) / 2
        if target > nums[m] {
            l = m + 1
        } else if target < nums[m] {
            r = m - 1
        } else {
            i = m
            
            if leftBias {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }
    
    return i
}

func searchRange(nums []int, target int) []int {
    left := binSearch(nums, target, true)
    right := binSearch(nums, target, false)
    return []int{left, right}
}