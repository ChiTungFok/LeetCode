package main

import "fmt"

func prisonAfterNDays(cells []int, n int) []int {
	var t int
	for i := len(cells) - 1; i >= 0; i-- {
		t = t*2 + cells[i]
	}
	var s []int
	var start int
	for i := 0; i < n; i++ {
		t = (^(t<<1 ^ t>>1)) & 0x7E
		if t == start && len(s) > 0 {
			break
		}
		if i == 0 {
			start = t
		}
		s = append(s, t)
	}
	index := (n - 1) % len(s)
	t = s[index]
	var res []int
	for t > 0 {
		res = append(res, t%2)
		t /= 2
	}
	if len(res) < len(cells) {
		for len(res) < len(cells) {
			res = append(res, 0)
		}
	}
	return res
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
}
