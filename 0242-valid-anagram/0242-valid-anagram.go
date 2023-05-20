func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    } 
    
    hash := map[rune]int{}
    
    for _, value := range s {
        hash[value]++
    }
    
    for _, value := range t {
        hash[value]--
    }
    
    for _, value := range hash {
        if value != 0 {
            return false
        }
    }
    
    return true
}