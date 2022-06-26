func containsDuplicate(nums []int) bool {
    hm := make(map[int]int)
    
    for _, iValue := range(nums) {
        _, found := hm[iValue]
        
        if found {
            return true
        }
        
        hm[iValue] = iValue
    }
    
    return false
}