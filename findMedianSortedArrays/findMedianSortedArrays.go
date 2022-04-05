package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		index1 int
		index2 int
		res    []int
	)
	for {
		if index1 == len(nums1) && index2 == len(nums2) {
			break
		}
		if index1 == len(nums1) {
			res = append(res, nums2[index2:]...)
			break
		}
		if index2 == len(nums2) {
			res = append(res, nums1[index1:]...)
			break
		}
		if nums1[index1] < nums2[index2] {
			res = append(res, nums1[index1])
			index1++
			continue
		}
		res = append(res, nums2[index2])
		index2++
	}
	if (len(nums1)+len(nums2))%2 > 0 {
		mid := (len(nums1) + len(nums2)) / 2
		return float64(res[mid])
	}
	mid := (len(nums1) + len(nums2)) / 2
	return (float64(res[mid-1]) + float64(res[mid])) / 2
}
