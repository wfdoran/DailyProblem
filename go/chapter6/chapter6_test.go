package chapter6

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	var root *TreeNode[int, int] = nil

	root = root.Insert(30, 0)
	root = root.Insert(20, 0)
	root = root.Insert(40, 0)
	root = root.Insert(70, 0)
	root = root.Insert(60, 0)
	root = root.Insert(80, 0)

	fmt.Println(root)
	fmt.Println(root.height)
	fmt.Println("======================================")
}

func TestBST2(t *testing.T) {
	var root *TreeNode[int, int] = nil

	for i := range 1000 {
		x := i + 1
		root = root.Insert(x, x*x)
	}

	// fmt.Println(root)
	fmt.Println(root.height)
	fmt.Println("======================================")
}

func TestBSTRemove(t *testing.T) {
	var root *TreeNode[int, int] = nil
	max := 30
	for i := range max {
		root = root.Insert(i, i*i)
	}

	for i := range max {
		// fmt.Println(root)
		ok, d, ret := root.Remove(i)
		if !ok {
			t.Error("AAA")
		}
		if d != i*i {
			t.Error("BBB")
		}
		root = ret
		root = root.Insert(i, d)
	}
	fmt.Println("======================================")
}

func TestBST3(t *testing.T) {
	n := 2000
	var root *TreeNode[int, int] = nil

	for i := range n {
		x := 1 + ((3 * i) % n)
		root = root.Insert(x, x*x)
	}

	fmt.Println(root.height)
	for i := 0; i < n; i += 2 {
		x := 1 + ((7 * i) % n)
		_, _, root = root.Remove(x)
	}

	fmt.Println(root.height)
	fmt.Println("======================================")
}

func TestBST4(t *testing.T) {
	var root *TreeNode[int, int] = nil

	n := 10
	for i := range n {
		x := 2*i + 1
		root = root.Insert(x, x*x)
	}

	fmt.Println(root)
	fmt.Println(root.key, root.left.key, root.right.key)

	for i := range 2*n + 1 {
		ok, floor, data := root.FindFloor(i)
		fmt.Println(i, ok, floor, data)
	}
	for i := range 2*n + 1 {
		ok, ceil, data := root.FindCeil(i)
		fmt.Println(i, ok, ceil, data)
	}
	fmt.Println("======================================")
}

func TestBST5(t *testing.T) {
	var root *TreeNode[int, int] = nil

	n := 20

	for i := range n / 2 {
		x := 1 + ((3 * i) % n)
		root = root.Insert(x, x*x)
	}

	for x := range root.Walk() {
		fmt.Println(x.key, x.data)
	}
	fmt.Println("======================================")
}
