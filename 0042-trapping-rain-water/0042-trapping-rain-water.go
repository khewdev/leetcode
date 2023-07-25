func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func trap(height []int) int {
    if height == nil {
        return 0
    }
    
    l, r := 0, len(height) - 1
    leftMax, rightMax := height[l], height[r]
    res := 0
    
    for l < r {
        if leftMax < rightMax {
            l++
            leftMax = max(leftMax, height[l])
            res += leftMax - height[l]
        } else {
            r--
            rightMax = max(rightMax, height[r])
            res += rightMax - height[r]
        }
    }
    
    return res   
}