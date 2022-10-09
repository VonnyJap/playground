package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var i = numRows
	var a [][]int

	for i > 0 {
		// row 0
		if numRows-i == 0 {
			a = append(a, []int{1})
			i--
			continue
		}
		// row 1
		if numRows-i == 1 {
			a = append(a, []int{1, 1})
			i--
			continue
		}
		var b = []int{}
		b = append(b, 1)
		fmt.Println(numRows - i - 1)
		for j := 1; j < numRows-i; j++ {
			b = append(b, a[numRows-i-1][j]+a[numRows-i-1][j-1])
		}
		b = append(b, 1)
		a = append(a, b)
		i--
	}

	return a
}
