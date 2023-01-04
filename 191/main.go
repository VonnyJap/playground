package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}

func hammingWeight(num uint32) int {

	one := 0
	arr := []int{}
	for {
		mod := num % 2
		num = num / 2
		arr = append(arr, int(mod))
		if mod == 1 {
			one++
		}
		if num == 0 {
			break
		}
	}

	fmt.Println(arr)
	return one
}
