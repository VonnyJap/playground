package main

import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
	fmt.Println(maximumRows(matrix, 2))

	matrix1 := [][]int{{0}}
	fmt.Println(maximumRows(matrix1, 1))
}

func maximumRows(matrix [][]int, numSelect int) int {

	countOne := map[int]int{}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				if _, ok := countOne[col]; !ok {
					countOne[col] = 1
					continue
				}
				countOne[col]++
			}
		}
	}

	countOneIter := []CountOne{}
	for col, count := range countOne {
		countOneIter = append(countOneIter, CountOne{col, count})
	}

	sort.Slice(countOneIter, func(i, j int) bool { return countOneIter[i].Count > countOneIter[j].Count })

	if len(countOneIter) > 0 && len(countOneIter) <= numSelect {
		countOneIter = countOneIter[0:numSelect]
	}
	countOneFinal := map[int]bool{}

	for _, countOne := range countOneIter {
		countOneFinal[countOne.Col] = true
	}

	countRow := 0
	for row := range matrix {
		good := true
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				if _, ok := countOneFinal[col]; !ok {
					good = false
					break
				}
			}
		}
		if good {
			countRow++
		}
	}

	return countRow
}

type CountOne struct {
	Col   int
	Count int
}

// brute force method
// init a DS with Col and Count
// and then we will select those cols
// and after we loop again and count how many rows?
