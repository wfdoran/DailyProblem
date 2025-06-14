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
}
