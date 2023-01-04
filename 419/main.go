package main

import "fmt"

func main() {
	fmt.Println(countBattleships([][]byte{
		[]byte("X..X"),
		[]byte("...X"),
		[]byte("...X"),
		[]byte("...."),
	}))
	fmt.Println(countBattleships([][]byte{
		[]byte("."),
	}))
}

func countBattleships(board [][]byte) int {

	battle := 0
	x := len(board)
	y := len(board[0])

	grid := &grid{
		content: board,
		visited: map[string]bool{},
	}

	// looking for "X" in the board
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if string(grid.content[i][j]) == "X" {
				coordinateStr := fmt.Sprintf("%d-%d", i, j)
				if _, ok := grid.visited[coordinateStr]; !ok {
					grid.dfs(i, j)
					battle++
				}
			}
		}
	}

	return battle
}

type grid struct {
	content [][]byte
	visited map[string]bool
}

func (g *grid) dfs(x, y int) {

	coordinateStr := fmt.Sprintf("%d-%d", x, y)
	if _, ok := g.visited[coordinateStr]; ok {
		return
	}

	g.visited[coordinateStr] = true
	adjacent := [][]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}

	for _, n := range adjacent {
		// make sure not out of the range
		if n[0] < 0 || n[0] >= len(g.content) || n[1] < 0 || n[1] >= len(g.content[0]) {
			continue
		}
		if string(g.content[n[0]][n[1]]) == "X" {
			nStr := fmt.Sprintf("%d-%d", n[0], n[1])
			if _, ok := g.visited[nStr]; !ok {
				g.dfs(n[0], n[1])
			}
		}
	}

}

// basically this is a problem of dfs
// how to find where to start?
// we can loop through everything and create adjacent list

// for this problem we need to get all the "X" in the grid
// also we need to introduce adjacent list
// create a function to dfs:
// 1. before dfs check if it has been visited before
// 2. if visited then just stop
// 3. create adjacent list and check for each if visited before or not
