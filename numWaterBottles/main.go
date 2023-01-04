package main

import "fmt"

func main() {

	fmt.Println(numWaterBottles(9, 3))
	fmt.Println(numWaterBottles(15, 4))

}

func numWaterBottles(numBottles int, numExchange int) int {

	if numBottles < numExchange {
		return numBottles
	}

	newBottles := 0
	curBottles := numBottles
	for curBottles-numExchange >= 0 {
		newBottles++
		curBottles -= numExchange
	}

	return (numBottles - curBottles) + numWaterBottles(newBottles+curBottles, numExchange)
}

// for recursive problem

// 9, 3
// 1. 9
// 2. 9 - 3 -> 1
// 3. 6 - 3 -> 1
// 4. 3 - 3 -> 1
// 5. no more - as long as left - exchange > 0 - continue the loop
// 6. total = 9 + recur(3, 3)
// 7. what will be the base case then?
