package main

import "fmt"

func main() {
	fmt.Println(majorityElementNaive([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElementNaive(nums []int) int {

	// use hash map and then we need to loop through that hashmap also
	var dict = map[int]int{}

	for _, n := range nums {
		if _, ok := dict[n]; ok {
			dict[n]++
			continue
		}
		dict[n] = 1
	}

	for i, v := range dict {
		if v > int(len(nums)/2) {
			return i
		}
	}
	return 0
}

func majorityElementBetter(nums []int) int {
	// sort and then use an integer to track
	target := int(len(nums) / 2)
	nums = mergesort(nums)

	var prev, count int
	for _, n := range nums {
		if n != prev {
			count = 1
		} else {
			count++
		}
		if count > target {
			return n
		}
		prev = n
	}

	return 0
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
