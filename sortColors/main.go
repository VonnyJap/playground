package main

import "fmt"

func main() {
	n := []int{2, 0, 1, 3}
	sortColors(n)
	fmt.Println(n)

	w := []int{2, 0, 2, 1, 1, 0}
	sortColors(w)
	fmt.Println(w)

	x := []int{2, 0, 1}
	sortColors(x)
	fmt.Println(x)
}

func sortColors(nums []int) {

	for {
		move := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				move++
			}
		}
		if move == 0 {
			break
		}
	}
}

func sortColors2(nums []int) {

	p0 := 0             // right most 0
	p2 := len(nums) - 1 // left most 2
	curr := 0

	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
		} else {
			curr++
		}
	}

}
