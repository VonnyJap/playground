package main

import "fmt"

func main() {
	fmt.Println(degreesFriendship([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 0, 1))
	fmt.Println(degreesFriendship([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 1, 2))
	fmt.Println(degreesFriendship([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 0, 2))
	fmt.Println(degreesFriendship([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 0, 0))
	fmt.Println(degreesFriendship([][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}, 1, 2))

}

func degreesFriendship(friends [][]int, userA, userB int) int {

	if userA == userB {
		return 0
	}

	g := &graph{
		adj:     map[int][]int{},
		visited: map[int]bool{},
	}
	for i, f := range friends {
		g.adj[i] = []int{}
		for j, b := range f {
			if b == 1 {
				g.adj[i] = append(g.adj[i], j)
			}
		}
	}

	return g.bfs(userA, userB)
}

type graph struct {
	adj     map[int][]int
	visited map[int]bool
}

func (g *graph) bfs(userA int, userB int) int {

	queue := [][]int{}
	queue = append(queue, []int{userA, 0})
	g.visited[userA] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current[0] == userB {
			return current[1]
		}
		for _, friend := range g.adj[current[0]] {
			if _, ok := g.visited[friend]; ok {
				continue
			}
			queue = append(queue, []int{friend, current[1] + 1})
		}
	}

	return -1
}
