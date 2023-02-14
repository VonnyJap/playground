package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{5, 1, 4, 2}, 6))
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
	fmt.Println(numRescueBoats([]int{2, 2}, 6))
	fmt.Println(numRescueBoats([]int{3, 2, 3, 2, 2}, 6))
}

// steps:
// 1. sort the array
// 2. init remaining to limit - arr[0]
// 3. boats++
// 4. loop from 1 up to end
// 	- if remaining >= arr[i]
// 		- remaining = remaining-arr[i]
// 	- else
// 		boats++
// 		remaining = limit -arr[i]

func numRescueBoats(people []int, limit int) int {

	sort.Ints(people)
	pairs := twoSum(people, limit)

	notpairs := len(people) - pairs*2

	return notpairs + pairs

}

// this will give the indexes

// return how many pairs
func twoSum(numbers []int, target int) int {

	start := 0
	end := len(numbers) - 1

	var pairs = 0

	for start < end {
		// fmt.Println("start: ", start)
		// fmt.Println("end: ", end)
		// fmt.Println("target: ", target)

		if numbers[start]+numbers[end] == target {
			pairs++
			start++
			end--
			continue
		}
		if numbers[start]+numbers[end] > target {
			end--
			continue
		}
		pairs++
		start++
		end--
	}
	return pairs
}

// sort and do two sum
// continue to do two sum until we cannot find twosum or number of people < 2
