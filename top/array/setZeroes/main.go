package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix1)
	fmt.Println(matrix1)

	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix2)
	fmt.Println(matrix2)
}

func setZeroes(matrix [][]int) {

	rows := []int{}
	cols := []int{}

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	// rows and cols will be sorted naturally

	for i, r := range rows {
		if i == 0 || rows[i] != rows[i-1] {
			for j := 0; j < n; j++ {
				matrix[r][j] = 0
			}
		}
	}

	for j, c := range cols {
		if j == 0 || cols[j] != cols[j-1] {
			for i := 0; i < m; i++ {
				matrix[i][c] = 0
			}
		}
	}

}

// arn:aws:iam::330848248187:role/home.vonny.aws.istio-ingressgateway
