package chapter3

import (
	"testing"
)

func TestLLReverse(t *testing.T) {
	ll := LinkedListNew[int]()

	for i := range 5 {
		ll.Append(i)
	}

	ll.Reverse()
	ll.Display()
}
