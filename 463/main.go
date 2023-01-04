package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1}, {1, 1}}))
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))

}

func islandPerimeter(grid [][]int) int {

	g := &Grid{
		content: grid,
		visited: map[string]bool{},
	}
	x := len(grid)
	y := len(grid[0])

	fmt.Println(g.content)
	perimeter := 0
out:
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if g.content[i][j] == 1 {
				fmt.Println("i,j:", i, j)
				perimeter = g.dfs(i, j)
				break out
			}
		}
	}

	return perimeter
}

// we can approach this problem using dfs and

// we build adjacency list and each time we found new cell
// add 4
// but each time there is adjacency, do --

type Grid struct {
	content [][]int
	visited map[string]bool
}

func (g *Grid) dfs(x, y int) int {

	coordinateStr := fmt.Sprintf("%d-%d", x, y)
	if _, ok := g.visited[coordinateStr]; ok {
		return 0
	}

	g.visited[coordinateStr] = true
	adjacent := [][]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}

	perimeter := 4
	for _, n := range adjacent {
		// make sure not out of the range
		if n[0] < 0 || n[0] >= len(g.content) || n[1] < 0 || n[1] >= len(g.content[0]) {
			continue
		}
		if g.content[n[0]][n[1]] == 1 {
			nStr := fmt.Sprintf("%d-%d", n[0], n[1])
			perimeter--
			if _, ok := g.visited[nStr]; !ok {
				perimeter += g.dfs(n[0], n[1])
			}
		}
	}

	return perimeter
}
