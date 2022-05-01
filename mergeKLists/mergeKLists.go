package mergeKLists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeArr []*ListNode

func (l ListNodeArr) Len() int {
	return len(l)
}

func (l ListNodeArr) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ListNodeArr) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l *ListNodeArr) Pop() any {
	if l.Len() == 0 {
		return nil
	}
	res := (*l)[l.Len()-1]
	(*l) = (*l)[0 : (*l).Len()-1]
	return res
}

func (l *ListNodeArr) Push(x any) {
	item, ok := x.(*ListNode)
	if !ok {
		return
	}
	*l = append(*l, item)
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := &ListNodeArr{}
	res := &ListNodeArr{}
	for _, list := range lists {
		if list != nil {
			(*arr) = append((*arr), list)
		}
	}
	heap.Init(arr)
	for len((*arr)) > 0 {
		x := heap.Pop(arr)
		if x == nil {
			continue
		}
		item, ok := x.(*ListNode)
		if !ok {
			continue
		}
		(*res) = append((*res), item)
		if item.Next != nil {
			heap.Push(arr, item.Next)
		}
	}
	for i := 0; i < len((*res)); i++ {
		if i < len((*res))-1 {
			(*res)[i].Next = (*res)[i+1]
		}
	}
	if len((*res)) > 0 {
		return (*res)[0]
	}
	return nil
}
