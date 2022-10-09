package main

import (
	"fmt"
	"sort"
)

func main() { fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20})) }

func lemonadeChange(bills []int) bool {
	// use hashmap to store the bills
	var changes []int

	for _, b := range bills {
		fmt.Println(b)
		change := b - 5
		if change != 0 {
			fmt.Println("peekaboo")
			fmt.Println("changes: ", changes)
			// change to while the changes queue not empty?
			var ret []int
			for len(changes) > 0 {
				current := changes[0]
				changes = changes[1:]
				if change-current < 0 {
					ret = append(ret, current)
					continue
				}
				change = change - current
				if change == 0 {
					break
				}
			}
			if change != 0 {
				return false
			}
			changes = append(changes, ret...)
		}
		changes = append(changes, b)
		sort.Slice(changes, func(i, j int) bool {
			return changes[i] > changes[j]
		})

	}

	return true
}
