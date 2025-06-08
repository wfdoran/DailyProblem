package chapter5

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache[float64](4)

	lru.Set(1, 1.0)
	lru.Set(2, 4.0)
	lru.Set(3, 9.0)
	lru.Set(4, 16.0)
	lru.Set(5, 25.0)

	_, err := lru.Get(1)
	fmt.Println(err)

	v, _ := lru.Get(2)
	fmt.Println(v)

	lru.Set(6, 36.0)
	v, _ = lru.Get(2)
	fmt.Println(v)

	_, err = lru.Get(3)
	fmt.Println(err)
}
