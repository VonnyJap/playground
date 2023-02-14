package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
	fmt.Println(findOrder(2, [][]int{{0, 1}, {1, 0}}))
	fmt.Println(findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {

	g := Graph{
		adjacent: map[int][]int{},
		visited:  map[int]bool{},
		orders:   []int{},
	}
	for _, req := range prerequisites {
		if _, ok := g.adjacent[req[0]]; ok {
			g.adjacent[req[0]] = append(g.adjacent[req[0]], req[1])
			continue
		}
		g.adjacent[req[0]] = []int{req[1]}
	}

	for i := 0; i < numCourses; i++ {
		g.dfs(i)
		if g.hasCycle {
			return []int{}
		}
	}

	return g.orders
}

type Graph struct {
	adjacent map[int][]int
	visited  map[int]bool
	orders   []int
	hasCycle bool
}

// how to implement?
// loop from i:=0 => i < n
// what is the stopping condition here? - if visited then no need to visit again

func (g *Graph) dfs(start int) {
	if val, ok := g.visited[start]; ok {
		// ok to have this as long as it is not true
		if val {
			g.hasCycle = true
		}
		return
	}
	g.visited[start] = true

	for _, course := range g.adjacent[start] {
		g.dfs(course)
		if g.hasCycle {
			return
		}
	}

	g.orders = append(g.orders, start)
	// backtrack
	g.visited[start] = false

}

// need to consider a cycle?

// [[0,1],[0,2],[1,2]]
// 3
