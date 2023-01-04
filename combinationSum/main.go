package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))

}

// assummed to be sorted
func combinationSum(candidates []int, target int) [][]int {

	if target <= 0 {
		return [][]int{}
	}
	final := [][]int{}

	for i, c := range candidates {
		result := combinationSum(candidates[i:], target-c)
		if len(result) == 0 && target-c == 0 {
			final = append(final, []int{c})
			continue
		}
		for _, r := range result {
			all := append([]int{c}, r...)
			if sum(all) == target {
				final = append(final, all)
			}
		}
	}

	return final
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// func combinationSumHelper(candidates []int, target int, temp *[]int) {

// 	if target <= 0 {
// 		return
// 	}
// 	for _, c := range candidates {
// 		*temp = append(*temp, c)
// 		combinationSumHelper(candidates, target-c, temp)
// 	}
// }

// 3,6 - target: 6

// brute force thoughts

// i think this is a recursion
