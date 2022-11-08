func twoSum(nums []int, target int) []int {
    if len(nums) == 0 {
        return nil
    }
    
    hashMap := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        
        _, found := hashMap[complement] 
        if found {
            return []int{i, hashMap[complement]}
        }
        
        hashMap[num] = i
    }
    
    return nil
}