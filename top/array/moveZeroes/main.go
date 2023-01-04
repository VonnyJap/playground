package main

func main() {}

func moveZeroes(nums []int) {

	lastNonZero := 0

	for _, n := range nums {
		if n != 0 {
			nums[lastNonZero] = n
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[lastNonZero] = 0
	}
}
