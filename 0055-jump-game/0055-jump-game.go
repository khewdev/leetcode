func canJump(nums []int) bool {
    goal := len(nums) - 1
    
    for i := len(nums) - 1; i >= 0; i-- {
        if i + nums[i] >= goal {
            goal = i
        }
    }
    
    if goal == 0 {
        return true
    } else {
        return false
    }
    
}