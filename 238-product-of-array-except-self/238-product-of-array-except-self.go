func productExceptSelf(nums []int) []int {
    size := len(nums)
    res := make([]int, size)
    left := 1
    
    for i := 0; i < size; i++ {
        if i > 0 {
            left *= nums[i - 1]
        }
        res[i] = left
    }
    
    right := 1
    for i := size - 1; i >= 0; i-- {
        if i < size - 1 {
            right *= nums[i + 1]
        }
        
        res[i] *= right
    }
    
    return res
}