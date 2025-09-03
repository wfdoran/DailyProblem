package chapter15

import (
	"fmt"
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
