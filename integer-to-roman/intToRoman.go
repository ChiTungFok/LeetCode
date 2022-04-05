package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	if num < 0 || num > 3999 {
		return ""
	}
	var dig int
	var s string
	for num != 0 {
		src := num % 10
		var dst int
		if dig > 0 {
			dst = src * 10 * dig
		} else {
			dst = src
		}
		if n, ok := m[dst]; ok {
			s = n + s
		} else {
			if src > 5 {
				if dig > 0 {
					s = m[50*dig] + strings.Repeat(m[10*dig], src-5) + s
				} else {
					s = m[5] + strings.Repeat(m[1], src-5) + s
				}
			} else {
				if dig > 0 {
					s = strings.Repeat(m[10*dig], src) + s
				} else {
					s = strings.Repeat(m[1], src) + s
				}
			}
		}
		if dig == 0 {
			dig = 1
		} else {
			dig *= 10
		}
		num = num / 10
	}
	return s
}

func main() {
	fmt.Println(intToRoman(27))
}
