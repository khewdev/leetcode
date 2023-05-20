func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

func replaceElements(arr []int) []int {
    j, temp := -1, 0
    
    for i := len(arr) - 1; i >= 0; i-- {
        temp = arr[i]
        arr[i] = j
        j = max(j, temp)
    }
    
    return arr
}