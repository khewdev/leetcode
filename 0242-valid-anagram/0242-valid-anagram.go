func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    } 
    
    alphabet := make([]int, 26)
    sBytes := []byte(s)
    tBytes := []byte(t)
    
    for i := 0; i < len(sBytes); i++ {
        alphabet[sBytes[i] - 'a']++
    }
    
    for i := 0; i < len(tBytes); i++ {
        alphabet[tBytes[i] - 'a']--
    }
    
    for i := 0; i < len(alphabet); i++ {
        if alphabet[i] != 0 {
            return false
        }
    }
    
    return true
}