package main

import "fmt"

func main() {
	fmt.Println(scs("ABCBDAB", "BDCABA"))
	fmt.Println(scsDP("AB", "CB"))
	fmt.Println(scsDP("geek", "eke"))
	fmt.Println(scsDP("AGGTAB", "GXTXAYB"))
	fmt.Println(scsDP("ABCBDAB", "BDCABA"))
}

// going to think about it recursive first and then we will go from there
// i know it is not optimized with the recursive but i will slowly think about next
func scs(A, B string) int {
	if len(A) == 0 || len(B) == 0 {
		return len(A) + len(B)
	}
	if A[len(A)-1] == B[len(B)-1] {
		return scs(A[0:len(A)-1], B[0:len(B)-1]) + 1
	}
	first := scs(A[0:len(A)-1], B) + 1
	second := scs(A, B[0:len(B)-1]) + 1
	if first < second {
		return first
	}
	return second
}

func scsDP(A, B string) int {

	var matrix [][]int
	for i := 0; i <= len(A); i++ {
		m := []int{}
		for j := 0; j <= len(B); j++ {
			m = append(m, 0)
		}
		matrix = append(matrix, m)
	}
	for i := 0; i <= len(A); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(B); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if string(A[i-1]) == string(B[j-1]) {
				matrix[i][j] = matrix[i-1][j-1] + 1
				continue
			}
			if matrix[i-1][j] < matrix[i][j-1] {
				matrix[i][j] = matrix[i-1][j] + 1
				continue
			}
			matrix[i][j] = matrix[i][j-1] + 1
		}
	}

	return matrix[len(A)][len(B)]
}
