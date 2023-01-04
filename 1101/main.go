package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestAcq(
		[][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}},
		6,
	))
	fmt.Println(earliestAcq([][]int{{9, 3, 0}, {0, 2, 1}, {8, 0, 1}, {1, 3, 2}, {2, 2, 0}, {3, 3, 1}}, 4))
}

func earliestAcq(logs [][]int, n int) int {

	// need to sort logs
	sort.Slice(logs, func(i, j int) bool { return logs[i][0] < logs[j][0] })

	graph := &Graph{
		Edges: map[int][]int{},
	}
	timestamp := -1
	for _, log := range logs {
		graph.AddEdge(log[1], log[2])
		graph.Visited = map[int]bool{}
		result := graph.DFS(0)
		fmt.Println(result)
		if len(result) == n {
			timestamp = log[0]
			break
		}
	}
	return timestamp
}

// what about using dfs with some kinds of memoization?
// like maybe we can add visited and the collected each time we run?

type Graph struct {
	Edges   map[int][]int
	Visited map[int]bool
}

func (g *Graph) AddEdge(u, v int) {
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = []int{v}
	} else {
		g.Edges[u] = append(g.Edges[u], v)
	}
	if _, ok := g.Edges[v]; !ok {
		g.Edges[v] = []int{u}
	} else {
		g.Edges[v] = append(g.Edges[v], u)
	}
	return
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
