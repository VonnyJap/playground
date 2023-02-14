package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)

	nums2 := []int{6, 6, 5, 6, 3, 8}
	wiggleSort(nums2)
	fmt.Println(nums2)
}

func wiggleSort(nums []int) {

	sort.Ints(nums)
	start := 0
	end := len(nums) - 1
	result := []int{}

	for start <= end {
		result = append(result, nums[start])
		if start != end {
			result = append(result, nums[end])
		}
		start++
		end--
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}
}

// sort and then swap adjacent

// class Solution {
// 	public:
// 			void swap(int& a, int& b) {
// 					int temp = a;
// 					a = b;
// 					b = temp;
// 			}

// 			void wiggleSort(vector<int>& nums) {
// 					sort(nums.begin(), nums.end());
// 					for (int i = 1; i < nums.size() - 1; i += 2) {
// 							swap(nums[i], nums[i + 1]);
// 					}
// 			}
// 	};
