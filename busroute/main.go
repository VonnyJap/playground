package main

import "fmt"

func main() {

	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 5, 7}, {3, 6, 8}}, 1, 6))
}

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	vertexMap, visited, queue, res := map[int][]int{}, make([]bool, len(routes)), []int{}, 0
	for i := 0; i < len(routes); i++ {
		for _, v := range routes[i] {
			tmp := vertexMap[v]
			tmp = append(tmp, i)
			vertexMap[v] = tmp
		}
	}
	fmt.Println(vertexMap)
	queue = append(queue, S)
	for len(queue) > 0 {
		res++
		qlen := len(queue)
		fmt.Println("qlen: ", qlen)
		for i := 0; i < qlen; i++ {
			vertex := queue[0]
			fmt.Println(vertex)

			queue = queue[1:]
			fmt.Println(queue)
			for _, bus := range vertexMap[vertex] {
				if visited[bus] == true {
					continue
				}
				visited[bus] = true
				for _, v := range routes[bus] {
					if v == T {
						return res
					}
					queue = append(queue, v)
				}
			}
		}
		fmt.Println("new queue: ", queue)
	}
	return -1
}

// so for the bus route we create something like a map to know which buses are intersecting
// once we know that we do bfs
// we loop through the whole queue because we consider that as one bus - if the source is having two bus routes, we consider we can chose one
