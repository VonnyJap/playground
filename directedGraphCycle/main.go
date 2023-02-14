package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
}

func minimumSemesters(n int, relations [][]int) int {

	hasPreqs := map[int]bool{}

	graph := &Graph{
		Vertices: []int{},
		Edges:    map[int][]int{},
		Visited:  map[int]bool{},
	}

	for i := 1; i <= n; i++ {
		graph.Vertices = append(graph.Vertices, i)
	}

	for _, relation := range relations {
		hasPreqs[relation[1]] = true
		if _, ok := graph.Edges[relation[0]]; ok {
			graph.Edges[relation[0]] = append(graph.Edges[relation[0]], relation[1])
			continue
		}
		graph.Edges[relation[0]] = []int{relation[1]}
	}

	max := 0
	for _, v := range graph.Vertices {
		if _, ok := hasPreqs[v]; ok {
			continue
		}
		depth := graph.traverse(v, []int{})
		if depth == -1 {
			return -1
		}
		if depth > max {
			max = depth
		}
	}

	return max
}

// create prereqs
// create graph and edges
// create visiting map
// create all vertices so we can loop through
// only start dfs from those dont have prereqs
// the dfs will return hasCycle and the level?

// Input: n = 3, relations = [[1,3],[2,3]]
// Input: n = 3, relations = [[1,3],[2,3],[3,4],[3,5]]

type Graph struct {
	Vertices []int
	Edges    map[int][]int
	Visited  map[int]bool
}

func (g *Graph) traverse(v int, path []int) int {

	g.Visited[v] = true
	path = append(path, v)

	depths := []int{}
	for _, node := range g.Edges[v] {
		if inSlice(path, node) {
			return -1
		}
		if !g.Visited[node] {
			depths = append(depths, g.traverse(node, path))
		}
	}

	if len(depths) == 0 {
		return 0
	}
	sort.Ints(depths)
	return depths[len(depths)-1]
}

func inSlice(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
