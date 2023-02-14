package main

import "fmt"

func main() {
	// fmt.Println(uniquePaths(2, 2))
	fmt.Println(uniquePaths(3, 2))

}

type coordinate struct {
	row, col int
}

func uniquePaths(m int, n int) int {

	g := &graph{
		row:     m,
		col:     n,
		visited: map[coordinate]bool{},
	}
	count := 0
	g.dfs(0, 0, &count)
	fmt.Println(count)
	return 0
}

type graph struct {
	row     int
	col     int
	visited map[coordinate]bool
}

// this solution is right because it is possible to get another path with backtrack
func (g *graph) dfs(startRow, startCol int, count *int) {

	if startRow < 0 || startRow >= g.row || startCol < 0 || startCol >= g.col {
		return
	}
	if val, ok := g.visited[coordinate{startRow, startCol}]; ok && val {
		return
	}
	if startRow == g.row-1 && startCol == g.col-1 {
		fmt.Println("startRow, startCol: ", startRow, startCol, "bang")
		*count = *count + 1
		fmt.Println("count: ", *count)
		return
	}
	fmt.Println("startRow, startCol: ", startRow, startCol)

	adjacent := [][]int{
		{startRow + 1, startCol},
		{startRow - 1, startCol},
		{startRow, startCol + 1},
		{startRow, startCol - 1},
	}
	fmt.Println(adjacent)

	// result := 0
	g.visited[coordinate{startRow, startCol}] = true
	for _, adj := range adjacent {
		g.dfs(adj[0], adj[1], count)
	}
	g.visited[coordinate{startRow, startCol}] = false

	return
}

// we can use dfs and backtracking
// we can init a matrix with 0 and mark it as 1 on the visit and then backtrack

// steps:
// 1. do dfs and stop when we find the m, n
// 2. we need to mark something to start from 0 to 1 and backtrack after that
// 3. if we dont find anything to run further - this will naturally stop
