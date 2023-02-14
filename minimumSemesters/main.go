package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}, {3, 4}, {2, 5}}))
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}, {3, 4}, {4, 2}}))
	fmt.Println(minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))

}

// create prereqs
// create graph and edges
// create visiting map
// create all vertices so we can loop through
// only start dfs from those dont have prereqs
// the dfs will return hasCycle and the level?

// Input: n = 3, relations = [[1,3],[2,3]]
// Input: n = 3, relations = [[1,3],[2,3],[3,4],[3,5]]

// func minimumSemesters(n int, relations [][]int) int {

// 	graph := &Graph{
// 		Vertices:   []int{},
// 		Edges:      map[int][]int{},
// 		Visited:    map[int]bool{},
// 		VisitedLen: map[int]int{},
// 	}

// 	for i := 1; i <= n; i++ {
// 		graph.Vertices = append(graph.Vertices, i)
// 	}

// 	for _, relation := range relations {
// 		if _, ok := graph.Edges[relation[0]]; ok {
// 			graph.Edges[relation[0]] = append(graph.Edges[relation[0]], relation[1])
// 			continue
// 		}
// 		graph.Edges[relation[0]] = []int{relation[1]}
// 	}

// 	for _, v := range graph.Vertices {
// 		if graph.hasCycle(v, []int{}) {
// 			return -1
// 		}
// 	}

// 	max := 0
// 	for _, v := range graph.Vertices {
// 		depth := graph.traverse(v)
// 		if depth > max {
// 			max = depth
// 		}
// 	}

// 	return max
// }

// type Graph struct {
// 	Vertices   []int
// 	Edges      map[int][]int
// 	Visited    map[int]bool
// 	VisitedLen map[int]int
// }

// func (g *Graph) hasCycle(v int, path []int) bool {

// 	g.Visited[v] = true
// 	path = append(path, v)
// 	for _, node := range g.Edges[v] {
// 		if inSlice(path, node) {
// 			return true
// 		}
// 		if !g.Visited[node] {
// 			if g.hasCycle(node, path) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func (g *Graph) traverse(v int) int {

// 	if _, ok := g.VisitedLen[v]; ok {
// 		return g.VisitedLen[v]
// 	}
// 	if len(g.Edges[v]) == 0 {
// 		return 1
// 	}
// 	depths := []int{}
// 	for _, node := range g.Edges[v] {
// 		depths = append(depths, 1+g.traverse(node))
// 	}

// 	sort.Ints(depths)
// 	g.VisitedLen[v] = depths[len(depths)-1]
// 	return depths[len(depths)-1]
// }

// func inSlice(haystack []int, needle int) bool {
// 	for _, item := range haystack {
// 		if item == needle {
// 			return true
// 		}
// 	}
// 	return false
// }

func minimumSemesters(n int, relations [][]int) int {

	graph := &Graph{
		Vertices: []int{},
		Edges:    map[int][]int{},
		Visited:  map[int]bool{},
	}

	for i := 1; i <= n; i++ {
		graph.Vertices = append(graph.Vertices, i)
	}

	for _, relation := range relations {
		if _, ok := graph.Edges[relation[0]]; ok {
			graph.Edges[relation[0]] = append(graph.Edges[relation[0]], relation[1])
			continue
		}
		graph.Edges[relation[0]] = []int{relation[1]}
	}

	max := 0
	for _, v := range graph.Vertices {
		depth := graph.traverse(v, []int{})
		if depth == -1 {
			return -1
		}
		if depth > max {
			max = depth
		}
	}

	return max + 1
}

type Graph struct {
	Vertices []int
	Edges    map[int][]int
	Visited  map[int]bool
}

func (g *Graph) traverse(v int, path []int) int {

	// g.Visited[v] = true
	path = append(path, v)

	depths := []int{}
	for _, node := range g.Edges[v] {
		if inSlice(path, node) {
			return -1
		}
		// if !g.Visited[node] {
		depth := g.traverse(node, path)
		if depth == -1 {
			return -1
		}
		depths = append(depths, 1+depth)
		// }
	}

	if len(depths) == 0 {
		return 0
	}
	sort.Ints(depths)
	return depths[len(depths)-1]
}

func inSlice(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

// using the concept of visiting and marking as -1 when visiting
// class Solution:
//     def minimumSemesters(self, N: int, relations: List[List[int]]) -> int:
//         graph = {i: [] for i in range(1, N + 1)}
//         for start_node, end_node in relations:
//             graph[start_node].append(end_node)

//         visited = {}

//         def dfs(node: int) -> int:
//             # return the longest path (inclusive)
//             if node in visited:
//                 return visited[node]
//             else:
//                 # mark as visiting
//                 visited[node] = -1

//             max_length = 1
//             for end_node in graph[node]:
//                 length = dfs(end_node)
//                 # we meet a cycle!
//                 if length == -1:
//                     return -1
//                 else:
//                     max_length = max(length+1, max_length)
//             # mark as visited
//             visited[node] = max_length
//             return max_length

//         max_length = -1
//         for node in graph.keys():
//             length = dfs(node)
//             # we meet a cycle!
//             if length == -1:
//                 print(visited)
//                 return -1
//             else:
//                 max_length = max(length, max_length)

//         print(visited)
//         print("hello")
//         return max_length
