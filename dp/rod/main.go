package main

import "fmt"

func main() {
	fmt.Println(rodCut([]int{1, 5, 8, 9, 10, 17, 17, 20}, 4))
}

func rodCut(prices []int, target int) int {

	var arr []int
	arr = append(arr, 0)

	for i := 1; i <= target; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if max < prices[j]+arr[i-j-1] {
				max = prices[j] + arr[i-j-1]
			}
		}
		arr = append(arr, max)
	}

	fmt.Println(arr)

	return arr[target]
}
