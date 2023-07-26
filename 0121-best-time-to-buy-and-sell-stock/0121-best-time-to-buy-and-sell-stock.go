func max(a int, b int) int {
    if a > b {
        return a
    } 
    return b
}

func maxProfit(prices []int) int {
    l, r := 0, 1
    res := 0
    
    for r < len(prices) {
        if prices[l] < prices[r] {
            res = max(res, prices[r] - prices[l])
        } else {
            l = r
        }
        r++
    }
    
    return res
}