package main

import "fmt"

func main() {

	fmt.Println(maxProfitAssignment([]int{64, 88, 97}, []int{53, 86, 89}, []int{98, 11, 6}))

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	var T []int
	for i := 0; i <= len(worker); i++ {
		T = append(T, 0)
	}

	for i, c := range worker {
		// at most no profit added by this worker
		T[i+1] = T[i+1-1]
		for j, d := range difficulty {
			if d <= c {
				if T[i+1] < T[i+1-1]+profit[j] {
					T[i+1] = T[i+1-1] + profit[j]
				}
			}
		}
	}
	return T[len(worker)]
}
