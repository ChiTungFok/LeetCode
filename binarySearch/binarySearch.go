package binarysearch

func search(nums []int, target int) int {
	var head, tail int
	tail = len(nums) - 1
	for head < tail {
		mid := nums[(head+tail)/2]
		if target == mid {
			return (head + tail) / 2
		}
		if target < mid {
			tail = (head + tail) / 2
			continue
		}
		if head == (head+tail)/2 {
			head++
			continue
		}
		head = (head + tail) / 2
	}
	if head == tail && nums[head] == target {
		return head
	}
	return -1
}
