package main

import "fmt"

func main() {
	fmt.Println(getRow(4))
}

func getRow(rowIndex int) []int {

	var result = make([][]int, rowIndex+1)

	result[0] = []int{1}
	if rowIndex == 0 {
		return result[rowIndex]
	}

	result[1] = []int{1, 1}
	if rowIndex == 1 {
		return result[rowIndex]
	}

	for i := 2; i <= rowIndex; i++ {
		result[i] = []int{}
		result[i] = append(result[i], 1)
		for j := 1; j < len(result[i-1]); j++ {
			result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
		}
		result[i] = append(result[i], 1)
	}

	return result[rowIndex]
}
