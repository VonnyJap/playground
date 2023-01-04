package main

import "fmt"

func main() {
	// fmt.Println(lis([]int{1, 2}))
	// fmt.Println(lis([]int{1, 2, 3}))
	// fmt.Println(lis([]int{3, 10, 2, 1, 20}))
	fmt.Println(lis([]int{50, 3, 10, 7, 40, 80}))
}

// func lis(arr []int) int {

// 	var result []int = make([]int, len(arr))

// 	result[0] = 1

// 	for i := 1; i < len(arr); i++ {
// 		for j := 0; j < i; j++ {
// 			if arr[i] > arr[j] {
// 				if result[i] < result[j]+1 {
// 					result[i] = result[j] + 1
// 				}
// 			}
// 		}
// 	}

// 	max := 0
// 	for _, r := range result {
// 		if max < r {
// 			max = r
// 		}
// 	}
// 	return max
// }

var memo = map[int]int{}

func lis(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	tempMax := 0

	for i := 1; i < len(arr); i++ {
		var currentLis int
		if _, ok := memo[i]; ok {
			fmt.Println("memo")
			currentLis = memo[i]
		} else {
			fmt.Println("no memo")
			currentLis = lis(arr[0:i])
		}
		if arr[i-1] < arr[len(arr)-1] {
			currentLis += 1
		}
		if currentLis > tempMax {
			tempMax = currentLis
		}
	}

	memo[len(arr)] = tempMax
	return tempMax
}

// base case is when len arr == 1 => then this is 1
// for anything less than the len(arr) - 1
// 1. check the lis and save it
// 2. if this current number is bigger than the comparison then +1
// 3. compare this with the max and update accordingly

// lets do some testing now
// 1,2
