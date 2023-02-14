package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestCellPath([][]int{{1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}}, 0, 0, 2, 0))
}

func shortestCellPath(matrix [][]int, sr, sc, tr, tc int) int {

	if len(matrix) == 0 {
		return 0
	}
	graph := &graph{
		matrix:  matrix,
		visited: map[coordinate]bool{},
		rows:    len(matrix),
		cols:    len(matrix[0]),
	}

	return graph.dfs(sr, sc, tr, tc)
}

// lets do memoization as well here?
type graph struct {
	matrix     [][]int
	visited    map[coordinate]bool
	rows, cols int
}

type coordinate struct {
	row, col int
}

// what can we do here?

func (g *graph) dfs(sr, sc, tr, tc int) int {

	g.visited[coordinate{sr, sc}] = true
	if sr == tr && sc == tc {
		return 0
	}

	adjacent := [][]int{}
	if sr-1 >= 0 && g.matrix[sr-1][sc] == 1 {
		adjacent = append(adjacent, []int{sr - 1, sc})
	}
	if sr+1 < g.rows && g.matrix[sr+1][sc] == 1 {
		adjacent = append(adjacent, []int{sr + 1, sc})
	}
	if sc-1 >= 0 && g.matrix[sr][sc-1] == 1 {
		adjacent = append(adjacent, []int{sr, sc - 1})
	}
	if sc+1 < g.cols && g.matrix[sr][sc+1] == 1 {
		adjacent = append(adjacent, []int{sr, sc + 1})
	}

	// need to fill this up
	// we check which one we can go from here and add to adjacent list

	// keep a min value or something and we keep on updating this?
	if len(adjacent) == 0 {
		return -1
	}

	min := int(math.Inf(1))
	for _, adj := range adjacent {
		if _, ok := g.visited[coordinate{adj[0], adj[1]}]; !ok {
			depth := g.dfs(adj[0], adj[1], tr, tc)
		}
	}

	return min
}
