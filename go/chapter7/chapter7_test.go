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
