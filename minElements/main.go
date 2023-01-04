package main

import "fmt"

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4))
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0))
	fmt.Println(minElements([]int{2, 2, 2, 5, 1, -2}, 5, 126614243))
}

func minElements(nums []int, limit int, goal int) int {

	// find the diff needs to be filled
	// convert to positive value

	// and then loop with modulo
	total := sum(nums)
	// fmt.Println("total:", total)
	diff := goal - total
	// fmt.Println("diff:", diff)

	if diff == 0 {
		return 0
	}

	if diff < 0 {
		diff *= -1
	}

	// fmt.Println("diff:", diff)

	i := 1

	for {
		if diff <= limit {
			break
		}
		diff = diff - limit
		i++
	}

	return i
}

func sum(nums []int) int {

	result := 0
	for _, n := range nums {
		result += n
	}

	return result
}
