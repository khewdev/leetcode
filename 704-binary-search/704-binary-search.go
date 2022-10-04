import "math"

func search(nums []int, target int) int {
    left := 0
    right := len(nums)
    
    for left + 1 < right {
        mid := int(math.Ceil(float64((right + left) / 2)))
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid
        } else {
            right = mid
        }
    }
    
    if nums[left] == target {
        return left
    }
    
    return -1
}