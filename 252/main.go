package main

import "sort"

func main() {

}

// a person can attend the meeting if the merge intervals == intervals
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}
	if len(intervals) == len(merge(intervals)) {
		return true
	}
	return false
}

// what about the brute force way of doing?
// sort by start time and then check one by one if overlap
// if so then return false

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
