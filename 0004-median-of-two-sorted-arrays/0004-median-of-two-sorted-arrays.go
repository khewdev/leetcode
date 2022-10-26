import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergedArr := []int{}
    mergedArr = append(nums1, nums2...)
    sort.Ints(mergedArr)
    
    mid := len(mergedArr) / 2
    
    if len(mergedArr) % 2 == 0 {
        return (float64(mergedArr[mid - 1]) + float64(mergedArr[mid])) / float64(2)
    } else {
        return float64(mergedArr[mid])
    }
}