package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {

	// sort based on the maximum that can be fit and then each time check with truckSize
	// golang has this kind of sort

	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[j][1] < boxTypes[i][1] })
	// fmt.Println("boxTypes:", boxTypes)

	remainingLoad := truckSize
	totalUnits := 0

	for _, box := range boxTypes {
		if remainingLoad < box[0] {
			totalUnits += remainingLoad * box[1]
			break
		}
		remainingLoad -= box[0]
		totalUnits += box[0] * box[1]
	}

	return totalUnits
}
