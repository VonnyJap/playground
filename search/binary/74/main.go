package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

// linear + binary
// get x and y

// this can be treated as very long array of m x n
// and then with that binary search can be performed
// though there is a translation between mxn to the row,col
func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	start := 0
	end := m*n - 1

	for start <= end {
		middle := (start + end) / 2
		row := middle / n
		col := middle % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			end = middle - 1
			continue
		}
		start = middle + 1
	}
	return false
}

// 0 1 2 3 		4 5 6 7				8 9 10 11
// matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
// m x n - 3 x 4 = 12
// 0 - 0,0 -
// 1 - 0,1
// 2 - 0,2
// 3 - 0,3
// 4 - 1,0 - 4/n, 4%n
// 5 - 1,1 - 5/n, 5%n
// 6 - 1,2 - 6/n, 6%n
// 7 - 1,3 - 7/n, 7%n

// using binary search to find number in between diagonal?
// and then thats going to be the row?
// and then binary search?
