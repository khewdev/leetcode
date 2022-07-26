func Max(x, y int) int {
    if x < y {
        return y
    }
    
    return x
}

func maxProfit(prices []int) int {
    leftPtr := 0
    rightPtr := 1
    maxProfit := 0
    
    for rightPtr < len(prices) {
        if prices[leftPtr] < prices[rightPtr] {
            profit := prices[rightPtr] - prices[leftPtr]
            maxProfit = Max(profit, maxProfit)
        } else {
            leftPtr = rightPtr
        }
        rightPtr += 1
    }
    
    return maxProfit
}
