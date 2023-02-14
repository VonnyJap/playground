package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {

	transpose(matrix)

	reverse(matrix)
}

func transpose(matrix [][]int) {

	rowStart := 0
	rowEnd := len(matrix)

	if rowEnd == 0 {
		return
	}
	colEnd := len(matrix[0])

	for i := rowStart; i < rowEnd; i++ {
		for j := i + 1; j < colEnd; j++ {
			temp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}

}

func reverse(matrix [][]int) {
	rowStart := 0
	rowEnd := len(matrix)

	if rowEnd == 0 {
		return
	}

	for i := rowStart; i < rowEnd; i++ {
		colStart := 0
		colEnd := rowEnd - 1
		for colStart < colEnd {
			matrix[i][colStart], matrix[i][colEnd] = matrix[i][colEnd], matrix[i][colStart]
			colStart++
			colEnd--
		}
	}
}
