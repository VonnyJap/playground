package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))

}

func canJump(nums []int) bool {
	temp := map[int]bool{}
	return canJumpUtil(0, nums, &temp)
}

func canJumpUtil(index int, nums []int, temp *map[int]bool) bool {

	if val, ok := (*temp)[index]; ok {
		return val
	}

	if len(nums)-1 == index {
		(*temp)[index] = true
		return true
	}
	if nums[index] == 0 {
		(*temp)[index] = false
		return false
	}
	// can optimize by doing furthest jump
	furthestJump := nums[index]
	if len(nums)-1-index < furthestJump {
		furthestJump = len(nums) - 1 - index
	}
	for i := 1; i <= furthestJump; i++ {
		ok := canJumpUtil(index+i, nums, temp)
		if ok {
			(*temp)[index] = true
			return true
		}
	}
	(*temp)[index] = false
	return false
}
