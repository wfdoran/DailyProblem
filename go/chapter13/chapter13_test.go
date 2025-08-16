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

func TestDecodeCount(t *testing.T) {
	fmt.Println(DecodeCount("111"))
	fmt.Println(DecodeCount("101"))
	fmt.Println(DecodeCount("131"))
	fmt.Println(DecodeCount("123711202231138152915281327239123"))

}
