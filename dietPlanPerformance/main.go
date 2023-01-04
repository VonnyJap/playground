package main

import "fmt"

func main() {
	fmt.Println(dietPlanPerformance([]int{1, 2, 3, 4, 5}, 1, 3, 3))
	fmt.Println(dietPlanPerformance([]int{3, 2}, 2, 0, 1))
	fmt.Println(dietPlanPerformance([]int{6, 5, 0, 0}, 2, 1, 5))

}

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {

	if len(calories) < k {
		return 0
	}

	points := 0

	total := sum(calories[0:k])
	if total < lower {
		points--
	}
	if total > upper {
		points++
	}

	for i := 1; i <= len(calories)-k; i++ {
		total = total - calories[i-1] + calories[i+k-1]
		// fmt.Println("total:", total)
		if total < lower {
			points--
		}
		if total > upper {
			points++
		}
	}

	return points

}

func sum(arr []int) int {

	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}
