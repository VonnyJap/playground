package main

import "fmt"

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	fmt.Println(jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}))
	fmt.Println(jobScheduling([]int{1, 2, 2, 3}, []int{2, 5, 3, 4}, []int{3, 4, 1, 2}))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var P = []int{}

	max := 0
	for _, end := range endTime {
		if end > max {
			max = end
		}
	}

	for i := 0; i <= max; i++ {
		P = append(P, 0)
	}

	// start from 1
	for i := 1; i < len(P); i++ {
		// set this as before in case nothing get done at this time
		P[i] = P[i-1]
		// check if anything get done by this time
		for j, end := range endTime {
			if end == i {
				// compare current with profit[j] + P[startTime] vs current P
				// lets do something here
				if profit[j]+P[startTime[j]] > P[i] {
					P[i] = profit[j] + P[startTime[j]]
				}
			}
		}
	}

	fmt.Println(P)
	return P[len(P)-1]
}

// use hash map and then also maybe we can use something else
func jobSchedulingEfficient(startTime []int, endTime []int, profit []int) int {
	var P = []int{}

	max := 0
	for _, end := range endTime {
		if end > max {
			max = end
		}
	}

	for i := 0; i <= max; i++ {
		P = append(P, 0)
	}

	// start from 1
	for i := 1; i < len(P); i++ {
		// set this as before in case nothing get done at this time
		P[i] = P[i-1]
		// check if anything get done by this time
		for j, end := range endTime {
			if end == i {
				// compare current with profit[j] + P[startTime] vs current P
				// lets do something here
				if profit[j]+P[startTime[j]] > P[i] {
					P[i] = profit[j] + P[startTime[j]]
				}
			}
		}
	}

	fmt.Println(P)
	return P[len(P)-1]
}

// this should use DP or recursion
// maybe use DP easier

// what about making jobSchedulingrecur?
