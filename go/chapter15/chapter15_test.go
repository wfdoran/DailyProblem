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
