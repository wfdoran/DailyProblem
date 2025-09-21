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

func TestDijkstra(t *testing.T) {
	g := NewGraph()

	v0 := g.AddVertex("0")
	v1 := g.AddVertex("1")
	v2 := g.AddVertex("2")
	v3 := g.AddVertex("3")
	v4 := g.AddVertex("4")
	v5 := g.AddVertex("5")

	g.AddWeightedDirectedEdge(v0, v1, 5.0)
	g.AddWeightedDirectedEdge(v0, v2, 3.0)
	g.AddWeightedDirectedEdge(v0, v5, 4.0)
	g.AddWeightedDirectedEdge(v1, v3, 8.0)
	g.AddWeightedDirectedEdge(v2, v3, 1.0)
	g.AddWeightedDirectedEdge(v3, v5, 10.0)
	g.AddWeightedDirectedEdge(v3, v4, 5.0)

	d := g.Dijkstra(v0)
	fmt.Println(d)
}
