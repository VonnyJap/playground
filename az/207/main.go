package main

func main() {}

// Input:
// numCourses(n) course to finish
// prerequisites which is [i, j] - finish i before doing j
// Output: true or false
// If no prerequisites then can be finished anytime

// 2 - [[1,0]] -> true
// 2 - [[1,0],[0,1]] -> false

// detect a cycle in the graph?

// 2: [1]
// 1: [2, 3]
// cycle between 1 and 2 because 2 is part of 1
// first I want to build the adjacency list first - map[int][]int{}

// steps:
// 1. build the map first from prerequisites - map[int]int{}
// 2. construct a loop and check for each course if there is a prerequisites
// 3. if there is a prerequisites for this course then check if there is a cycle between these two
// 	- if nothing then continue
// 	- otherwise return false immediately

// 4. create a helper function called hasCycle to check if the course is a prerequisites for each of the prerequisites as well

// func canFinish(numCourses int, prerequisites [][]int) bool {

// 	prerequisitesMap := map[int][]int{}

// 	for _, p := range prerequisites {
// 		course := p[0]
// 		pre := p[1]
// 		if _, ok := prerequisitesMap[course]; ok {
// 			prerequisitesMap[course] = append(prerequisitesMap[course], pre)
// 			continue
// 		}
// 		prerequisitesMap[course] = []int{pre}
// 	}

// 	for course, prereqs := range prerequisitesMap {
// 		for _, c := range {

// 		}
// 	}

// 	// then for each of map we want to check for cycle
// }

// can this be solved with recursion?

type Graph struct {
	nodes []int
	edges map[int][]int
}

func generateGraph(num int, edges [][]int) Graph {
	g := Graph{nodes: make([]int, 0), edges: make(map[int][]int, 0)}
	// generate node
	for i := 0; i < num; i++ {
		node := i
		g.nodes = append(g.nodes, node)
	}
	// generate graph direct edges
	for i := 0; i < len(edges); i++ {
		sNode := edges[i][0]
		eNode := edges[i][1]
		g.edges[sNode] = append(g.edges[sNode], eNode)
	}
	return g
}

func checkCycle(g Graph, num int) bool {

	/*
		visited means those nodes have been visited
		visiting  means those nodes are visiting in this traverse
	*/
	visited := make(map[int]bool, 0)
	visiting := make(map[int]bool, 0)
	for i := 0; i < num; i++ {
		if dfs(g, i, visited, visiting) == true {
			return false
		}
	}
	return true
}

func dfs(g Graph, i int, visited map[int]bool, visiting map[int]bool) bool {
	if visited[i] != true {
		visited[i] = true
		visiting[i] = true
		for j := 0; j < len(g.edges[i]); j++ {
			neighbour := g.edges[i][j]
			if visited[neighbour] != true && dfs(g, neighbour, visited, visiting) {
				return true
			} else if visiting[neighbour] == true {
				return true
			}
		}
	}
	visiting[i] = false
	return false
}
