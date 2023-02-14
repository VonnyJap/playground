package main

import "fmt"

func main() {

	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))

}

func uniquePaths(m int, n int) int {
	g := Graph{
		m:       m,
		n:       n,
		visited: map[Coordinate]bool{},
	}

	return g.dfs(Coordinate{0, 0}, map[Coordinate]int{})
}

type Graph struct {
	m, n    int
	visited map[Coordinate]bool
}

type Coordinate struct {
	x, y int
}

func (g *Graph) dfs(start Coordinate, memo map[Coordinate]int) int {

	if val, ok := memo[start]; ok {
		return val
	}

	if start.x == g.m-1 && start.y == g.n-1 {
		memo[start] = 1
		return memo[start]
	}
	if start.x >= g.m || start.y >= g.n || start.x < 0 || start.y < 0 {
		memo[start] = 0
		return memo[start]
	}
	if visited, ok := g.visited[start]; visited && ok {
		memo[start] = 0
		return memo[start]
	}

	g.visited[start] = true
	total := 0
	points := []Coordinate{
		Coordinate{start.x + 1, start.y},
		Coordinate{start.x, start.y + 1},
	}

	for _, point := range points {
		total += g.dfs(point, memo)
	}
	g.visited[start] = false

	memo[start] = total
	return memo[start]
}

// we can do this recursively for example using dfs?
// do we need to backtrack? - maybe yes
