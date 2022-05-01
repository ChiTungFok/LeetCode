package main

import "container/heap"

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(i interface{}) {
	item, ok := i.(int)
	if !ok {
		return
	}
	*ih = append(*ih, item)
}

func (ih *IntHeap) Pop() interface{} {
	if len(*ih) == 0 {
		return nil
	}
	item := (*ih)[len(*ih)-1]
	(*ih) = (*ih)[0 : len(*ih)-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
			continue
		}
		top := heap.Pop(h)
		if num > top.(int) {
			heap.Push(h, num)
		} else {
			heap.Push(h, top)
		}
	}
	return heap.Pop(h).(int)
}
