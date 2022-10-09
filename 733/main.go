package main

import "fmt"

func main() {
	image1 := [][]int{{1, 1, 1}, {1, 1, 0}}
	image1 = floodFill(image1, 1, 1, 2)
	fmt.Println(image1)

	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	image2 = floodFill(image2, 0, 0, 0)
	fmt.Println(image2)

}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	var x = len(image)
	var y = len(image[0])
	var queue = [][]int{[]int{sr, sc}}
	var visited = map[string]bool{
		fmt.Sprintf("%d,%d", sr, sc): true,
	}

	var oldColor = image[sr][sc]
	if oldColor == color {
		return image
	}
	image[sr][sc] = color
	for len(queue) > 0 {
		current := queue[0]
		currentX := current[0]
		currentY := current[1]
		queue = queue[1:]
		// need to build adjacent lists
		adjacent := [][]int{}
		// case 1 - sr + 1
		if currentX+1 < x {
			adjacent = append(adjacent, []int{currentX + 1, currentY})
		}
		// case 2 - sr - 1
		if currentX-1 >= 0 {
			adjacent = append(adjacent, []int{currentX - 1, currentY})
		}
		// case 3 - sc + 1
		if currentY+1 < y {
			adjacent = append(adjacent, []int{currentX, currentY + 1})
		}
		// case 4 - sc - 1
		if currentY-1 >= 0 {
			adjacent = append(adjacent, []int{currentX, currentY - 1})
		}
		for _, adj := range adjacent {
			if image[adj[0]][adj[1]] == oldColor {
				if _, ok := visited[fmt.Sprintf("%d,%d", adj[0], adj[1])]; !ok {
					visited[fmt.Sprintf("%d,%d", adj[0], adj[1])] = true
					image[adj[0]][adj[1]] = color
					queue = append(queue, []int{adj[0], adj[1]})
				}
			}
		}
	}
	return image
}

// to approach this problem is like approaching bfs problem
// to start with add the sr+sc into the queue
// loop through the queue as long as it is not empty
// at the same time create a list of visited block
// we only visit when the initial color is the same with sr+sc
// once done we return the image
