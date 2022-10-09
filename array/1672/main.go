package main

func main() {}

func maximumWealth(accounts [][]int) int {

	max := 0
	for _, i := range accounts {
		res := 0
		for _, w := range i {
			res += w
		}
		if res > max {
			max = res
		}
	}
	return max
}
