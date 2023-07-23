func twoSum(nums []int, target int) []int {
    if len(nums) == 0 {
        return nil
    }
    
    l, r := 0, len(nums) - 1
    
    for l < r {
        if target > 0 && nums[r] > target {
            r-=1
            continue
        }

        curSum := nums[l] + nums[r]
        
        if curSum < target {
            l+=1
        } else if curSum > target {
            r-=1
        } else {
            return []int{l + 1, r + 1}
        }
        
        
    }
            
    return nil
}
