package main

import (
	"fmt"
	"math"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	ptr := head
	arr := []int{}

	for ptr != nil {
		arr = append(arr, ptr.Val)
		ptr = ptr.Next
	}

	fmt.Println(arr)
	result := 0

	for i, bin := range arr {
		result += bin * int(math.Pow(2, float64(len(arr)-1-i)))
	}

	return result
}
