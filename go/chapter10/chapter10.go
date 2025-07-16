package chapter10

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
}

func NewGraph() *Graph {
	g := new(Graph)
	g.num_vertices = 0
	g.num_edges = 0
	return g
}

func (g *Graph) AddVertex(label string) {
	g.num_vertices += 1
	g.vertex_label = append(g.vertex_label, label)
}

func (g *Graph) AddEdge(src int, dst int) {
	e := Edge{directed: false, src: src, dst: dst, weight: 0.0}

	g.num_edges += 1
	g.edge = append(g.edge, e)
}

func (g *Graph) AddWeightedEdge(src int, dst int, weight float64) {
	e := Edge{directed: false, src: src, dst: dst, weight: weight}

	g.num_edges += 1
	g.edge = append(g.edge, e)
}

func (g *Graph) AddDirectedEdge(src int, dst int) {
	e := Edge{directed: true, src: src, dst: dst, weight: 0.0}

	g.num_edges += 1
	g.edge = append(g.edge, e)
}

func (g *Graph) AddWeightedDirectedEdge(src int, dst int, weight float64) {
	e := Edge{directed: true, src: src, dst: dst, weight: weight}

	g.num_edges += 1
	g.edge = append(g.edge, e)
}
