package main

import "fmt"

func main() {

	// fmt.Println(maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3))
	// fmt.Println(maxSubArrayLen([]int{-2, -1, 2, 1}, 1))

	fmt.Println(maxSubArrayLen([]int{-2, 1, -3, 4, -2}, 1))
}

// using sliding window
// steps:
// 1. get the total sum and if thats equal to k - return immediately len(arr)
// 2. init start at 0 index and loop up to len(nums)/2
// 3. maintain 2 sum => right & left and init at 0
// 4. in each loop update right and left after removing i elements from each left and right
// 5. break immediately if we found k

// [1,-1,5] looking for 5 - total = 5 => return immediately

// [1,-1,5] looking for 4 - total = 5, fromLeft = total, fromRight = total
// i = 0 => fromLeft = fromLeft  - arr[len(arr)-1-i] = 0 => len(nums)-1-i
// => fromRight = fromRight  - arr[i] = 4 => len(nums)-1-i
// => middle = fromLeft - arr[i] - arr[len(arr)-1-i] == 4 => len(nums)-2-2*i
// i = 1 => fromLeft = 0 - 1 = - 1 => len(nums)-1-i = 1
// => fromRight = fromRight  - arr[i] = 5 => len(nums)-1-i = 1
// => middle = fromLeft - arr[i] - arr[len(arr)-1-i] == 4 => len(nums)-2-2*i // if this become 0 or less than skip

func maxSubArrayLen(nums []int, k int) int {

	total := 0
	for _, n := range nums {
		total += n
	}
	length := len(nums)
	fmt.Println("total:", total)
	fmt.Println("length:", length)

	if total == k {
		return length
	}

	leftSum := total
	rightSum := total

	count := 0

	for i := 0; i < length; i++ {

		leftSum -= nums[length-1-i]
		rightSum -= nums[i]
		fmt.Println("leftSum, rightSum:", leftSum, rightSum)
		if rightSum == k || leftSum == k {
			if length-1-i > count {
				count = length - 1 - i
			}
		}
		if length-2-2*i > 0 {
			middle := leftSum + rightSum - total // this one is wrong maybe?
			fmt.Println("middle: ", middle)
			if middle == k {
				if length-2-2*i > count {
					count = length - 2 - 2*i
				}
			}
		}
	}

	return count
}

// func maxSubArrayLen(nums []int, k int) int {

// 	prefix := make([]int, len(nums)+1)
// 	prefix[0] = 0
// 	for i := 1; i <= len(nums); i++ {
// 		prefix[i] = nums[i-1] + prefix[i-1]
// 	}
// 	// fmt.Println(prefix)
// 	count := 0

// 	for i := 1; i < len(prefix); i++ {
// 		for j := 0; j < i; j++ {
// 			if prefix[i]-prefix[j] == k {
// 				temp := i - j
// 				if temp > count {
// 					count = temp
// 				}
// 				break
// 			}
// 		}
// 	}

// 	return count
// }

// Given an integer array nums and an integer k, return the maximum length of a subarray that sums to k. If there is not one, return 0 instead.
// No sorting because we need to have subarray

// what about the brute force?
// we can do two loops - first loop is normal and second loop is j := i+1 and then keeping a sum to keep track and also len of subarray

// what about using the sliding window?
// maybe same thing like the product?

// steps:
// 1. get the prefix sum and store in prefix
// 2. for all j = 0 until j <= index then get the diff and check?
// 3. when we found something we break the loop
