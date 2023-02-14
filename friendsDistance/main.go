package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(friendsDistance([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 0, 1))
	fmt.Println(friendsDistance([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 1, 2))
	fmt.Println(friendsDistance([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 0, 2))

}

func friendsDistance(arr [][]int, friend1, friend2 int) int {
	if friend1 == friend2 {
		return 0
	}

	g := graph{
		adjacent: map[int][]int{},
		visited:  map[int]bool{},
	}

	for i, f := range arr {
		g.adjacent[i] = []int{}
		for j, b := range f {
			if b == 1 {
				g.adjacent[i] = append(g.adjacent[i], j)
			}
		}
	}

	return g.bfs(friend1, friend2)
}

type graph struct {
	adjacent map[int][]int
	visited  map[int]bool
}

func (g *graph) bfs(start, end int) int {

	queue := [][]int{} // store the friend and degree
	queue = append(queue, []int{start, 0})
	g.visited[start] = true

	degree := int(math.Inf(1))

out:
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, adj := range g.adjacent[current[0]] {
			if adj == end {
				degree = current[1] + 1
				break out
			}
			if _, ok := g.visited[adj]; !ok {
				queue = append(queue, []int{adj, current[1] + 1})
			}
		}
	}

	return degree

}

// problem like this normally can be done with bst
// so in each time traversal happening, keep track the level
// start is from the friend1
// it will stop adding when finding the friend2
