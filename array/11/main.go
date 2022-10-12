package main

import "fmt"

// Container With Most Water
// Input: array of height - total n - height is always be a positive number?
// vertical lines drawn from (x,y)-(i,0) -> (x,y)-(i, height[i])

// brute force
// 1. for each height starting from i=1, go back to the previous starting from 0 up to this i
// 2. whichever is smaller between the two heights will be used as the height of container
// 3. and the width is i - j (where j < i)
// 4. init a max value which is 0 - each loop we will compare with this current max and update accordingly
// 5. return the max

func main() {
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 8, 6}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8}))
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8, 9}))
	fmt.Println(maxAreaOpt([]int{1, 1}))
	fmt.Println(maxAreaOpt([]int{1, 8, 6}))
	fmt.Println(maxAreaOpt([]int{1, 8, 6, 2}))
	fmt.Println(maxAreaOpt([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxAreaOpt([]int{2, 3, 10, 5, 7, 8, 9}))

}

func maxArea(height []int) int {

	var amount = 0

	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			h := min([]int{height[i], height[j]})
			if amount < h*(i-j) {
				amount = h * (i - j)
			}
		}
	}

	return amount
}

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}
	return min
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

// dp way
// 1. to maximize the amount - (max height) x (max width)
// 2. as i++ -> width gets bigger but height might vary and we need to know whats the min_height
// 3. compare it with a local optimum and update the value accordingly
// 4. return the last element as optimum result

func maxAreaOpt(height []int) int {

	var amount = 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		amount = max([]int{amount, min([]int{height[left], height[right]}) * width})
		if height[left] <= height[right] {
			left++
			continue
		}
		right--
	}
	return amount
}
