import "math"

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
            leftMax = int(math.Max(float64(leftMax), float64(height[l])))
            res += leftMax - height[l]
        } else {
            r--
            rightMax = int(math.Max(float64(rightMax), float64(height[r])))
            res += rightMax - height[r]
        }
    }
    
    return res   
}