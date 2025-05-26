package chapter3

import "fmt"

type LinkedListNode[T any] struct {
	val  T
	next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	start *LinkedListNode[T]
	end   *LinkedListNode[T]
}

func LinkedListNew[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil}
}

func LinkedListNodeNew[T any](val T) LinkedListNode[T] {
	return LinkedListNode[T]{val, nil}
}

func (x *LinkedList[T]) Append(val T) {
	node := LinkedListNodeNew[T](val)

	if x.start == nil {
		x.start = &node
		x.end = &node
	} else {
		x.end.next = &node
		x.end = &node
	}
}

func (x *LinkedList[T]) Reverse() {
	if x.start == nil {
		return
	}

	var orig *LinkedListNode[T] = x.start
	var prev *LinkedListNode[T] = nil
	var curr *LinkedListNode[T] = x.start

	for {
		temp := curr.next
		curr.next = prev
		prev = curr

		if temp == nil {
			x.start = curr
			x.end = orig
			break
		}
		curr = temp
	}
}

func (x *LinkedList[T]) Display() {
	for n := x.start; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}
