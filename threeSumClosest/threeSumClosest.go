package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var min float64 = math.MaxFloat64
	var res int
	sort.Ints(nums)
	for i := range nums {
		if i == len(nums)-1 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			temp := nums[i+1:]
			for j := range temp {
				var tail int = len(temp) - 1
				if j == tail {
					break
				}
				if j == 0 || temp[j] != temp[j-1] {
					for tail > j {
						if math.Abs(float64(target-(nums[i]+temp[j]+temp[tail]))) < min {
							min = math.Abs(float64(target - (nums[i] + temp[j] + temp[tail])))
							res = nums[i] + temp[j] + temp[tail]
						}
						tail--
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}
