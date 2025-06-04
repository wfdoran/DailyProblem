package chapater4

import (
	"cmp"
	"errors"
)

type Stack[T any] struct {
	stack []T
}

func StackNew[T any]() *Stack[T] {
	a := new(Stack[T])
	return a
}

func (a *Stack[T]) Push(val T) {
	a.stack = append(a.stack, val)
}

func (a *Stack[T]) Length() int {
	return len(a.stack)
}

func (a *Stack[T]) Pop() (T, error) {
	a_len := len(a.stack)

	if a_len == 0 {
		var temp T
		return temp, errors.New("Stack is empty")
	}

	rv := a.stack[a_len-1]
	a.stack = a.stack[:a_len-1]
	return rv, nil
}

func (a *Stack[T]) Peek() (T, error) {
	a_len := len(a.stack)

	if a_len == 0 {
		var temp T
		return temp, errors.New("Stack is empty")
	} else {
		return a.stack[a_len-1], nil
	}
}

type MaxStack[T cmp.Ordered] struct {
	vals *Stack[T]
	maxs *Stack[T]
}

func MaxStackNew[T cmp.Ordered]() *MaxStack[T] {
	a := new(MaxStack[T])
	a.vals = new(Stack[T])
	a.maxs = new(Stack[T])
	return a
}

func (a *MaxStack[T]) Push(val T) {
	a.vals.Push(val)

	curr, err := a.maxs.Peek()
	if err != nil || val > curr {
		a.maxs.Push(val)
	} else {
		a.maxs.Push(curr)
	}
}

func (a *MaxStack[T]) Pop() (T, error) {
	val, err := a.vals.Pop()
	a.maxs.Pop()

	return val, err
}

func (a *MaxStack[T]) Peek() (T, error) {
	return a.vals.Peek()
}

func (a *MaxStack[T]) Max() (T, error) {
	return a.maxs.Peek()
}

func IsBalanced(s string) bool {
	a := StackNew[rune]()
	partner := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			a.Push(v)
		}

		if v == ')' || v == '}' || v == ']' {
			r, err := a.Pop()
			if err != nil {
				return false
			}

			if r != partner[v] {
				return false
			}
		}
	}
	return a.Length() == 0
}
