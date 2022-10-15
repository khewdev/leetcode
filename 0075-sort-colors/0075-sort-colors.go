func sortColors(nums []int)  {
    swapCount := 1
    
    for swapCount > 0 {
        swapCount = 0
        for itemIndex := 0; itemIndex < len(nums)-1; itemIndex++ {
            if nums[itemIndex] > nums[itemIndex+1] {
                nums[itemIndex], nums[itemIndex+1] = nums[itemIndex+1], nums[itemIndex]
                swapCount += 1
            }
        }
    }
}