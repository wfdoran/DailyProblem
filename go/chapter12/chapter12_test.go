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

func TestRegEx(t *testing.T) {
	if !RegEx("ra.", "ray") {
		t.Errorf("1")
	}
	if !RegEx("*at", "chat") {
		t.Errorf("2")
	}
	if RegEx("*z", "hello") {
		t.Errorf("3")
	}
	if RegEx("*z.", "tubz") {
		t.Errorf("4")
	}
	if !RegEx("*z*", "tubz") {
		t.Errorf("5")
	}
}

func TestArrayMinMax(t *testing.T) {
	a := []int{4, 2, 7, 5, -1, 3, 6}
	fmt.Println(ArrayMinMax(a))

}
