package main

func main() {}

type Graph struct {
	Vertices int
	Edges    map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	if _, ok := g.Edges[u]; ok {
		g.Edges[u] = append(g.Edges[u], v)
		return
	}
	g.Edges[u] = []int{v}
}

func (g *Graph) findParent(parent []int, vertice int) int {
	if parent[vertice] == vertice {
		return vertice
	}
	return g.findParent(parent, parent[vertice])
}

func (g *Graph) union(parent []int, x, y int) []int {
	parent[x] = y
	return parent
}

func (g *Graph) isCyclic() bool {

	parent := make([]int, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		parent[i] = i
	}

	for i := range g.Edges {
		for j := range g.Edges[i] {
			x := g.findParent(parent, i)
			y := g.findParent(parent, j)
			if x == y {
				return true
			}
			parent = g.union(parent, x, y)
		}
	}

	return false
}
