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
