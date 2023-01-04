package main

import "fmt"

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))

}

func xorOperation(n int, start int) int {

	result := make([]int, n)

	for i := range result {
		result[i] = start + 2*i
	}

	xor := 0
	for _, num := range result {
		xor = xor ^ num
	}

	return xor
}
