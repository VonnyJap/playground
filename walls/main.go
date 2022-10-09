package main

func main() {

}

const GATE = 0

const INF = 2147483647

var DIRECTIONS = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func wallsAndGates(rooms [][]int) {

	var m = len(rooms)
	if m == 0 {
		return
	}
	var n = len(rooms[0])
	var queue [][]int

	for row := range rooms {
		for col := range rooms[row] {
			if rooms[row][col] == GATE {
				queue = append(queue, []int{row, col})
			}
		}
	}

	// we are given multi array
	for len(queue) > 0 {
		var current []int
		current, queue = queue[0], queue[1:]
		row := current[0]
		col := current[1]
		// this is the adjacent
		for _, dir := range DIRECTIONS {
			r := row + dir[0]
			c := col + dir[1]
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != INF {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			queue = append(queue, []int{r, c})
		}
	}
}
