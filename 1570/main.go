package main

import "fmt"

func main() {
	v1 := Constructor([]int{1, 0, 0, 2, 3})
	v2 := Constructor([]int{0, 3, 0, 4, 0})
	fmt.Println((&v1).dotProduct(v2))

	v3 := Constructor([]int{0, 1, 0, 0, 0})
	v4 := Constructor([]int{0, 0, 0, 0, 2})
	fmt.Println((&v3).dotProduct(v4))

	v5 := Constructor([]int{0, 1, 0, 0, 2, 0, 0})
	v6 := Constructor([]int{1, 0, 0, 0, 3, 0, 4})
	fmt.Println((&v5).dotProduct(v6))
}

type SparseVector struct {
	Len     int
	NonZero map[int]int
}

func Constructor(nums []int) SparseVector {

	vector := SparseVector{
		Len:     len(nums),
		NonZero: map[int]int{},
	}
	for i, n := range nums {
		if n != 0 {
			vector.NonZero[i] = n
		}
	}

	return vector
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {

	var result = 0
	for pos, val1 := range this.NonZero {
		if val2, ok := vec.NonZero[pos]; ok {
			result += val1 * val2
		}
	}
	return result
}
