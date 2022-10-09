package main

import "fmt"

func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))

}

func nearestValidPoint(x int, y int, points [][]int) int {

	// first find the valid points and set aside the remaining points
	// so given those valid points, calculate the distance

	// this solution is O(N)
	var min int = 10000000
	var index int = -1
	for i, point := range points {
		if point[0] == x || point[1] == y {
			distanceX := x - point[0]
			if distanceX < 0 {
				distanceX = distanceX * (-1)
			}
			distanceY := y - point[1]
			if distanceY < 0 {
				distanceY = distanceY * (-1)
			}
			distance := distanceX + distanceY
			if distance < min {
				min = distance
				index = i
			}
		}
	}

	return index
}
