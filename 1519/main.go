package main

import "fmt"

func main() {

	fmt.Println(countSubTrees(4, [][]int{{0, 1}, {0, 2}, {2, 3}}, "abae"))
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(countSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))

}

type Graph struct {
	edges   map[int][]int
	labels  []string
	visited map[int]bool
}

// do dfs for each start point so we can get the subtree
// we can do memoization on this
// for each subtree we can do check if the labels are same

func (g *Graph) DFS(v int) []int {
	if _, ok := g.visited[v]; ok {
		return []int{}
	}
	g.visited[v] = true
	var set = []int{}
	set = append(set, v)

	for _, node := range g.edges[v] {
		set = append(set, g.DFS(node)...)
	}
	return set
}

func countSubTrees(n int, edges [][]int, labels string) []int {

	// first lets init the Graph
	graph := &Graph{
		edges: map[int][]int{},
		// visited: map[int]bool{},
	}

	for _, edge := range edges {
		if _, ok := graph.edges[edge[0]]; ok {
			graph.edges[edge[0]] = append(graph.edges[edge[0]], edge[1])
			continue
		}
		graph.edges[edge[0]] = []int{edge[1]}
	}

	// fmt.Println(graph.edges)
	// build the edges first
	result := make([]int, n)

	start := 0
	for start < n {
		graph.visited = map[int]bool{}
		sub := graph.DFS(start)
		// fmt.Println(start, sub)
		for _, child := range sub {
			if labels[start] == labels[child] {
				result[start]++
			}
		}
		start++
	}

	return result
}
