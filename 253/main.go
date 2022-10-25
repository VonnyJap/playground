package main

import "fmt"

func main() {

	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))
	fmt.Println(minMeetingRooms([][]int{{9, 10}, {4, 9}, {4, 17}}))
	fmt.Println(minMeetingRooms([][]int{{13, 15}, {1, 13}}))
	fmt.Println(minMeetingRooms([][]int{{3, 18}, {7, 16}, {5, 18}}))
	fmt.Println(minMeetingRooms([][]int{{1, 5}, {8, 9}, {8, 9}}))
}

// intervals = [[0,30],[5,10],[15,20]]
// lets think [[0,30],[5,10],[15,20]]
// brute force solution possible

// 0 - 30
// 5 - 10 - cannot be added in ++
// 15 - 20 - can be added into the previous one

// one way - we can maintain a list of meeting rooms and check if any available

// 0 - 30 -> {1: [[0,30]]}
// 5 - 10 -> // end must be smaller than start of current or has start has to be bigger than current end
// set a var ok - when this is not ok then we set it to false
// // when it is false we need to add more room
func minMeetingRooms(intervals [][]int) int {

	rooms := map[int][][]int{}
	rooms[0] = [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		added := false
		for num, room := range rooms {
			ok := true
			for _, slot := range room {
				// check to see if it will conflict with any slot in this room
				// if yes then we break and continue to the next room
				// how will it conflict then if the start of this room is in between the interval or end of this room is between the interval
				if interval[1] == slot[1] && interval[0] == slot[0] {
					ok = false
					break
				}
				if interval[1] <= slot[0] || interval[0] >= slot[1] {
					continue
				}
				ok = false
				break
			}
			if ok {
				rooms[num] = append(rooms[num], interval)
				added = true
				break
			}
		}
		if !added {
			rooms[len(rooms)+1] = [][]int{interval}
		}
		// fmt.Println(rooms)
	}

	return len(rooms)
}

// minheap solution
// https://medium.com/@lingyong/leetcode-253-meeting-rooms-ii-ee98e40a446f
// minheap all children greater than or same with the current node

// or do sort two pointers start and end to find the current meeting
