package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices int
	Adjacent [][]int
}

func Construct(v int) *Graph {
	var graph = &Graph{}
	graph.Vertices = v
	graph.Adjacent = make([][]int, v)
	for i := range graph.Adjacent {
		graph.Adjacent[i] = make([]int, v)
	}
	fmt.Println(graph.Adjacent)
	return graph
}

// this is literally the distance from this start
func (g *Graph) dijkstra(start int) {
	dist := []int{}
	for i := 0; i < g.Vertices; i++ {
		dist = append(dist, int(math.Inf(1)))
	}
	dist[start] = 0
	// shortest path set
	sptSet := []bool{}
	for i := 0; i < g.Vertices; i++ {
		sptSet = append(sptSet, false)
	}

	for i := 0; i < g.Vertices; i++ {
		index := g.findMinDistance(dist, sptSet)
		sptSet[index] = true
		for j := 0; j < g.Vertices; j++ {
			if g.Adjacent[i][j] > 0 && sptSet[j] == false && dist[j] > dist[i]+g.Adjacent[i][j] {
				dist[j] = dist[i] + g.Adjacent[i][j]
			}
		}
	}
	fmt.Println("distance from src ", start, dist)
}

func (g *Graph) findMinDistance(dist []int, sptSet []bool) (index int) {

	min := int(math.Inf(1))

	for i := 0; i < g.Vertices; i++ {
		if dist[i] < min && sptSet[i] == false {
			min = dist[i]
			index = i
		}
	}
	return
}

func main() {
	graph := Construct(9)
	graph.Adjacent = [][]int{
		[]int{0, 4, 0, 0, 0, 0, 0, 8, 0},
		[]int{4, 0, 8, 0, 0, 0, 0, 11, 0},
		[]int{0, 8, 0, 7, 0, 4, 0, 0, 2},
		[]int{0, 0, 7, 0, 9, 14, 0, 0, 0},
		[]int{0, 0, 0, 9, 0, 10, 0, 0, 0},
		[]int{0, 0, 4, 14, 10, 0, 2, 0, 0},
		[]int{0, 0, 0, 0, 0, 2, 0, 1, 6},
		[]int{8, 11, 0, 0, 0, 0, 1, 0, 7},
		[]int{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}
}
