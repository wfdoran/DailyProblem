package chapater4

import "errors"

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
