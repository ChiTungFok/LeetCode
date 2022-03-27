package main

func longestPalindrome(s string) string {
	var max string
	for i := range s {
		odd := f(i, i, s)
		even := f(i, i+1, s)
		if len(odd) > len(even) && len(odd) > len(max) {
			max = odd
		}
		if len(even) > len(odd) && len(even) > len(max) {
			max = even
		}
	}
	return max
}

func f(left, right int, s string) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
