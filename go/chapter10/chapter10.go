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
	nbrs_made    bool
	nbrs         map[int][]int
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
	for i := range g.num_vertices {
		g.nbrs[i] = make([]int, 0)
	}

	for _, e := range g.edge {
		g.nbrs[e.src] = append(g.nbrs[e.src], e.dst)
		if !e.directed {
			g.nbrs[e.dst] = append(g.nbrs[e.dst], e.src)
		}
	}

	g.nbrs_made = true
}

func (g *Graph) OutDegree(v int) int {
	g.generate_nbrs()
	return len(g.nbrs[v])
}
