func Max(x, y int)int{
    if x >= y {
        return x
    }
    
    return y
}

func maxArea(height []int) int {
    
    // length of input array
    size := len(height)
    
    // tow pointers, left init as 0, right init as size-1
    left, right := 0, size-1
    
    // maximal width between leftmost stick and rightmost stick
    maxWidth := size-1
    
    // area also known as the amount of water
    area := 0
    
    // trade-off between width and height
    // scan each possible width and compute maximal area
    for width := maxWidth ; width > 0 ; width-- {
        
        if height[left] < height[right]{
            
            // the height of lefthand side is shorter
            area = Max(area, width * height[left])
            
            // update left index to righthand side
            left += 1
            
        }else{
            
            // the height of righthand side is shorter
            area = Max(area, width * height[right])
            
            // update left index to righthand side
            right -= 1
        }
        
    }
    return area
}