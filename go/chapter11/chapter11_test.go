package chapter11

import (
	"fmt"
	"testing"
)

func TestGraph1(t *testing.T) {
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
