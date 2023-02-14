package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}

func groupThePeople(groupSizes []int) [][]int {

	dict := map[int][]int{}
	result := [][]int{}
	for _, size := range groupSizes {
		if _, ok := dict[size]; !ok {
			dict[size] = []int{}
		}
	}

	for i, size := range groupSizes {
		dict[size] = append(dict[size], i)
		if len(dict[size]) == size {
			result = append(result, dict[size])
			dict[size] = []int{}
		}
	}

	return result
}

// Input: groupSizes = [3,3,3,3,3,1,3]
// Output: [[5],[0,1,2],[3,4,6]]

// create map[int][]int{}

// 3 - [0,1,2] - del the key when done
// 3 - [3,4,6] - del the key when done
// 1 - [5] - del the key when done
