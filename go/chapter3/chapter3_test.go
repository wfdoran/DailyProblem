package chapter3

import (
	"fmt"
	"testing"
)

func TestLLReverse(t *testing.T) {
	a := LinkedListNew[int]()

	for i := range 5 {
		a.Append(i)
	}

	a.Reverse()
	fmt.Println(LLToInt(a))
}

func TestLLAddInts(t *testing.T) {
	a := LLFromInt(99)
	b := LLFromInt(25)
	c := LLAddInts(a, b)
	fmt.Println(LLToInt(c))
}

func TestSwapNextTwo(t *testing.T) {
	a := LLFromInt(37)
	SwapNextTwo(&a.start)
	fmt.Println(LLToInt(a))
}

func TestLLAlternate(t *testing.T) {
	a := LLFromInt(54321)
	LLAlternate(a)
	fmt.Println(LLToInt(a))
}

func TestLLLength(t *testing.T) {
	{
		a := LLFromInt(123456)
		expect := 6
		result := a.Length()
		if result != expect {
			t.Error("LL Length: result = ", result, "expect = ", expect)
		}
	}
}
