package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPowersOfThree(12))
	fmt.Println(checkPowersOfThree(91))
	fmt.Println(checkPowersOfThree(21))

}

func checkPowersOfThree(n int) bool {

	oneChance := true
	for n > 1 {
		if n%3 == 0 {
			n /= 3
			oneChance = true
		} else if oneChance && n%3 == 1 {
			n--
			oneChance = false
		} else {
			return false
		}
	}
	return true
}

func getPowersOfThree(n int) []int {

	result := []int{}
	i := 0
	for {
		result = append(result, int(math.Pow(float64(3), float64(i))))
		if result[len(result)-1] >= n {
			result = result[:len(result)-1]
			break
		}
		i++
	}

	return result
}

// get all power of three in an arr until the number is less than n?
// and then do twosums?
