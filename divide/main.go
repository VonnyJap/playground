package main

import "fmt"

func main() {
	n := 1
	w := -n
	fmt.Println(n)
	fmt.Println(w)

}

// steps:
// check if ctr negative or positive
// setup recursive funtion
// - input
// 	- dividend
// 	- divisor
// 	- ctr
// - base condition
// 	- if dividend == 0
// 		return ctr+1
// 	- dividend < divisor
// 		return ctr

// -7 -- -3 -- 0
// -4 -- -3 -- 1
// -1 -- -3 -- 2
// stop

// 7 -- -3 -- 0
// 4 -- -3 -- -1
// 1 -- -3 -- -2
// 4 -- -3 -- stop
// ctr--
// dividend+divisor < divisor

// -7 -- 3 - 0
// -4 -- 3 - -1
// -1 -- 3 - -2
// ctr--
// dividend+divisor < divisor

// ctr = 0
// for dividend-divisor > divisor
// 	ctr++
// 	dividend = dividend-divisor
// this depends on the operation
