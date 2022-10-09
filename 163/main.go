package main

import "fmt"

func main() {
	fmt.Println(findMissingRanges([]int{-1}, -2, -1))
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	// what I will do is I will put lower in front
	// and then upper at the back
	// and then if difference is 0 or 1 then no need to do anything
	// if different 2 then just insert number after +1
	// if more than 2 then do something else
	if len(nums) == 0 {
		r := fmt.Sprintf("%d", lower)
		if upper-lower != 0 {
			r = fmt.Sprintf("%s->%d", r, upper)
		}
		return []string{r}
	}

	var result = []string{}

	for i := 0; i < len(nums)-1; i++ {
		var diff = nums[i+1] - nums[i]
		if diff == 0 || diff == 1 {
			continue
		}
		if diff == 2 {
			result = append(result, fmt.Sprintf("%d", nums[i]+1))
			continue
		}
		result = append(result, fmt.Sprintf("%d->%d", nums[i]+1, nums[i+1]-1))
	}

	lb := nums[0] - lower
	switch lb {
	case 0:
		break
	case 1:
		result = append([]string{fmt.Sprintf("%d", lower)}, result...)
		break
	default:
		result = append([]string{fmt.Sprintf("%d->%d", lower, nums[0]-1)}, result...)
	}

	ub := upper - nums[len(nums)-1]
	switch ub {
	case 0:
		break
	case 1:
		result = append(result, fmt.Sprintf("%d", upper))
		break
	default:
		result = append(result, fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper))
	}
	return result
}
