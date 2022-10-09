package main

import "fmt"

func main() {
	g := Graph{rs: map[int][]int{}}
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	fmt.Printf("%#v\n", g.rs)
	fmt.Printf("%#v\n", g.bfs(2))

}

type Graph struct {
	rs map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	if _, ok := g.rs[u]; !ok {
		g.rs[u] = []int{v}
		return
	}
	g.rs[u] = append(g.rs[u], v)
}

func (g *Graph) bfs(start int) []int {
	// create visited thing
	var visited = map[int]bool{}
	visited[start] = true
	var queue = []int{}
	queue = append(queue, start)
	var bfs = []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		bfs = append(bfs, current)

		// visit all nodes adjacent to current and add them to queue
		for _, node := range g.rs[current] {
			_, ok := visited[node]
			if ok {
				continue
			}
			queue = append(queue, node)
			visited[node] = true
		}
	}

	return bfs
}
