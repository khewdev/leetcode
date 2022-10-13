func search(nums []int, target int) int {
    minIndex := 0
    maxIndex := len(nums) - 1
    
    for minIndex <= maxIndex {
        midIndex := (minIndex + maxIndex) / 2
        midItem := nums[midIndex]
        
        if midItem == target {
            return midIndex
        }
        
        if midItem < target {
            minIndex = midIndex + 1
        } else {
            maxIndex = midIndex - 1
        }
    }
    
    return -1
}