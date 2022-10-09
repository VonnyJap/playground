package main

import "fmt"

func main() {

	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	var dict = map[int]int{}
	var stack = []int{}

	for i := 0; i < len(nums2); i++ {
		for len(stack) != 0 {
			if nums2[i] < stack[len(stack)-1] {
				break
			}
			dict[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for _, s := range stack {
		dict[s] = -1
	}

	var result = []int{}

	for _, n := range nums1 {
		if c, ok := dict[n]; ok {
			result = append(result, c)
			continue
		}
		result = append(result, -1)
	}

	return result
}

// public class Solution {
// 	public int[] nextGreaterElement(int[] nums1, int[] nums2) {
// 			Stack<Integer> stack = new Stack<>();
// 			HashMap<Integer, Integer> map = new HashMap<>();

// 			for (int i = 0; i < nums2.length; i++) {
// 					while (!stack.empty() && nums2[i] > stack.peek())
// 							map.put(stack.pop(), nums2[i]);
// 					stack.push(nums2[i]);
// 			}

// 			while (!stack.empty())
// 					map.put(stack.pop(), -1);

// 			int[] res = new int[nums1.length];
// 			for (int i = 0; i < nums1.length; i++) {
// 					res[i] = map.get(nums1[i]);
// 			}

// 			return res;
// 	}
// }
