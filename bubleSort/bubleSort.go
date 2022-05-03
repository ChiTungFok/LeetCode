package main

import "fmt"

func bubleSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(bubleSort([]int{12, 31, 23, 5, 12, 32, 6, 6, 1, 23, 123, 55, 1}))
}
