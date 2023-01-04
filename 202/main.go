package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))

}

func isHappy(n int) bool {

	if n == 1 {
		return true
	}
	if n < 4 {
		return false
	}
	num := 0
	for {
		num += (n / 10) * (n / 10)
		n -= (n / 10) * 10
		if n < 10 {
			num += (n % 10) * (n % 10)
			break
		}
	}

	return isHappy(num)
}

// as long as number is 10 or more then we need to iter and then recurse isHappy with new number
