func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    
    for i, iValue := range(nums) {
        complement := target - iValue
        _, jValue := hm[complement]
        
        if jValue {
            return []int{i, hm[complement]}
        }
        
        hm[iValue] = i
    }
    
    return nil
}