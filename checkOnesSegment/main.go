package main

import "fmt"

func main() {
	fmt.Println(checkOnesSegment("1001"))
	fmt.Println(checkOnesSegment("110"))
	fmt.Println(checkOnesSegment("1"))
	fmt.Println(checkOnesSegment("1000"))
	fmt.Println(checkOnesSegment("10"))

}

func checkOnesSegment(s string) bool {

	index := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if len(index) != 0 {
				if i-index[len(index)-1] > 1 {
					return false
				}
			}
			index = append(index, i)
		}
	}

	return true
}
