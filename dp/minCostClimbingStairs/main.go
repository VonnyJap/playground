package main

import "fmt"

func main() {
	// fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))

}
func minCostClimbingStairs(cost []int) int {

	min := cost[0]
	start := 0

	if cost[1] <= min {
		min = cost[1]
		start = 1
	}

	for {
		fmt.Println(start)
		if start+1 >= len(cost) {
			fmt.Println("break")
			break
		}
		if cost[start+2] <= cost[start+1] {
			start += 2
			min += cost[start+2]
			continue
		}
		start += 1
		min += cost[start+1]
	}

	return min
}

// [1,100,1,1,1,100,1,1,100,1]
// 0-2-4-6-7-9
// cost = [10,15,20]

// when we start => compare whether we will go to 0 or 1
// thats going to be our start
