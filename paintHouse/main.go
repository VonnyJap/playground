package main

import "fmt"

func main() {
	fmt.Println(solution([]int{1, 3, 2, 1}))
	fmt.Println(solution([]int{1, 2, 1, 2}))

}

func solution(sequence []int) bool {

	remove := 0

	i := 0
	for i < len(sequence) {
		if i == 0 {
			i++
			continue
		}
		if sequence[i]-sequence[i-1] <= 0 {
			if remove >= 1 {
				return false
			}
			remove++
			sequence = append(sequence[:i], sequence[i+1:]...)
			continue
		}
		i++
	}

	return true
}
