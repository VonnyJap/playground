package main

func main() {

}

type MyCalendar struct {
	Schedules [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{Schedules: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {

	for _, schedule := range this.Schedules {
		intersect := findSingleIntersection(schedule, []int{start, end})
		if len(intersect) != 0 {
			return false
		}
	}

	this.Schedules = append(this.Schedules, []int{start, end})
	return true
}

// MyCalendar will hold a list of start+end
// I can use the merge interval to take care of the calendar
// if the merge is hitting the merging than we will issue a false and not changing the list
// if ok then we will do merging

// test merge interval skills now

func findSingleIntersection(firstList []int, secondList []int) []int {
	// when there is no overlap then we will return empty
	if (secondList[0] >= firstList[0] && secondList[0] >= firstList[1]) || (firstList[0] >= secondList[0] && firstList[0] >= secondList[1]) {
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
