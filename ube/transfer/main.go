package main

import "fmt"

func main() {
	fmt.Println(minTransfers([][]int{{0, 1, 10}, {2, 0, 5}}))
}

// build a map so we know the end state of what each person has

func minTransfers(transactions [][]int) int {
	var dict = map[int]int{}

	for _, tr := range transactions {
		p0 := tr[0]
		p1 := tr[1]
		if _, ok := dict[p0]; !ok {
			dict[p0] = 0
		}
		dict[p0] = dict[p0] - tr[2]
		if _, ok := dict[p1]; !ok {
			dict[p1] = 0
		}
		dict[p1] = dict[p1] + tr[2]
	}

	var debt = []int{}

	// loop through the dict and add to the debt
	for _, m := range dict {
		debt = append(debt, m)
	}
	fmt.Println("debt: ", debt)
	return getMinTransfers(0, debt)
}

func getMinTransfers(id int, debt []int) int {

	for id < len(debt) && debt[id] == 0 {
		id++
	}
	if id == len(debt) {
		return 0
	}
	var min = 1000000000000000
	for i := id + 1; i < len(debt); i++ {
		if debt[i]*debt[id] < 0 {
			fmt.Println("0: ", debt[id])
			fmt.Println("1: ", debt[i])
			debt[i] += debt[id]
			minLocal := getMinTransfers(id+1, debt) + 1
			if minLocal < min {
				min = minLocal
			}
			debt[i] -= debt[id]
			fmt.Println("2: ", debt[i])
		}
	}
	return min
}
