func spiralOrder(matrix [][]int) []int {
    res := []int{}
    
    left, right := 0, len(matrix[0])
    top, bottom := 0, len(matrix)
    
    for left < right && top < bottom {
        // get every i in the top row
        for i := left; i < right; i++ {
            res = append(res, matrix[top][i])
        }
        top += 1
        
        // get every i in the right col
        for i := top; i < bottom; i++ {
            res = append(res, matrix[i][right-1])
        }
        right -= 1
        
        if !(left < right && top < bottom) {
            break
        }
        
        // get every i in the bottom row
        for i := right - 1; i > left - 1; i-- {
            res = append(res, matrix[bottom-1][i])
        }
        bottom -= 1
        
        // get every i in the left col
        for i := bottom -1; i > top - 1; i-- {
            res = append(res, matrix[i][left])
        }
        left += 1
    }
    
    return res
}