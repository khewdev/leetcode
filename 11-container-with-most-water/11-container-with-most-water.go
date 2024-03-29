import "math"

func maxArea(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    
    for l < r {
        area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
        res = int(math.Max(float64(area), float64(res)))
        
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    
    return res
}