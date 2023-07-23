func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    freq := make([][]int, len(nums)+1)
    res := []int{}

    for _, v := range nums {
        count[v]++
    }

    for i, v := range count {
        freq[v] = append(freq[v], i) 
    }

    for i:=len(freq)-1; i>0; i-- {
        for _, v := range freq[i] {
            res = append(res, v)
        }

        if len(res) == k {
            return res
        }
    }

    return res
}