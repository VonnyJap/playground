package main

import (
	"fmt"
)

func main() {

	fmt.Println(minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 8))
	fmt.Println(minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 12))

}

// for this one we are going to find the slots that we can merge
// if there is something that we can merge and its duration satisfied then return it
// whatever get merged needs to be within intersecting
// actually we are finding intersection between the two

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {

	intersection := intersect(slots1, slots2)

	for _, interval := range intersection {
		cur := interval[1] - interval[0]
		if cur == duration {
			return interval
		}
		if cur > duration {
			return []int{interval[0], interval[0] + duration}
		}
	}

	return []int{}
}

// maybe we can use merge interval idea to get the intersection
// and we ignore the last add to result if the currentEnd and currentStart equals to the last thing

func intersect(slots1 [][]int, slots2 [][]int) [][]int {

	p1 := 0
	p2 := 0
	result := [][]int{}

	for p1 < len(slots1) && p2 < len(slots2) {
		l := max([]int{slots1[p1][0], slots2[p2][0]})
		r := min([]int{slots1[p1][1], slots2[p2][1]})
		if l <= r {
			result = append(result, []int{l, r})
		}
		if slots1[p1][1] < slots2[p2][1] {
			p1++
			continue
		}
		p2++
	}

	return result
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}
	return min
}
