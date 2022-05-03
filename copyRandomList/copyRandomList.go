package copyrandomlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	r := make(map[*Node]*Node)
	var pre *Node
	h := head
	for head != nil {
		node := &Node{
			Val: head.Val,
		}
		m[node] = head
		r[head] = node
		if pre != nil {
			pre.Next = node
		}
		pre = node
		head = head.Next
	}
	newHead := r[h]
	for newHead != nil {
		newHead.Random = r[m[newHead].Random]
		newHead = newHead.Next
	}
	return r[h]
}
