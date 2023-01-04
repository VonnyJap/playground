package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{-3, 1, -8, 4, -1, 2, 1, -5, 5}
	fmt.Println(maxSubArraySum(arr))

}

func maxSubArraySum(arr []int) int {
	result := make([]int, len(arr))
	result[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > result[i-1]+arr[i] {
			result[i] = arr[i]
			continue
		}
		result[i] = result[i-1] + arr[i]
	}

	max := int(math.Inf(-1))
	fmt.Println(result)
	for _, r := range result {
		if max < r {
			max = r
		}
	}

	return max
}
