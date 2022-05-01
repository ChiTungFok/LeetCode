package main

import (
	"fmt"
	"sort"
)

func arraySign(nums []int) int {
	sort.Ints(nums)
	i := sort.SearchInts(nums, 0)
	if i >= len(nums) {
		if i%2 == 1 {
			return -1
		}
		return 1
	}
	if nums[i] == 0 {
		return 0
	}
	if i%2 == 1 {
		return -1
	}
	return 1
}

func main() {
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
