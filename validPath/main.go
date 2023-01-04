package main

import "fmt"

func main() {
	// fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	// fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	fmt.Println(validPath(10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5))

}

func validPath(n int, edges [][]int, source int, destination int) bool {

	if n == 1 {
		return true
	}

	g := &graph{
		vertex:  n,
		edges:   map[int][]int{},
		visited: map[int]bool{},
	}

	for _, ed := range edges {
		if _, ok := g.edges[ed[0]]; !ok {
			g.edges[ed[0]] = []int{}
		}
		if _, ok := g.edges[ed[1]]; !ok {
			g.edges[ed[1]] = []int{}
		}
		g.edges[ed[0]] = append(g.edges[ed[0]], ed[1])
		g.edges[ed[1]] = append(g.edges[ed[1]], ed[0])
	}

	fmt.Println(g.edges)

	return g.dfs(source, destination)
}

type graph struct {
	vertex  int
	edges   map[int][]int
	visited map[int]bool
}

func (g *graph) dfs(source int, destination int) bool {

	g.visited[source] = true

	for _, edge := range g.edges[source] {
		if edge == destination {
			return true
		}
		if _, ok := g.visited[edge]; ok {
			continue
		}
		if g.dfs(edge, destination) {
			return true
		}
	}

	return false
}
