package main

import "fmt"

func main() {

	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}

func numIdenticalPairs(nums []int) int {
	// sort first and then 3 + 2 + 1 + 0
	// loop to find identical

	nums = mergesort(nums)

	result := 0
	same := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			same++
			result += same
			continue
		}
		// add to result
		same = 0
	}

	return result
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l1 := mergesort(arr[:len(arr)/2])
	l2 := mergesort(arr[len(arr)/2:])
	return merge(l1, l2)
}

func merge(a, b []int) []int {
	var c []int

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
			continue
		}
		c = append(c, b[j])
		j++
	}

	for i < len(a) {
		c = append(c, a[i])
		i++
	}
	for j < len(b) {
		c = append(c, b[j])
		j++
	}

	return c
}
