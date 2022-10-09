package main

import (
	"fmt"
	"sort"
)

type Matrix [3][3]int

func (m Matrix) Len() int { return len(m) }
func (m Matrix) Less(i, j int) bool {
	for x := range m[i] {
		if m[i][x] == m[j][x] {
			continue
		}
		return m[i][x] < m[j][x]
	}
	return false
}

func (m *Matrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func main() {
	var matrix = [3][3]int{
		{2, 3, 1},
		{6, 3, 5},
		{1, 4, 9},
	}
	m := Matrix(matrix)
	sort.Sort(&m)
	fmt.Println(m)
}
