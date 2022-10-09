package main

import "fmt"

func main() {
	fmt.Println(getMaximumGenerated(7))
	fmt.Println(getMaximumGenerated(2))
	fmt.Println(getMaximumGenerated(3))
}

func getMaximumGenerated(n int) int {

	var arr = make([]int, n+1)
	arr[0] = 0
	if n == 0 {
		return 0
	}
	arr[1] = 1
	if n == 1 {
		return 1
	}

	for i := 1; i <= int(n/2); i++ {
		arr[i*2] = arr[i]
		if i*2+1 > n {
			break
		}
		arr[i*2+1] = arr[i] + arr[i+1]
	}

	fmt.Println(arr)

	max := 0

	for _, n := range arr {
		if max < n {
			max = n
		}
	}

	return max
}
