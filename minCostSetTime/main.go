package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostSetTime(1, 2, 1, 600))
	fmt.Println(minCostSetTime(0, 1, 2, 76))
	fmt.Println(minCostSetTime(0, 1, 4, 9))

}

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {

	minutes := targetSeconds / 60
	seconds := targetSeconds % 60

	cost1 := getCost(startAt, moveCost, pushCost, minutes, seconds)
	cost2 := getCost(startAt, moveCost, pushCost, minutes-1, seconds+60)

	if cost1 > cost2 {
		return cost2
	}

	return cost1
}

func getCost(startAt int, moveCost int, pushCost int, minutes, seconds int) int {

	cost := 0
	if minutes > 99 || seconds > 99 || minutes < 0 {
		return 10000000000
	}

	list := []int{}
	if minutes < 10 {
		if minutes > 0 {
			list = append(list, []int{minutes}...)
		}
	} else {
		list = append(list, []int{minutes / 10, minutes % 10}...)
	}

	if seconds < 10 {
		if len(list) > 0 {
			list = append(list, []int{0, seconds}...)
		} else {
			list = append(list, []int{seconds}...)
		}
	} else {
		list = append(list, []int{seconds / 10, seconds % 10}...)
	}

	for _, num := range list {
		if num != startAt {
			cost += moveCost
		}
		cost += pushCost
		startAt = num
	}

	return cost
}
