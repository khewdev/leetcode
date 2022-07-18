func productExceptSelf(nums []int) []int {
    size := len(nums)
    res := make([]int, size)
    
    left := 1
    for i := 0; i < size; i++ {
        // if i > 0 {
        //     left *= nums[i - 1]
        // }
        res[i] = left
        left *= nums[i]
    }
    
    right := 1
    for i := size - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]

        // if i < size - 1 {
        //     right *= nums[i + 1]
        // }
        
    }
    
    return res
}