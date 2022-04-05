package main

import "strconv"

func decodeString(s string) string {
	stack := Stack{}
	for _, r := range s {
		if string(r) != "]" {
			stack.Push(string(r))
			continue
		}
		sta := Stack{}
		for stack.Cap() > 0 && stack.Top() != "[" {
			sta.Push(stack.Pop())
		}
		stack.Pop()
		nstack := Stack{}
		var num string
		for stack.Top() >= "0" && stack.Top() <= "9" {
			nstack.Push(stack.Pop())
		}
		for nstack.Cap() > 0 {
			num += nstack.Pop()
		}
		n, _ := strconv.ParseInt(num, 10, 32)
		ss := Clone(&sta, int(n))
		for _, sss := range ss {
			for sss.Cap() > 0 {
				stack.Push(sss.Pop())
			}
		}
	}
	var res string
	for _, s := range stack.s {
		res += s
	}
	return res
}

type Stack struct {
	s []string
}

func (s *Stack) Cap() int {
	return len(s.s)
}

func (s *Stack) Push(ss string) {
	s.s = append(s.s, ss)
}

func (s *Stack) Pop() string {
	if len(s.s) == 0 {
		return ""
	}
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

func (s *Stack) Top() string {
	if len(s.s) == 0 {
		return ""
	}
	return s.s[len(s.s)-1]
}

func Clone(s *Stack, n int) []*Stack {
	if s == nil {
		return nil
	}
	var res []*Stack
	for i := 0; i < n; i++ {
		ss := &Stack{}
		ss.s = append(ss.s, s.s...)
		res = append(res, ss)
	}
	return res
}
