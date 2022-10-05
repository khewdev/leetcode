func max(x int, y int) int {
    if x < y {
        return y
    }
    return x
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxArea(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    
    for l < r {
        area := (r - l) * min(height[l] ,height[r])
        res = max(res, area)
        
        if height[l] < height[r] {
            l += 1
        } else {
            r -= 1
        }
    }
    
    return res
}