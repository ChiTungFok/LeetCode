package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var m []map[rune]int
	for index, c := range s {
		cm := make(map[rune]int)
		cm[c] = index
		m = append(m, cm)
	}
	var max int
	for index, c := range s {
		if index == 0 {
			max = Max(len(m[index]), max)
			continue
		}
		if _, ok := m[index-1][c]; ok {
			mergeExist(m[index-1], m[index], c)
		} else {
			mergeMap(m[index-1], m[index])
		}
		max = Max(len(m[index]), max)
	}
	for index, cm := range m {
		fmt.Println(string(s[index]), " ", len(cm))
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeMap(a, b map[rune]int) map[rune]int {
	for k, v := range a {
		b[k] = v
	}
	return b
}

func mergeExist(a, b map[rune]int, c rune) map[rune]int {
	index := a[c]
	for k, v := range a {
		if v <= index {
			continue
		}
		b[k] = v
	}
	return b
}

func main() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))
}
