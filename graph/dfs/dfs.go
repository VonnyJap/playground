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

// what do I want to do when detect a cycle?
// from this point I want to do dfs and check if it is in path
func (g *Graph) hasCycle(v int, path []int) bool {

	g.Visited[v] = true
	path = append(path, v)
	for _, node := range g.Edges[v] {
		if inSlice(path, node) {
			return true
		}
		if !g.Visited[v] {
			if g.hasCycle(node, path) {
				return true
			}
		}
	}
	return false
}

func inSlice(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
