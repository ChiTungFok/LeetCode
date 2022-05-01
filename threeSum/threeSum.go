package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	m := make(map[int]struct{})
	sort.Ints(nums)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}
		var head, tail, target int
		if i < len(nums)-1 {
			head = i + 1
		} else {
			break
		}
		mp := make(map[int]struct{})
		tail = len(nums) - 1
		target = 0 - num
		for head < tail {
			if _, ok := mp[nums[head]]; ok {
				head++
				continue
			}
			if nums[head]+nums[tail] == target {
				m[num] = struct{}{}
				mp[nums[head]] = struct{}{}
				res = append(res, []int{num, nums[head], nums[tail]})
				head++
				continue
			}
			if nums[head]+nums[tail] < target {
				head++
				continue
			}
			if nums[head]+nums[tail] > target {
				tail--
				continue
			}
		}
	}
	return res
}
