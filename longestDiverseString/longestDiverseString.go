package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestDiverseString(a int, b int, c int) string {
	sl := []*struct {
		cnt int
		b   byte
	}{
		{
			cnt: a,
			b:   'a',
		},
		{
			cnt: b,
			b:   'b',
		},
		{
			cnt: c,
			b:   'c',
		},
	}
	var res []byte
	for {
		sort.Slice(sl, func(i, j int) bool {
			return sl[i].cnt > sl[j].cnt
		})
		hasNext := false
		for _, item := range sl {
			if item.cnt <= 0 {
				continue
			}
			if len(res) < 2 {
				res = append(res, item.b)
				item.cnt--
				hasNext = true
				break
			}
			if res[len(res)-1] == res[len(res)-2] && res[len(res)-1] == item.b {
				continue
			}
			res = append(res, item.b)
			item.cnt--
			hasNext = true
			break
		}
		if !hasNext {
			break
		}
	}
	return string(res)
}

func findLongestDiverseString(a int, b int, c int, s string) string {
	if strings.Contains(s, "aaa") {
		return ""
	}
	if strings.Contains(s, "bbb") {
		return ""
	}
	if strings.Contains(s, "ccc") {
		return ""
	}
	if a <= 0 && b <= 0 && c <= 0 {
		return s
	}
	var as, bs, cs string
	if a > 0 {
		as = findLongestDiverseString(a-1, b, c, s+"a")
	}
	if b > 0 {
		bs = findLongestDiverseString(a, b-1, c, s+"b")
	}
	if c > 0 {
		cs = findLongestDiverseString(a, b, c-1, s+"c")
	}
	return max(max(as, s), max(bs, cs))
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestDiverseString(1, 1, 7))
}
