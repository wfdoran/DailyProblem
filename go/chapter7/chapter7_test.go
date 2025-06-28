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
