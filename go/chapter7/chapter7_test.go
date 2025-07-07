package chapter7

import (
	"fmt"
	"testing"
)

func TestTrie1(t *testing.T) {
	root := TrieRoot()

	root.Insert("bear")
	root.Insert("cat")
	root.Insert("coat")
	root.Insert("dog")

	fmt.Println(root)

}

func TestTrie2(t *testing.T) {
	root := TrieRoot()

	root.Insert("bear")
	root.Insert("cat")
	root.Insert("coat")
	root.Insert("catalog")
	root.Insert("dog")

	fmt.Println(root.Find("cat"))
	fmt.Println(root.Find("do"))
	fmt.Println(root.Find("cata"))
}

func TestTrie3(t *testing.T) {
	root := TrieRoot()

	root.Insert("dog")
	root.Insert("deer")
	root.Insert("deal")
	root.Insert("ant")
	root.Insert("anter")

	fmt.Println(root.AutoComplete("de"))
	fmt.Println(root.AutoComplete("ant"))

}

func TestTrie4(t *testing.T) {
	root := TrieRoot()

	root.InsertValue("dog", 1)
	root.InsertValue("donut", 2)
	root.InsertValue("delay", 4)

	fmt.Println(root.Sum("d"))
	fmt.Println(root.Sum("do"))
	fmt.Println(root.Sum("doo"))
}

func TestTrie5(t *testing.T) {
	a := []uint{4, 6, 7}
	fmt.Println("max = ", MaxXOR(a))

}
