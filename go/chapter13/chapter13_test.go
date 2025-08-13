package chapter13

import (
	"fmt"
	"testing"
)

func TestNumClimbs(t *testing.T) {
	for i := int64(1); i < int64(20); i++ {
		fmt.Println(NumClimbs(i))
	}
}
