package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	var now = head
	for now != nil {
		next = now.Next
		now.Next = pre
		pre = now
		now = next
	}
	return pre
}
