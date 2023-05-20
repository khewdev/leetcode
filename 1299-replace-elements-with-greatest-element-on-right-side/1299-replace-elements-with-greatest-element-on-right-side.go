func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

func replaceElements(arr []int) []int {
    j, temp := -1, 0
    
    for i := len(arr) - 1; i >= 0; i-- {
        // set current index value to temp
        temp = arr[i]
        
        // in the first iteration we set last element to -1
        // in following iterations we will set j to current index, which we got the largest value of current index all the way to right side from previous iterations
        arr[i] = j
        
        // set largest value of current index all the way to right side to j
        j = max(j, temp)
    }
    
    return arr
}