package main

import "fmt"

func main() {

	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))

}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	nums = mergesort(nums)
	fmt.Println("sorted nums: ", nums)

	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v != nums[i-1] {
			if i == len(nums)-1 {
				return v
			}
			if v != nums[i+1] {
				return v
			}
		}
	}
	return nums[0]
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

// sort them and loop through
// when you found something not equal to previous, then thats the one

// another method is to use maps or hash table in this case - this is for the search algo
// or we can use the bit manipulation
