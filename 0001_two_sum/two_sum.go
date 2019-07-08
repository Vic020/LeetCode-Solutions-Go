package main

func twoSum(nums []int, target int) []int {
	b := map[int]int{}
	for i, v := range nums {
		r, e := b[target-v]
		if e {
			return []int{r, i}
		}
		b[v] = i
	}
	return []int{}
}
