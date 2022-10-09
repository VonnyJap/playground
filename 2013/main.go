package main

import "fmt"

func main() {

	ds := Constructor()
	(&ds).Add([]int{3, 10})
	(&ds).Add([]int{11, 2})
	(&ds).Add([]int{3, 2})
	fmt.Println((&ds).Count([]int{11, 10}))
	fmt.Println((&ds).Count([]int{14, 8}))
	(&ds).Add([]int{11, 2})
	fmt.Println((&ds).Count([]int{11, 10}))
}

type DetectSquares struct {
	Points [][]int
}

func Constructor() DetectSquares {
	return DetectSquares{[][]int{}}
}

func (this *DetectSquares) Add(point []int) {
	this.Points = append(this.Points, point)
}

func (this *DetectSquares) Count(point []int) int {
	// linear search probably?
	xgroup := [][]int{}
	ygroup := [][]int{}

	for _, p := range this.Points {
		if p[0] == point[0] {
			xgroup = append(xgroup, p)
		}
		if p[1] == point[1] {
			ygroup = append(ygroup, p)
		}
	}

	if len(xgroup) == 0 || len(ygroup) == 0 {
		return 0
	}
	// fmt.Println(xgroup)
	// fmt.Println(ygroup)

	last := [][]int{}
	for _, x := range xgroup {
		for _, y := range ygroup {
			last = append(last, []int{y[0], x[1]})
		}
	}

	var count = 0

	for _, l := range last {
		for _, p := range this.Points {
			if l[0] == p[0] && l[1] == p[1] {
				count++
			}
		}
	}
	return count
}
