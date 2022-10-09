package main

import "fmt"

func main() {

	g := &Graph{Nodes: map[string]*Node{}}

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")
	d := NewNode("D")
	e := NewNode("E")
	f := NewNode("F")

	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddNode(d)
	g.AddNode(e)
	g.AddNode(f)

	g.AddEdge(a, b, 1)
	g.AddEdge(b, a, 1)
	g.AddEdge(a, c, 1)
	g.AddEdge(c, a, 1)
	g.AddEdge(a, d, 1)
	g.AddEdge(d, a, 1)
	g.AddEdge(b, e, 1)
	g.AddEdge(e, b, 1)

	fmt.Println(g.DFS(a))
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

// how do I perform DFS and BFS now?

// for BFS
// starting node is marked as Visited
// and then add the starting node to the queue
// when queue not empty
// check adjacency and if not visited: add the node to queue and mark as visited
// if not white - dont explore
// when things are all explored then set it to black
