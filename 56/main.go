package main

import (
	"fmt"
	"sort"
)

// step
// 1. sort array
// 2. check if the new start <= currentEnd
// 3. update currentend accordingly
// 4. if no interval - merge the current element to the result and update currentEnd and start

func merge(intervals [][]int) [][]int {

	// lets sort by start time
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	// get current start and end from the first in the sorted intervals
	currentStart := intervals[0][0]
	currentEnd := intervals[0][1]
	result := [][]int{}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= currentEnd {
			if intervals[i][1] > currentEnd {
				currentEnd = intervals[i][1]
			}
		} else {
			result = append(result, []int{currentStart, currentEnd})
			currentStart = intervals[i][0]
			currentEnd = intervals[i][1]
		}
	}
	result = append(result, []int{currentStart, currentEnd})
	return result
}

func main() {
	// [1,3],[2,6]

	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))

}
