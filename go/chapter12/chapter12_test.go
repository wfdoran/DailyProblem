package chapter12

import (
	"fmt"
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	for x := range TowerOfHanoi(3, 1, 3) {
		fmt.Println(x)
	}

}
