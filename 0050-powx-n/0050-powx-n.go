import "math"

func helper(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    
    if n == 0 {
        return 1
    }
    
    res := helper(x, n / 2)
    res *= res
    
    if n % 2 == 0 {
        return float64(res)
    } else {
        return float64(x) * float64(res)
    }
}

func myPow(x float64, n int) float64 {
    res := helper(x, int(math.Abs(float64(n))))
    if n >= 0 {
        return res
    } else {
        return 1 / res
    }
}