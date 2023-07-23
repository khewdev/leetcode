func twoSum(nums []int, target int) []int {
    if len(nums) == 0 {
        return nil
    }
    
    hashMap := make(map[int]int)
    
    for i, v := range nums {
        complement := target - v
        
        _, found := hashMap[complement]
        
        if found {
            return []int{hashMap[complement], i+1}
        }
        
        hashMap[v] = i+1
    }
    
    return nil
}
