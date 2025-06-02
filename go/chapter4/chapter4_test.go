package chapater4

import (
	"fmt"
	"testing"
)

func TestStackPushPop(t *testing.T) {
	a := StackNew[int]()

	for i := range 3 {
		a.Push(i)
	}

	for {
		v, err := a.Pop()
		if err == nil {
			fmt.Println(v)
		} else {
			break
		}
	}
}
