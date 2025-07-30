package chapter11

import (
	"fmt"
	"testing"
)

func TestFenwick(t *testing.T) {
	amts := []int{1, 2, 3, 4, 5}
	n := len(amts)
	a := make([]int, n+1)

	for i, amt := range amts {
		FenwickUpdate(a, i+1, amt)
	}

	fmt.Println(a)

	for i := 1; i <= n; i++ {
		fmt.Println(i, FenwickSum(a, i))
	}

	for i := 1; i <= n; i++ {
		fmt.Println(i, FenwickRange(a, i, i))
	}
	for i := 1; i <= n-1; i++ {
		fmt.Println(i, FenwickRange(a, i, i+1))
	}

}

func TestUnionFind(t *testing.T) {
	fmt.Println("==========================")
	uf := UnionFindInit(6)
	uf.Merge(0, 2)
	uf.Merge(1, 3)
	uf.Merge(2, 4)

	for i := range 6 {
		fmt.Println(i, uf.GetClass(i))
	}
}
