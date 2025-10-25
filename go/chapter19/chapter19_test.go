package chapter19

import (
	"fmt"
	"testing"
)

func TestFindSlow(t *testing.T) {
	s := "abracadabra"
	target := "abr"
	fmt.Println(RabinKarp(target, s))
}
