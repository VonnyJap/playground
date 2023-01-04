package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}

func plusOne(digits []int) []int {

	index := len(digits) - 1

	for index >= 0 {
		plus := digits[index] + 1
		if plus < 10 {
			digits[index] = plus
			break
		}
		digits[index] = 0
		index--
	}

	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}

	return digits

}
