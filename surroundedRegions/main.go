package main

import "fmt"

func main() {
	board := [][]byte{[]byte("XXXX"), []byte("XOOX"), []byte("XXOX"), []byte("XOXX")}
	solve(board)
	fmt.Println(board)

	board1 := [][]byte{[]byte("OXXOX"), []byte("XOOXO"), []byte("XOXOX"), []byte("OXOOO"), []byte("XXOXO")}
	solve(board1)
	fmt.Println(board1)
}

func solve(board [][]byte) {

	if len(board) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])

	visited := map[int]map[int]bool{}
	forbidden := map[int]map[int]bool{}

	for i := 0; i < m; i++ {
		visited[i] = map[int]bool{}
		forbidden[i] = map[int]bool{}
	}

	queue := [][]int{{0, 0}}
	visited[0][0] = true

	// we need to go through the border first and then doing the rest

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current[0]][current[1]] = true

		adjacent := [][]int{}
		if current[0]+1 < m {
			adjacent = append(adjacent, []int{current[0] + 1, current[1]})
		}
		if current[0]-1 >= 0 {
			adjacent = append(adjacent, []int{current[0] - 1, current[1]})
		}
		if current[1]+1 < n {
			adjacent = append(adjacent, []int{current[0], current[1] + 1})
		}
		if current[1]-1 >= 0 {
			adjacent = append(adjacent, []int{current[0], current[1] - 1})
		}

		adjacentForbidden := false

		for _, adj := range adjacent {
			if val, ok := forbidden[adj[0]][adj[1]]; ok && val {
				adjacentForbidden = true
			}
			if val, ok := visited[adj[0]][adj[1]]; ok && val {
				continue
			}
			queue = append(queue, adj)
		}

		if board[current[0]][current[1]] == 'O' {
			if current[0] == 0 || current[0] == m-1 || current[1] == 0 || current[1] == n-1 {
				forbidden[current[0]][current[1]] = true
				continue
			}
			if adjacentForbidden {
				continue
			}
			board[current[0]][current[1]] = 'X'
		}
	}
}

// do bfs as so

// 1. init a visited map
// 2. perform bfs from (0,0)
// 3. adjacent list - x+1, x-1, y+1, y-1
// 4. x >= 0 and y>=0 and x<len(board) and y<len(board[0])
// 5. check if the value is "0"
// - check if at the border - if yes then continue
// - check if len(x)-2 or len(y)-2 => check that the next is not "0"
