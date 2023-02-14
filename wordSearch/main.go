package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, "ABCCED"))
	fmt.Println(exist([][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, "SEE"))
	fmt.Println(exist([][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, "ABCB"))
	fmt.Println(exist([][]byte{[]byte("a")}, "a"))
	fmt.Println(exist([][]byte{[]byte("a"), []byte("a")}, "a"))
	fmt.Println(exist([][]byte{[]byte("CAA"), []byte("AAA"), []byte("BCD")}, "AAB"))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		if word == "" {
			return true
		}
		return false
	}

	row := len(board)
	col := len(board[0])

	for i, r := range board {
		for j, c := range r {
			if c == word[0] {
				grh := graph{
					matrix: board,
					row:    row,
					col:    col,
					target: word,
				}
				if grh.dfs(i, j, word) {
					return true
				}
			}
		}
	}
	return false
}

type graph struct {
	matrix [][]byte
	target string
	row    int
	col    int
}

func (g *graph) dfs(x, y int, target string) bool {
	if len(target) == 0 {
		return true
	}
	if x < 0 || x >= g.row || y < 0 || y >= g.col || g.matrix[x][y] != target[0] {
		return false
	}

	g.matrix[x][y] = '#'

	adjacent := [][]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}

	ret := true
	for _, adj := range adjacent {
		ret = g.dfs(adj[0], adj[1], target[1:])
		if ret {
			break
		}
	}

	g.matrix[x][y] = target[0]

	return ret
}

// another idea here is actually to use dfs

// steps:
// 1. in exist loop through the matrix and select only the cell that is equal to the first letter of the word
// 2. in the recursive
// - if target == "" => true
// - if target OOB or not same with first letter => false
// - mark the cell as something "-"
// - loop to all boundaries and remove first letter from the target
// - do cleanup
// - return true/false
