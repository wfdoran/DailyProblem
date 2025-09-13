package chapter15

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDutchFlagSort(t *testing.T) {
	a := []rune{'G', 'B', 'R', 'R', 'B', 'R', 'R'}
	DutchFlagSort(a)
	fmt.Println(a)
}

func TestPancakeSort(t *testing.T) {
	a := []int{5, 7, 10, 1, 8, 23, 11, 15, 25, 2}
	PancakeSort(a)
	fmt.Println(a)
}

func TestRadixSort(t *testing.T) {
	n := 1000000
	a := make([]uint32, n)
	for i := range n {
		a[i] = rand.Uint32()
	}

	RadixSort(a)

	for i := range n - 1 {
		if a[i] > a[i+1] {
			t.Error("out of order")
		}
	}
}

func TestFindInSorted(t *testing.T) {
	a := []int{1, 2, 3, 5, 6, 7}

	fmt.Println(FindInSorted(a, 1))
	fmt.Println(FindInSorted(a, 4))
	fmt.Println(FindInSorted(a, 6))
	fmt.Println(FindInSorted(a, 7))
}

func TestFindSlow(t *testing.T) {
	a := []int{1, 2, 3, 5, 6, 7}

	fmt.Println(FindSlow(a, 1))
	fmt.Println(FindSlow(a, 4))
	fmt.Println(FindSlow(a, 6))
	fmt.Println(FindSlow(a, 7))
}

func TestFindInRotated(t *testing.T) {
	a := []int{18, 25, 2, 8, 10, 13}

	fmt.Println(FindInRotated(a, 8))
	fmt.Println(FindInRotated(a, 2))
}
