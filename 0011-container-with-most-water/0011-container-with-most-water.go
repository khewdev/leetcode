func maxArea(height []int) int {
    start := 0
    end := len(height) - 1
    maxArea := 0
    for start < end {
        currArea := end - start
        if height[start] < height[end] {
            currArea *= height[start]
            start++
        } else {
            currArea *= height[end]
            end--
        }
        if maxArea < currArea {
            maxArea = currArea
        }
    }
    
    return maxArea
}