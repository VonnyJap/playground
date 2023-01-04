package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

// divide by 10 and then keep in stack?
// if last 0 then dont add

func reverse(x int) int {

	stack := []int{}
	negative := false

	if x < 0 {
		x *= -1
		negative = true
	}

	for {
		if x <= 0 {
			break
		}
		stack = append(stack, x%10)
		x = x / 10
	}

	num := 0

	for i, v := range stack {
		if num == 0 && v == 0 {
			continue
		}
		num += v * int(math.Pow(10, float64(len(stack)-1-i)))
	}

	if negative {
		num *= -1
		if num < int(math.Pow(-2, 31)) {
			return 0
		}
	}

	if num > int(math.Pow(2, 31))-1 {
		return 0
	}

	return num

}

// 123
