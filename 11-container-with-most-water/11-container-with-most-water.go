func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func maxArea(height []int) int {
    size := len(height)
    
    var maxArea int = 0
    var aPtr int
    var bPtr int = size - 1
    
    for aPtr < bPtr {
        if(height[aPtr] < height[bPtr]) {
            maxArea = Max(maxArea, height[aPtr] * (bPtr - aPtr))
            aPtr++
        } else {
            maxArea = Max(maxArea, height[bPtr] * (bPtr - aPtr))
            bPtr--
        }
    }
    
    return maxArea
}