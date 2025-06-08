package chapter5

import "errors"

type DLLNode[T any] struct {
	value T
	prev  *DLLNode[T]
	next  *DLLNode[T]
}

func (self *DLLNode[T]) InsertAfter(value T) {
	n := new(DLLNode[T])

	n.value = value
	n.prev = self
	n.next = self.next

	n.prev.next = n
	n.next.prev = n
}

// ///////////////////////////////////////////////////////////////////
// flow: Tail -> Head
type DblLinkList[T any] struct {
	length int
	head   *DLLNode[T]
	tail   *DLLNode[T]
}

func NewDblLinkList[T any]() *DblLinkList[T] {
	x := new(DblLinkList[T])
	x.length = 0
	x.head = new(DLLNode[T])
	x.tail = new(DLLNode[T])
	x.head.next = nil
	x.head.prev = x.tail
	x.tail.next = x.head
	x.tail.prev = nil

	return x
}

func (self *DblLinkList[T]) InsertTail(value T) {
	self.tail.InsertAfter(value)
	self.length++
}

func (self *DblLinkList[T]) RemoveHead() (T, error) {
	if self.length == 0 {
		var bogus T
		return bogus, errors.New("DblLinkList is empty")
	}

	n := self.head.prev

	nn := n.prev
	nn.next = self.head
	nn.next.prev = nn

	self.length--

	/* should not be needed */
	n.next = nil
	n.prev = nil

	return n.value, nil
}

func (self *DblLinkList[T]) MoveToTail(n *DLLNode[T]) {
	n.prev.next = n.next
	n.next.prev = n.prev

	n.next = self.tail.next
	n.prev = self.tail

	n.next.prev = n
	n.prev.next = n
}

// //////////////////////////////////////////////////////////////////
type KeyValuePair[T any] struct {
	key   int
	value T
}

type LRUCache[T any] struct {
	max_n int
	m     map[int]*DLLNode[KeyValuePair[T]]
	hist  *DblLinkList[KeyValuePair[T]]
}

func NewLRUCache[T any](max_n int) *LRUCache[T] {
	c := new(LRUCache[T])
	c.max_n = max_n
	c.m = make(map[int]*DLLNode[KeyValuePair[T]])
	c.hist = NewDblLinkList[KeyValuePair[T]]()
	return c
}

func (self *LRUCache[T]) Set(key int, value T) {
	n, ok := self.m[key]

	in := KeyValuePair[T]{key, value}

	if ok {
		n.value = in
		self.hist.MoveToTail(n)
	} else {
		self.hist.InsertTail(in)
		self.m[key] = self.hist.tail.next
	}

	if self.hist.length > self.max_n {
		out, _ := self.hist.RemoveHead()
		delete(self.m, out.key)
	}
}

func (self *LRUCache[T]) Get(key int) (T, error) {
	n, ok := self.m[key]

	if ok {
		self.hist.MoveToTail(n)
		return n.value.value, nil
	} else {
		var bogus T
		return bogus, errors.New("key not in cache")
	}
}
