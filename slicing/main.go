package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[4:7]
	fmt.Printf("before sorting a, b = %v\n", b) // before sorting a, b = [5 6 7]

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	// b is not [5,6,7] anymore. If we code something to use [5,6,7] in b, then
	// there can be some unpredicted behaviours
	fmt.Printf("after sorting a, b = %v\n", b) // after sorting a, b = [5 4 3]

	// Fix:
	// To avoid that, that part of the slice should be copied to a different slice.
	// Then that values are in different underlying array and changes to a will
	// not be affected to that
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	d := make([]int, 3)
	copy(d, c[4:7])
	fmt.Println(c)
	fmt.Printf("before sorting c, d = %v\n", d) // before sorting c, d = [5 6 7]

	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	fmt.Println(c)
	fmt.Printf("after sorting c, d = %v\n", d) //after sorting c, d = [5 6 7]
}
