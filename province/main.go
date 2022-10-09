package main

import "fmt"

func main() {

	matrix := [][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	}

	graph := NewGraph()
	for i, v := range matrix {
		graph.AddNode(NewNode(fmt.Sprintf("%d", i)))
		for j, c := range v {
			if (i != j) && (c == 1) {
				graph.AddEdge(NewNode(fmt.Sprintf("%d", i)), NewNode(fmt.Sprintf("%d", j)), 1)
			}
		}
	}

	count := 0
	for _, n := range graph.Nodes {
		if n.State == Unvisited {
			graph.DFS(*n)
			count++
		}
	}

	fmt.Println("Total group: ", count)
}

type State int64

const (
	Unvisited State = 0
	Visited   State = 1
	Visiting  State = 2
)

func (s State) String() string {
	switch s {
	case Visited:
		return "black"
	case Visiting:
		return "gray"
	}
	return "white"
}

type Node struct {
	ID       string
	State    State
	Adjacent map[string]int // key: node ID, value: weight
}

func NewNode(id string) Node {
	return Node{
		ID:       id,
		State:    Unvisited,
		Adjacent: map[string]int{},
	}
}

func (n Node) String() string {
	return n.ID
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{Nodes: map[string]*Node{}}
}

func (g *Graph) AddNode(n Node) {
	g.Nodes[n.ID] = &n
}

func (g *Graph) AddEdge(src Node, dest Node, weight int) {
	if _, ok := g.Nodes[src.ID]; !ok {
		g.Nodes[src.ID] = &src
	}
	if _, ok := g.Nodes[dest.ID]; !ok {
		g.Nodes[dest.ID] = &dest
	}
	g.Nodes[src.ID].Adjacent[dest.ID] = weight
}

func (g *Graph) BFS(start Node) []string {
	var bfs []string
	var queue []Node
	g.Nodes[start.ID].State = Visited
	queue = append(queue, start)

	for len(queue) > 0 {
		var current Node
		current, queue = queue[0], queue[1:]
		bfs = append(bfs, current.ID)
		for id := range current.Adjacent {
			node := g.Nodes[id]
			if node.State == Unvisited {
				node.State = Visiting
				queue = append(queue, *node)
			}
		}
		current.State = Visited
	}
	return bfs
}

func (g *Graph) DFS(start Node) []string {
	var dfs []string
	dfs = append(dfs, start.ID)
	g.Nodes[start.ID].State = Visited
	for id := range start.Adjacent {
		node := g.Nodes[id]
		if node.State == Unvisited {
			dfs = append(dfs, g.DFS(*node)...)
		}
	}
	return dfs
}
