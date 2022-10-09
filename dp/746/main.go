package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {

	if len(cost) == 0 {
		return 0
	}

	var result = make([]int, len(cost))
	result[0] = cost[0]
	if len(cost) == 1 {
		return result[0]
	}

	result[1] = cost[1]
	// start from index 2
	for i := 2; i < len(cost); i++ {
		if result[i-1] > result[i-2] {
			result[i] = result[i-2] + cost[i]
			continue
		}
		result[i] = result[i-1] + cost[i]
	}

	if result[len(cost)-1] > result[len(cost)-2] {
		return result[len(cost)-2]
	}
	return result[len(cost)-1]
}

// memoization
// class Solution:
//     def minCostClimbingStairs(self, cost: List[int]) -> int:
//         def minimum_cost(i):
//             # Base case, we are allowed to start at either step 0 or step 1
//             if i <= 1:
//                 return 0

//             # Check if we have already calculated minimum_cost(i)
//             if i in memo:
//                 return memo[i]

//             # If not, cache the result in our hash map and return it
//             down_one = cost[i - 1] + minimum_cost(i - 1)
//             down_two = cost[i - 2] + minimum_cost(i - 2)
//             memo[i] = min(down_one, down_two)
//             return memo[i]

//         memo = {}
//         return minimum_cost(len(cost))
