func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func trap(height []int) int {
    if height == nil || len(height) < 3 {
        return 0
    }
    
    // left pointer and right pointer
    l, r := 0, len(height) - 1
    
    // initializing left max and right max
    leftMax, rightMax := height[l], height[r]
    
    // result
    res := 0
    
    // while left pointer < right pointer
    for l < r {
        if leftMax < rightMax {
            l++ // increment left pointer by 1 (move right 1 step)
            
            // update leftMax with the max of current leftMax and 
            // the newly shifted left pointer's height
            leftMax = max(leftMax, height[l]) 
            
            // newly updated leftMax - newly shifted left pointer height works because
            // newly shifted left pointer height must be lower than leftMax
            // in order to trap water
            res += leftMax - height[l]
        } else {
            r-- // decrement right pointer by 1 (move left 1 step)
            
            // update rightMax with the max of current rightMax and 
            // the newly shifted right pointer's height
            rightMax = max(rightMax, height[r])

            // newly updated rightMax - newly shifted right pointer height
            res += rightMax - height[r]
        }
    }
    
    return res   
}