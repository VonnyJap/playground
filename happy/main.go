package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {

	// as long as n/10 greater than 0 then we keep on looping?
	var total int
	for { // this needs to become something to exit as well
		rem := n % 10
		total += rem * rem
		n = int((n - rem) / 10)
		if n >= 10 {
			continue
		}
		total += n * n
		if total == 1 {
			return true
		}
		n = total
		total = 0
	}

	return false

}

// detect a cycle using hashset to catch the false
