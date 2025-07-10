package chapter9

import (
	"fmt"
	"testing"
)

func TestHeap1(t *testing.T) {
	var a []int

	a = HeapInsert[int](a, 10)
	a = HeapInsert[int](a, 6)
	a = HeapInsert[int](a, 7)
	a = HeapInsert[int](a, 2)
	a = HeapInsert[int](a, 5)

	fmt.Println(a)

	for range 5 {
		var min_val int
		min_val, a = HeapDelete(a)
		fmt.Println(a, min_val)
	}
}
