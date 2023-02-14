package main

import "fmt"

func main() {
	fmt.Println(canBeIncreasing([]int{1, 2, 10, 5, 7}))
	fmt.Println(canBeIncreasing([]int{2, 3, 1, 2}))
	fmt.Println(canBeIncreasing([]int{1, 1, 1}))

}

func canBeIncreasing(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		arr := []int{}
		for j := 0; j < len(nums); j++ {
			if i != j {
				arr = append(arr, nums[j])
			}
		}
		if isIncreasing(arr) {
			return true
		}
	}

	return false
}

func isIncreasing(nums []int) bool {

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] <= 0 {
			return false
		}
	}
	return true
}

// When we find a drop, we check if the current number nums[i] is greater than the number before the previous one nums[i - 2].

// If so, the number nums[i - 1] needs to be removed.
// Otherwise, the current number needs to be removed (nums[i]).
// For simplicity, I just assign the previous value to the current number (nums[i] = nums[i - 1]).
// And, of course, we return false if we find a second drop.

// bool canBeIncreasing(vector<int>& nums) {
// 	int cnt = 0;
// 	for (int i = 1; i < nums.size() && cnt < 2; ++i) {
// 			if (nums[i - 1] >= nums[i]) {
// 					++cnt;
// 					if (i > 1 && nums[i - 2] >= nums[i])
// 							nums[i] = nums[i - 1];
// 			}
// 	}
// 	return cnt < 2;
// }
