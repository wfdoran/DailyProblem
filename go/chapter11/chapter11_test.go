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

func TestBloomFilter(t *testing.T) {
	filter := BloomFilterInit(1024)

	v := uint64(1)
	for range 100 {
		v *= uint64(0x555555555555555)
		v += uint64(10)

		filter.Add(v)
	}

	fmt.Println("========================")
	v = uint64(1)
	for range 10 {
		for range 10 {
			v *= uint64(0x555555555555555)
			v += uint64(10)
		}

		fmt.Print(filter.Check(v), " ")
	}
	fmt.Println()

	v = uint64(2)
	for range 10 {
		v *= uint64(0x555555555555555)
		v += uint64(10)

		fmt.Print(filter.Check(v), " ")
	}
	fmt.Println()
}
