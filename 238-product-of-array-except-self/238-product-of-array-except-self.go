func productExceptSelf(nums []int) []int {
    size := len(nums)
    res := make([]int, size)
    
    prefix := 1
    for i := 0; i < size; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }
    
    postfix := 1
    for i := size-1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }
    
    return res
}