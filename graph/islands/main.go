package main

import "fmt"

func main() {
	grid1 := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	fmt.Println(numIslands(grid1))

	grid2 := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(grid2))
}

func numIslands(grid [][]byte) int {

	graph := &gridS{
		mapping: grid,
		visited: map[string]bool{},
	}
	island := 0
	for i := range graph.mapping {
		for j := range graph.mapping[i] {
			if graph.mapping[i][j] == '0' {
				continue
			}
			sourceStr := fmt.Sprintf("%d,%d", i, j)
			if _, ok := graph.visited[sourceStr]; !ok {
				graph.dfs([]int{i, j})
				island++
			}
		}
	}
	return island
}

type gridS struct {
	mapping [][]byte
	visited map[string]bool
}

// since we only care about visited or not then no need to return the node

// for traversing - use dfs to do so
// how to determine adjacent list in the case of this grid
// given x, y coordinate adjacent list is (x+1,y), (x-1,y), (x,y+1), (x,y-1)
// but we need to check that x and y cannot go beyond 0 or grid length and width
// and we only add adjacent == 1
func (g *gridS) dfs(source []int) {
	x := source[0]
	y := source[1]
	sourceStr := fmt.Sprintf("%d,%d", source[0], source[1])
	if _, ok := g.visited[sourceStr]; ok {
		return
	}
	g.visited[sourceStr] = true
	adj := [][]int{
		[]int{x + 1, y},
		[]int{x - 1, y},
		[]int{x, y + 1},
		[]int{x, y - 1},
	}

	for _, coordinate := range adj {
		if coordinate[0] < 0 || coordinate[0] > len(g.mapping)-1 {
			continue
		}
		if coordinate[1] < 0 || coordinate[1] > len(g.mapping[0])-1 {
			continue
		}
		if g.mapping[coordinate[0]][coordinate[1]] == '0' {
			continue
		}
		sourceStr := fmt.Sprintf("%d,%d", coordinate[0], coordinate[1])
		if _, ok := g.visited[sourceStr]; ok {
			continue
		}
		g.dfs(coordinate)
	}

	// check adjacent list

}

// init visited map as well "x,y"
// for each point in the grid -> (x,y) -> check if it is visited or not
// 1. if not visited then we traverse - path length more than 1
// 	- increment the island for each of this
// 2. else already visited hence path length is zero
