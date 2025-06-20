package chapter6

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	root := NewTreeNode[int, int](50, 0)

	root = root.Insert(30, 0)
	root = root.Insert(20, 0)
	root = root.Insert(40, 0)
	root = root.Insert(70, 0)
	root = root.Insert(60, 0)
	root = root.Insert(80, 0)

	fmt.Println(root)
	fmt.Println(root.height)
}

func TestBST2(t *testing.T) {
	root := NewTreeNode[int, int](0, 0)

	for i := range 1000 {
		x := i + 1
		root = root.Insert(x, x*x)
	}

	// fmt.Println(root)
	fmt.Println(root.height)
}

func TestBSTRemove(t *testing.T) {
	max := 30

	root := NewTreeNode[int, int](0, 0)
	for i := range max {
		x := i + 1
		d := x * x
		root = root.Insert(x, d)
	}

	for i := range max + 1 {
		// fmt.Println(root)
		ok, d, ret := root.Remove(i)
		if !ok {
			t.Error("AAA")
		}
		if d != i*i {
			t.Error("BBB")
		}
		root = ret

		// fmt.Println(i, ok, d)
		// fmt.Println(root)

		root = root.Insert(i, d)
		// fmt.Println(root)
		// fmt.Println()
	}

}
