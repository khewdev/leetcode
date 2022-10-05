import "math"

func maxArea(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    
    for l < r {
        area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
        res = int(math.Max(float64(res), float64(area)))
        
        if height[l] < height[r] {
            l += 1
        } else {
            r -= 1
        }
    }
    
    return res
}