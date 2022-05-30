func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    
    for i, iValue := range (nums) {
        _, jValue := hm[iValue]
        if jValue {
            return []int{i, hm[iValue]}
        }
        
        hm[target - iValue] = i
    }
    
    return nil
}