package chapter5

import (
	"errors"
	"fmt"
)

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

func BestCut(bricks [][]int) int {
	m := make(map[int]int)

	row_len := -1
	for _, row := range bricks {
		pos := 0

		for _, a := range row {
			pos += a
			m[pos] += 1
		}

		if row_len == -1 {
			row_len = pos
		} else {
			if row_len != pos {
				fmt.Println("Errr!!!!")
			}
		}
	}

	max_cut := 0
	max_val := 0

	for key, val := range m {
		if key != row_len && val > max_val {
			max_cut = key
			max_val = val
		}
	}
	return max_cut
}

type SparseArray[T any] struct {
	size         int
	m            map[int]T
	have_default bool
	default_val  T
}

func SparseArrayInit[T any](size int) *SparseArray[T] {
	out := new(SparseArray[T])
	out.size = size
	out.m = make(map[int]T)
	out.have_default = false
	return out
}

func (a *SparseArray[T]) SetDefault(val T) {
	a.have_default = true
	a.default_val = val
}

func (a *SparseArray[T]) Set(i int, val T) error {
	if i < 0 || i >= a.size {
		return errors.New("invalid i")
	}
	a.m[i] = val
	return nil
}

func (a *SparseArray[T]) Get(i int) (T, error) {
	var bogus T
	if i < 0 || i >= a.size {
		return bogus, errors.New("invalid i")
	}

	val, ok := a.m[i]
	if ok {
		return val, nil
	}
	if a.have_default {
		return a.default_val, nil
	}
	return bogus, errors.New("no value for this i")
}
