package main

import "fmt"

func main() {

	fmt.Println(SpiralCopy([][]int{{7, 8}}))
	fmt.Println(SpiralCopy([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}))
	fmt.Println(SpiralCopy([][]int{{7, 8, 9}, {10, 11, 12}, {13, 14, 15}}))
	fmt.Println(SpiralCopy([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}))

}

// 7 8
// 9 10

// rowStart=0, rowEnd=1, colStart=0, colEnd=1
// result = []int{7,8} - rowStart=1

func SpiralCopy(inputMatrix [][]int) []int {

	result := []int{}

	rowStart := 0
	rowEnd := len(inputMatrix) - 1

	colStart := 0
	colEnd := len(inputMatrix[0]) - 1

	for rowStart <= rowEnd && colStart <= colEnd {

		for i := colStart; i <= colEnd; i++ {
			result = append(result, inputMatrix[rowStart][i])
		}
		rowStart++

		// from last col
		for i := rowStart; i <= rowEnd; i++ {
			result = append(result, inputMatrix[i][colEnd])
		}
		colEnd--

		if rowStart <= rowEnd {
			for i := colEnd; i >= colStart; i-- {
				result = append(result, inputMatrix[rowEnd][i])
			}
			rowEnd--
		}

		// from first col
		if colStart <= colEnd {
			for i := rowEnd; i >= rowStart; i-- {
				result = append(result, inputMatrix[i][colStart])
			}
			colStart++
		}
	}

	return result
}

func reverse(arr []int) []int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
