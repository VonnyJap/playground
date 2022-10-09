package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
}

var ops = map[string]int{
	"--X": -1,
	"X--": -1,
	"++X": 1,
	"X++": 1,
}

func finalValueAfterOperations(operations []string) int {

	x := 0
	for _, o := range operations {
		if v, ok := ops[o]; ok {
			x = x + v
		}
	}

	return x

}
