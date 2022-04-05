package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var n int
	var (
		meet5    bool
		meet10   bool
		meet50   bool
		meet100  bool
		meet500  bool
		meet1000 bool
	)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'I' && (meet5 || meet10) {
			n = n - 1
			continue
		}
		if s[i] == 'X' && (meet50 || meet100) {
			n = n - 10
			continue
		}
		if s[i] == 'C' && (meet500 || meet1000) {
			n = n - 100
			continue
		}
		if s[i] == 'V' {
			meet5 = true
		}
		if s[i] == 'X' {
			meet10 = true
		}
		if s[i] == 'L' {
			meet50 = true
		}
		if s[i] == 'C' {
			meet100 = true
		}
		if s[i] == 'D' {
			meet500 = true
		}
		if s[i] == 'M' {
			meet1000 = true
		}
		n += m[s[i]]
	}
	return n
}

func main() {
	fmt.Println(romanToInt("IV"))
}
