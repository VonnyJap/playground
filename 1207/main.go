package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
}
func uniqueOccurrences(arr []int) bool {

	var dict1 = map[int]int{}
	var dict2 = map[int]bool{}

	for _, a := range arr {
		if _, ok := dict1[a]; ok {
			dict1[a]++
			continue
		}
		dict1[a] = 1
	}

	for _, val := range dict1 {
		if _, ok := dict2[val]; ok {
			return false
		}
		dict2[val] = true
	}

	return true
}
