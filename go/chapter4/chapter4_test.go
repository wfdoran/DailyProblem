package chapater4

import (
	"fmt"
	"testing"
)

func TestStackPushPop(t *testing.T) {
	a := StackNew[int]()

	for i := range 3 {
		a.Push(i)
	}

	for {
		v, err := a.Pop()
		if err == nil {
			fmt.Println(v)
		} else {
			break
		}
	}
	fmt.Println()
}

func TestMaxStack(t *testing.T) {
	a := MaxStackNew[int]()

	vals := []int{1, 3, 5, 2, 4}
	for _, v := range vals {
		a.Push(v)
	}

	for {
		maxv, err := a.Max()
		if err == nil {
			fmt.Println(maxv)
		} else {
			break
		}
		a.Pop()
	}
}

func TestIsBalanced(t *testing.T) {
	if !IsBalanced("([])[]({})") {
		t.Error("Ex 1")
	}

	if IsBalanced("([()]") {
		t.Error("Ex 2")
	}

	if IsBalanced("((()") {
		t.Error("Ex 3")
	}
}
