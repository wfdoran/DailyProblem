package chapter10

import (
	"fmt"
	"testing"
)

func TestGraph1(t *testing.T) {
	g := NewGraph()

	v0 := g.AddVertex("Hello")
	v1 := g.AddVertex("X")
	v2 := g.AddVertex("World")
	v3 := g.AddVertex("Y")

	g.AddEdge(v0, v1)
	g.AddEdge(v0, v2)
	g.AddEdge(v1, v2)
	g.AddEdge(v0, v3)

	for v := range []int{v0, v1, v2, v3} {
		fmt.Println(g.OutDegree(v))
	}
}
