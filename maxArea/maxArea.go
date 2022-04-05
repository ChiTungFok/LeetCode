package main

import "fmt"

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	var m int
	for start < end {
		m = max(m, min(height[start], height[end])*(end-start))
		if height[start] > height[end] {
			end--
			continue
		} else {
			start++
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 1}))
}
