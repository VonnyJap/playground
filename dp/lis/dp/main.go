package main

import "fmt"

func main() {
	fmt.Println(lis([]int{7, 5, 3}))
}

func lis(arr []int) int {

	var result []int = make([]int, len(arr))

	result[0] = 1

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if result[i] < result[j]+1 {
					result[i] = result[j] + 1
				}
			}
		}
	}

	max := 0
	for _, r := range result {
		if max < r {
			max = r
		}
	}
	return max
}
