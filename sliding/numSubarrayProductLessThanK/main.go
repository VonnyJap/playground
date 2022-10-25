package main

import "fmt"

func main() {

	fmt.Println(numSubarrayProductLessThanK([]int{10, 0, 2}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))

}

// func numSubarrayProductLessThanK(nums []int, k int) int {

// 	product := make([]int, len(nums))
// 	product[0] = nums[0]

// 	for i := 1; i < len(nums); i++ {
// 		product[i] = nums[i] * product[i-1]
// 	}

// 	total := 0

// 	// fmt.Println(product)

// 	for i := 0; i < len(nums); i++ {
// 		// fmt.Println("bong", product[i])
// 		if product[i] < k {
// 			// fmt.Println("yap")
// 			total++
// 		}
// 		for j := 0; j < i; j++ {
// 			// fmt.Println("bang", product[i], product[j])
// 			// failed when there is 0
// 			if product[i] == 0 || product[i]/product[j] < k {
// 				// fmt.Println("yeah")
// 				total++
// 			}
// 		}
// 		// fmt.Println(total)
// 	}

// 	return total
// }

func numSubarrayProductLessThanK(nums []int, k int) int {

	if k <= 1 {
		return 0
	}

	total := 0
	product := 1
	left := 0

	for right, num := range nums {
		product *= num
		for product >= k {
			product /= nums[left]
			left++
		}
		total += right - left + 1
	}

	return total
}
