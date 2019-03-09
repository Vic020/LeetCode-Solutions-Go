package main

import (
	"fmt"
	"math"
	"sort"
)

// method 1
// 两个数组合成一个数组，然后排序，直接算中间两个位置的数
// 28 ms
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[int(l/2-1)]+nums[int(l/2)]) / 2

	} else {
		return float64(nums[int(math.Floor(float64(l)/2))])
	}

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{2, 4}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 3}))
}
