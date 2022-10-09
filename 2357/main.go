package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{0}))

}

func minimumOperations(nums []int) int {
	// sort and get first non zero and then check if sum is 0 if not then do it again
	// ctr := 0
	// eliminate 0 and then count how many distinct numbers using map
	// nums = mergesort(nums)
	var new = []int{}
	for i := range nums {
		if nums[i] != 0 {
			new = append(new, nums[i])
		}
	}
	var dict = map[int]int{}
	for _, n := range new {
		if _, ok := dict[n]; ok {
			dict[n]++
			continue
		}
		dict[n] = 1
	}
	return len(dict)
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l1 := mergesort(arr[:len(arr)/2])
	l2 := mergesort(arr[len(arr)/2:])
	return merge(l1, l2)
}

func merge(a []int, b []int) []int {
	var c []int

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		c = append(c, a[i])
	}
	for ; j < len(b); j++ {
		c = append(c, b[j])
	}
	return c
}
