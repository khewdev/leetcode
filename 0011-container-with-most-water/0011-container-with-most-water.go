import "math"

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    res := 0
    
    for l < r {
        area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
        
        res = int(math.Max(float64(res), float64(area)))
        
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    
    return res
}