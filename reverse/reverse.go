package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var y int
	var stack []int
	var flag bool
	if x < 0 {
		x = -x
		flag = true
	}
	for x != 0 {
		stack = append(stack, x%10)
		x /= 10
	}
	for _, n := range stack {
		y = y*10 + n
	}
	if flag {
		y = -y
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

func main() {
	fmt.Println(reverse(1534236469))
}
