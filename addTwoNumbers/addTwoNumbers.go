package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addSecondWay(l1 *ListNode, l2 *ListNode, res []int, c int) ([]int, int) {
	if l1 == nil && l2 == nil {
		return res, c
	}
	var (
		l1n    int
		l2n    int
		r      int
		l1next *ListNode
		l2next *ListNode
	)
	if l1 != nil {
		l1n = l1.Val
		l1next = l1.Next
	}
	if l2 != nil {
		l2n = l2.Val
		l2next = l2.Next
	}
	r = l1n + l2n + c
	res = append(res, r%10)
	c = r / 10
	return addSecondWay(l1next, l2next, res, c)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c int
	var res []int
	res, c = addSecondWay(l1, l2, res, c)
	if c != 0 {
		res = append(res, c)
	}
	return newList(res)
}

func add(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		return nil
	}
	if getLength(a) != getLength(b) {
		return nil
	}
	curA := a
	curB := b
	var res []int
	var c int
	for {
		if curA == nil || curB == nil {
			break
		}
		r := curA.Val + curB.Val + c
		res = append(res, r%10)
		c = r / 10
		curA = curA.Next
		curB = curB.Next
	}
	if c > 0 {
		res = append(res, c)
	}
	return newList(res)
}

func addSuffix(smaller *ListNode, long int) *ListNode {
	if smaller == nil {
		return nil
	}
	var (
		pre    *ListNode
		cur    = smaller
		length = 0
	)
	for {
		if cur == nil {
			break
		}
		length++
		pre = cur
		cur = cur.Next
	}
	if pre == nil {
		return smaller
	}
	cur = pre
	for {
		if length == long {
			break
		}
		length++
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	return smaller
}

func getLength(l *ListNode) int {
	if l == nil {
		return 0
	}
	var (
		cur    = l
		length = 0
	)
	for {
		if cur == nil {
			break
		}
		length++
		cur = cur.Next
	}
	return length
}

func newList(l []int) *ListNode {
	head := &ListNode{}
	p := head
	for i, num := range l {
		p.Val = num
		if i == len(l)-1 {
			break
		}
		p.Next = &ListNode{}
		p = p.Next
	}
	return head
}

func printList(l *ListNode) {
	if l == nil {
		fmt.Println("")
		return
	}
	var (
		cur = l
		res []int
	)
	for {
		if cur == nil {
			break
		}
		res = append(res, cur.Val)
		cur = cur.Next
	}
	fmt.Println(res)
}

func main() {
	l1 := newList([]int{2, 4, 3})
	l2 := newList([]int{7, 0, 8})
	printList(addTwoNumbers(l1, l2))
}
