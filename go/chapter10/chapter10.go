package chapter10

import "math"

type Edge struct {
	directed bool
	src      int
	dst      int
	weight   float64
}

type Graph struct {
	num_vertices int
	vertex_label []string
	num_edges    int
	edge         []Edge
	nbrs_made    bool
	nbrs         map[int][]int
	nbr_edges    map[int][]int // index into edge
}

func NewGraph() *Graph {
	g := new(Graph)
	g.num_vertices = 0
	g.num_edges = 0
	g.nbrs_made = false
	return g
}

func (g *Graph) AddVertex(label string) int {
	g.num_vertices += 1
	g.vertex_label = append(g.vertex_label, label)
	g.nbrs_made = false
	return g.num_vertices - 1
}

func (g *Graph) AddEdge(src int, dst int) {
	e := Edge{directed: false, src: src, dst: dst, weight: 0.0}

	g.num_edges += 1
	g.edge = append(g.edge, e)
	g.nbrs_made = false
}

func (g *Graph) AddWeightedEdge(src int, dst int, weight float64) {
	e := Edge{directed: false, src: src, dst: dst, weight: weight}

	g.num_edges += 1
	g.edge = append(g.edge, e)
	g.nbrs_made = false
}

func (g *Graph) AddDirectedEdge(src int, dst int) {
	e := Edge{directed: true, src: src, dst: dst, weight: 0.0}

	g.num_edges += 1
	g.edge = append(g.edge, e)
	g.nbrs_made = false
}

func (g *Graph) AddWeightedDirectedEdge(src int, dst int, weight float64) {
	e := Edge{directed: true, src: src, dst: dst, weight: weight}

	g.num_edges += 1
	g.edge = append(g.edge, e)
	g.nbrs_made = false
}

func (g *Graph) generate_nbrs() {
	if g.nbrs_made {
		return
	}

	g.nbrs = make(map[int][]int)
	g.nbr_edges = make(map[int][]int)
	for i := range g.num_vertices {
		g.nbrs[i] = make([]int, 0)
		g.nbr_edges[i] = make([]int, 0)
	}

	for idx, e := range g.edge {
		g.nbrs[e.src] = append(g.nbrs[e.src], e.dst)
		if !e.directed {
			g.nbrs[e.dst] = append(g.nbrs[e.dst], e.src)
		}
		g.nbr_edges[e.src] = append(g.nbr_edges[e.src], idx)
	}

	g.nbrs_made = true
}

func (g *Graph) OutDegree(v int) int {
	g.generate_nbrs()
	return len(g.nbrs[v])
}

func (g *Graph) Dijkstra(src int) []float64 {
	if !g.nbrs_made {
		g.generate_nbrs()
	}
	d := make([]float64, g.num_vertices)

	for v := range g.num_vertices {
		d[v] = math.MaxFloat64
	}

	d[src] = 0.0
	var stack []int
	stack = append(stack, src)

	for len(stack) > 0 {
		v := stack[0]
		stack = stack[1:]

		for _, e_idx := range g.nbr_edges[v] {
			e := g.edge[e_idx]
			dst := e.dst
			dd := d[v] + e.weight

			if dd < d[dst] {
				d[dst] = dd
				stack = append(stack, dst)
			}
		}
	}

	return d
}
