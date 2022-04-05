package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := range nums {
		if i == len(nums)-1 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			temp := nums[i+1:]
			var tail int = len(temp) - 1
			for j := range temp {
				if j == tail {
					break
				}
				if j == 0 || temp[j] != temp[j-1] {
					for nums[i]+temp[j]+temp[tail] > 0 && tail > j {
						tail--
					}
					if nums[i]+temp[j]+temp[tail] == 0 {
						res = append(res, []int{nums[i], temp[j], temp[tail]})
					}
				}
			}
		}
	}
	return res
}
