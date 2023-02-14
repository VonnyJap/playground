package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	adjacent := map[int]int{}
	adjacent[len(gas)-1] = 0
	for i := 0; i < len(gas)-1; i++ {
		adjacent[i] = i + 1
	}

	for start := range adjacent {
		graph := &graph{
			adjacent: adjacent,
			visited:  map[int]bool{},
			gas:      gas,
			cost:     cost,
			target:   start,
		}
		if graph.loopBack(start, gas[start]) {
			return start
		}
	}
	return -1
}

type graph struct {
	adjacent  map[int]int
	visited   map[int]bool
	gas, cost []int
	target    int
}

func (g *graph) loopBack(start int, gasTank int) bool {
	if start == g.target && g.visited[start] {
		return true
	}
	g.visited[start] = true
	next := g.adjacent[start]

	if gasTank < g.cost[start] {
		return false
	}

	return g.loopBack(next, gasTank-g.cost[start]+g.gas[next])
}

// it can only move clockwise so 1-2-3-4-5
// so create adjacent map as such [0->1,1->2,2->3,3->4,4->5,5->0]
// and also visited map - so we can mark the starting as visited and we should never get in again
// and then create a dfs util for this
