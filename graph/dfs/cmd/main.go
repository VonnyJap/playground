package main

import (
	"fmt"
)

func main() {
	g := &Graph{
		Edges:   map[int][]int{},
		Visited: map[int]bool{},
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	// g.AddEdge(1, 2)
	// g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	// g.AddEdge(3, 3)
	fmt.Printf("%#v\n", g.Edges)
	fmt.Printf("%#v\n", g.hasCycle(2, []int{}))
}

// how to handle disconnected graph then
// for all the nodes in the graph - we will check if visited - if not then we run dfs
// if we need to increment dfs that means there are multiple disconnected graphs

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
		if !g.Visited[node] {
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
