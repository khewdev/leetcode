func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}
func maxSubArray(nums []int) int {
    maxSub := nums[0]
    curSum := 0
    
    for _, num := range nums {
        if curSum < 0 {
            curSum = 0
        }
        curSum += num
        maxSub = max(maxSub, curSum)
    } 
    
    return maxSub
}