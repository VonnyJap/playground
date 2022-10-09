package main

import (
	"fmt"

	"github.com/vonny/interview/graph/dfs"
)

func main() {
	g := &dfs.Graph{
		Edges:   map[int][]int{},
		Visited: map[int]bool{},
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	fmt.Printf("%#v\n", g.Edges)
	fmt.Printf("%#v\n", g.DFS(2))
}

// how to handle disconnected graph then
// for all the nodes in the graph - we will check if visited - if not then we run dfs
// if we need to increment dfs that means there are multiple disconnected graphs
