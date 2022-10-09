package main

import "fmt"

var movement = map[string][]int{
	"L": []int{0, -1},
	"R": []int{0, 1},
	"U": []int{-1, 0},
	"D": []int{1, 0},
}

type Object string

const (
	Food  Object = "F"
	Empty Object = " "
	Snake Object = "S"
)

type SnakeGame struct {
	matrix [][]string
	snake  [][]int
	food   [][]int // like a queue
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	var matrix [][]string
	matrix = make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
	}
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = string(Empty)
		}
	}
	// init first food
	matrix[food[0][0]][food[0][1]] = string(Food)
	// init the head of snake
	matrix[0][0] = string(Snake)
	fmt.Println(matrix)
	return SnakeGame{
		matrix: matrix,
		snake:  [][]int{{0, 0}},
		food:   food[1:],
	}
}

// length
func (this *SnakeGame) Move(direction string) int {
	// move the snake head basically from something to another thing and when hit the wall then -1

	var value = 0
	fmt.Println("snake: ", this.snake)
	for i, pos := range this.snake {
		x := pos[0] + movement[direction][0]
		y := pos[1] + movement[direction][1]
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		if x < 0 || x >= len(this.matrix) || y < 0 || y >= len(this.matrix[0]) || this.matrix[x][y] == string(Snake) {
			value = -1
			break
		}
		if this.matrix[x][y] == string(Food) {
			value = 1
		}
		// keep this for visualization for now
		this.matrix[pos[0]][pos[1]] = string(Empty)
		this.matrix[x][y] = string(Snake)
		this.snake[i] = []int{x, y}
		// if the snake eats something that it will grow
	}

	// also if head and tail bla - do I have to check empty thing or anything?

	// when value is 1 then take out more foods and increase the snake
	if value == 1 {
		// need to do something here
		if len(this.food) > 0 {
			new := this.food[0]
			this.matrix[new[0]][new[1]] = string(Food)
		}
		if len(this.food) > 1 {
			this.food = this.food[1:]
		} else {
			this.food = [][]int{}
		}
		tail := this.snake[len(this.snake)-1]
		for _, c := range movement {
			x := tail[0] + c[0]
			y := tail[1] + c[1]
			if x > 0 && x < len(this.matrix) && y > 0 && y < len(this.matrix[0]) && this.matrix[x][y] == string(Empty) {
				this.snake = append(this.snake, []int{y, y})
				this.matrix[x][y] = string(Snake)
				break
			}
		}
	}

	fmt.Println(this.matrix)
	return value
}

func main() {
	sn := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	fmt.Println((&sn).Move("R"))
	fmt.Println((&sn).Move("D"))
	fmt.Println((&sn).Move("R"))
	fmt.Println((&sn).Move("U"))

}

// Move with no food - remove tail, add new head
// Move with food - tail the same, add new head
// tail cannot be bitten in move with food or move with no food
// we can check if new head already in the queue using hash
