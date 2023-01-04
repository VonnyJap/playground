package main

import "fmt"

func main() {
	// fmt.Println(totalFruit([]int{1, 2, 1}))
	// fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	// fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	// fmt.Println(totalFruit([]int{1, 2, 3, 2, 1}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))

}

// func totalFruit(fruits []int) int {

// 	dict := map[int]bool{}

// 	for _, f := range fruits {
// 		if _, ok := dict[f]; !ok {
// 			dict[f] = true
// 		}
// 	}
// 	fmt.Println(dict)

// 	// go from the front to the back - we can only reduce from the front anyway
// 	// we keep on making the window size smaller by doing so
// 	i := 0
// 	for i < len(fruits) {
// 		if len(dict) == 2 {
// 			break
// 		}
// 		delete(dict, fruits[i])
// 		i++
// 	}
// 	fmt.Println(i)
// 	return len(fruits) - i
// }

func totalFruit(fruits []int) int {
	dict := map[int]int{}

	start := 0
	max := 0
	temp := 0

	for _, n := range fruits {
		if temp > max {
			max = temp
		}
		temp++
		dict[n]++
		// this loop is only entered when the dict is more than 2
		for len(dict) > 2 {
			dict[fruits[start]]--
			if dict[fruits[start]] == 0 {
				delete(dict, fruits[start])
			}
			temp--  // adjust temp
			start++ // adjust start
		}
	}

	// for _, v := range dict {
	// 	max += v
	// }
	if temp > max {
		max = temp
	}

	return max
}

// this can also be done with max subarray i think in dp

// store the type in dict - populate in each loop
// [1,2,1]
// init a start
// in the loop
// check if element is in dict
// 1. yes - not adding anything -> window++
// 2. no
// 	- add to dict
// 	- check
// 		- if dict > 2 -> removed one equal to the start index from dict, update start - update max when reaching this point
// 		- else window++

// return the window?

// start = 0
// [1,2,1]
// 1 -> window: 1, dict = {1:1}
// 2 -> window: 2, dict = {1:1, 2:1}
// 1 -> window: 3, dict = {1:2, 2:1}

// start = 0
// [0,1,2,2]
// 0 -> window: 1, dict = {0:true}
// 1 -> window: 2, dict = {0:true, 1:true}
// 2 -> window: 2, dict = {1:true, 2:true} - start = 0 + 1 - we have to loop from start until we managed to get down to dict == 2
// 2 -> window: 3, dict = {1:true, 2:true}

// start = 0
// [1,2,3,2,1]
// 1 -> window: 1, dict = {1:true}
// 2 -> window: 2, dict = {1:true, 2:true}
// 3 -> window: 2, dict = {2:true, 3:true} - start = 0 + 1
// 2 -> window: 3, dict = {2:true, 3:true} - start = 0 + 1

// remember do brute force and see how

// single rows of fruit trees arranged from left to right

// brute force
// [1,2,1]
// from 1 -> 2, 1
// from 2 -> 1
// [0,1,2,2]
// from 0 -> 1
// from 1 -> 2
// store this in the map to check
// [1,2,3,2,2]
