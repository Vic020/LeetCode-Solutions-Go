package main

import (
	"fmt"
	"math"
	"sort"
)

// method 1
// 两个数组合成一个数组，然后排序，直接算中间两个位置的数
// 28 ms O(nlogn)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums) // O(nlogn)
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[int(l/2-1)]+nums[int(l/2)]) / 2

	} else {
		return float64(nums[int(math.Floor(float64(l)/2))])
	}

}


// 20ms 5.4MB O((m+n)/2)
// 在已知长度和已排序的情况下，我们可以直接计算出需要移动指针次数，
// 直接在单端移动指针需要的次数即可
// --------------原思路
// 这道题的难度在于同时两个数组同时判断， 指针同时向中间数逼近
// 结束的标志是
// while *p1 <= *p2:
//    moveLeft()
//    moveRight()
// we know that nums1 & nums2 is not empty from problem content,
// So we can obviously confirm that the number of steps which are pushed from each side to the middle
// is related with len(nums1) + len(nums2)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var w1, w2 int
	// init
	p1, p2 := 0, 0
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 {
		if l2%2 == 0 {
			return float64(nums2[l2/2-1]+nums2[l2/2]) / 2
		} else {
			return float64(nums2[int(math.Floor(float64(l2)/2))])
		}
	}
	if l2 == 0 {
		if l1%2 == 0 {
			return float64(nums1[l1/2-1]+nums1[l1/2]) / 2
		} else {
			return float64(nums1[int(math.Floor(float64(l1)/2))])
		}
	}

	if nums1[p1] < nums2[p2] {
		w1 = nums1[p1]
	} else {
		w1 = nums2[p2]
	}

	steps := l1 + l2

	if steps%2 == 0 {

		for steps = steps / 2; steps != 0; steps-- {

			if nums1[p1] < nums2[p2] {
				if p1+1 < l1 {
					if nums1[p1+1] < nums2[p2] {
						p1++
						w2 = w1
						w1 = nums1[p1]
					} else {
						p1++
						w2 = w1
						w1 = nums2[p2]
					}
				} else {
					w2 = w1
					w1 = nums2[p2]
					if p2+1 < l2 {
						p2++
					}
				}
			} else {
				if p2+1 < l2 {
					if nums1[p1] < nums2[p2+1] {

						p2++
						w2 = w1
						w1 = nums1[p1]

					} else {
						if p2+1 < l2 {
							p2++
						}
						w2 = w1
						w1 = nums2[p2]

					}
				} else {
					w2 = w1
					w1 = nums1[p1]
					if p1+1 < l1 {
						p1++
					}
				}
			}
		}
		return float64(w1+w2) / 2
	} else {

		for steps = int(math.Floor(float64(steps) / 2)); steps != 0; steps-- {
			if nums1[p1] < nums2[p2] {
				if p1+1 < l1 {
					if nums1[p1+1] < nums2[p2] {
						p1++
						w1 = nums1[p1]
					} else {
						p1++
						w1 = nums2[p2]
					}
				} else {
					w1 = nums2[p2]
					if p2+1 < l2 {
						p2++
					}
				}
			} else {
				if p2+1 < l2 {
					if nums1[p1] < nums2[p2+1] {

						p2++
						w1 = nums1[p1]

					} else {
						if p2+1 < l2 {
							p2++
						}
						w1 = nums2[p2]

					}
				} else {
					w1 = nums1[p1]
					if p1+1 < l1 {
						p1++
					}
				}
			}
		}
		return float64(w1)
	}

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{100000}, []int{100001}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{2, 4}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{5}, []int{1, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 5, 11, 13}, []int{3, 7, 9}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{5, 6, 8}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4}))
}
