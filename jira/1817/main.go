package main

import "fmt"

func main() {
	findingUsersActiveMinutes([][]int{{1, 1}, {2, 2}, {2, 3}}, 4)
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {

	a := make([]int, k)

	var dict = map[int]map[int]bool{}

	for _, log := range logs {
		// 1st el is the user
		// 2nd el is the active minute
		if _, ok := dict[log[0]]; !ok {
			dict[log[0]] = map[int]bool{}
		}
		if _, ok := dict[log[0]][log[1]]; !ok {
			dict[log[0]][log[1]] = true
		}
	}

	fmt.Println("dict: ", dict)

	for _, val := range dict {
		a[len(val)-1]++
		// whats the length of this one

	}
	fmt.Println("a: ", a)

	return a
}
