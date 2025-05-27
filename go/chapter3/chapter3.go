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

func LLFromInt(x int) *LinkedList[int] {
	a := LinkedListNew[int]()

	for x > 0 {
		a.Append(x % 10)
		x /= 10
	}

	return a
}

func LLToInt(a *LinkedList[int]) int {
	rv := 0
	base := 1

	an := a.start
	for an != nil {
		rv += base * an.val
		base *= 10
		an = an.next
	}
	return rv
}

func LLAddInts(a *LinkedList[int], b *LinkedList[int]) *LinkedList[int] {
	c := LinkedListNew[int]()

	carry := 0
	an := a.start
	bn := b.start

	for an != nil || bn != nil || carry > 0 {
		a_value := 0
		if an != nil {
			a_value = an.val
			an = an.next
		}

		b_value := 0
		if bn != nil {
			b_value = bn.val
			bn = bn.next
		}

		total := a_value + b_value + carry

		c.Append(total % 10)
		carry = total / 10
	}

	return c
}
