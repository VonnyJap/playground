package main

import "fmt"

func main() {

	// fmt.Println(findSingleIntersection([]int{1, 3}, []int{2, 4}))
	// fmt.Println(findSingleIntersection([]int{1, 3}, []int{4, 5}))
	// fmt.Println(findSingleIntersection([]int{1, 5}, []int{2, 3}))
	// fmt.Println(findSingleIntersection([]int{1, 5}, []int{1, 3}))
	// fmt.Println(findSingleIntersection([]int{1, 5}, []int{3, 5}))
	// fmt.Println(findSingleIntersection([]int{1, 5}, []int{1, 5}))
	// fmt.Println(findSingleIntersection([]int{1, 3}, []int{3, 5}))

	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	result := [][]int{}

	if len(firstList) == 0 || len(secondList) == 0 {
		return result
	}
	ptr1 := 0
	ptr2 := 0
	for {
		if ptr1 >= len(firstList) || ptr2 >= len(secondList) {
			break
		}
		intersection := findSingleIntersection(firstList[ptr1], secondList[ptr2])
		if len(intersection) != 0 {
			result = append(result, intersection)
		}
		if firstList[ptr1][1] == secondList[ptr2][1] {
			ptr1++
			ptr2++
			continue
		}
		if firstList[ptr1][1] < secondList[ptr2][1] {
			ptr1++
			continue
		}
		ptr2++
	}

	return result
}

func findSingleIntersection(firstList []int, secondList []int) []int {
	// when there is no overlap then we will return empty
	if (secondList[0] > firstList[0] && secondList[0] > firstList[1]) || (firstList[0] > secondList[0] && firstList[0] > secondList[1]) {
		return []int{}
	}
	result := []int{}
	if secondList[0] > firstList[0] {
		result = append(result, secondList[0])
	} else {
		result = append(result, firstList[0])
	}
	if secondList[1] < firstList[1] {
		result = append(result, secondList[1])
	} else {
		result = append(result, firstList[1])
	}
	return result
}

// ptr1 pointing to firstList
// ptr2 pointing to secondList
// 1. case of not overlap
// 	- as long as secondList[0] > firstList[0] && secondList[1] > firstList[1] || firstList[0] > secondList[0] && firstList[1] > secondList[1]
//  - increment both pointer
// 2. case of an overlap
// 	- increment whichever pointer with smaller end?
// - if same, increment both
