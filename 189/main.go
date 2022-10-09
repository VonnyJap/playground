package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// rotateSlow(nums, 1)
	fmt.Println(nums)

	// fmt.Println(reverse([]int{1, 2, 3, 4}))
	rotate(nums, 1)
	fmt.Println(nums)
}

// i am going to do brute force solution first
func rotateSlow(nums []int, k int) {

	// for i := 1; i <= k; i++ {
	// current := nums[0]
	// nums[0] = nums[len(nums)-1]
	// for j := 1; j < len(nums)-1; j++ {
	// current = nums[j+1]
	// nums[j+1] = current
	// nums[2], nums[1] = nums[1], nums[0]
	// fmt.Println(nums)
	k %= len(nums)
	// nums[3], nums[2] = nums[2], nums[1]
	for i := 1; i <= k; i++ {
		current := nums[0]
		for j := 1; j < len(nums); j++ {
			nums[j], current = current, nums[j]
		}
		nums[0] = current
	}
}

// what can we do on the rotating using two pointers

func rotate(nums []int, k int) {

	a := reverse(nums[:len(nums)-k])
	fmt.Println(a)

	b := reverse(nums[len(nums)-k:])
	fmt.Println(b)

	nums = reverse(append(a, b...))

}

func reverse(arr []int) []int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
