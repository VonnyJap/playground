package main

import "fmt"

func main() {

	nums := []int{0, 1, 0, 2, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	ptr1 := 0
	ptr2 := len(nums) - 1

	for ptr1 <= ptr2 {
		if nums[ptr1] == 0 {

			// anything in between rotate by 1
			// before ptr1 no need to do anything
			// after ptr2 no need to do anything as well
			// middle := nums[ptr1:ptr2]
			first := nums[:ptr1]
			middle := append(nums[ptr1+1:ptr2+1], nums[ptr1])
			second := nums[ptr2+1:]

			nums = append(first, append(middle, second...)...)
			ptr2--
		} else {
			ptr1++
		}
	}
}
