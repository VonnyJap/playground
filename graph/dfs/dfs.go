package dfs

type Graph struct {
	Edges   map[int][]int
	Visited map[int]bool
}

func (g *Graph) AddEdge(u, v int) {
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = []int{v}
		return
	}
	g.Edges[u] = append(g.Edges[u], v)
}

func (g *Graph) DFS(v int) []int {
	if _, ok := g.Visited[v]; ok {
		return []int{}
	}
	g.Visited[v] = true
	var set = []int{}
	set = append(set, v)

	for _, node := range g.Edges[v] {
		set = append(set, g.DFS(node)...)
	}
	return set
}
