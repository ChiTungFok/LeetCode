package main

type List struct {
	len  int
	root *Element
}

type Element struct {
	pre   *Element
	next  *Element
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *List
	elements map[int]*Element
}

func NewList() *List {
	root := &Element{}
	root.pre = root
	root.next = root
	return &List{
		root: root,
	}
}

func (l *List) PushFront(key, value int) *Element {
	if l == nil {
		return nil
	}
	e := &Element{
		key:   key,
		value: value,
	}
	e.pre = l.root
	e.next = l.root.next
	e.next.pre = e
	l.root.next = e
	l.len++
	return e
}

func (l *List) MoveToFront(e *Element) {
	if e == nil || l == nil {
		return
	}
	e.pre.next = e.next
	e.next.pre = e.pre
	e.pre = l.root
	e.next = l.root.next
	e.next.pre = e
	l.root.next = e
}

func (l *List) Back() *Element {
	if l == nil || l.root == nil {
		return nil
	}
	if l.root.pre == l.root {
		return nil
	}
	return l.root.pre
}

func (l *List) Remove(e *Element) {
	if l == nil || e == nil {
		return
	}
	e.pre.next = e.next
	e.next.pre = e.pre
	e.next = nil
	e.pre = nil
	l.len--
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     NewList(),
		elements: make(map[int]*Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.elements[key]; ok {
		if e == nil {
			return -1
		}
		this.list.MoveToFront(e)
		return e.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.elements[key]; ok {
		if e == nil {
			return
		}
		e.value = value
		this.list.MoveToFront(e)
		return
	}
	if this.list.len < this.capacity {
		val := this.list.PushFront(key, value)
		this.elements[key] = val
		return
	}
	this.eliminate()
	e := this.list.PushFront(key, value)
	this.elements[key] = e
}

func (this *LRUCache) eliminate() {
	e := this.list.Back()
	this.list.Remove(e)
	if e == nil {
		return
	}
	delete(this.elements, e.key)
}

func main() {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
}
