package main

import "fmt"

func main() {
	// fmt.Println(judgeCircle("LL"))
	// fmt.Println(judgeCircle("UD"))
	// fmt.Println(judgeCircle("LRLR"))
	fmt.Println(judgeCircle("RLUURDDDLU"))
}

func judgeCircle(moves string) bool {
	fmt.Println(len(moves))
	if len(moves)%2 != 0 {
		return false
	}

	var movesCount = map[string]int{
		"R": 0,
		"L": 0,
		"U": 0,
		"D": 0,
	}

	// divide the move into two and check with the validMoves
	for i := 0; i < len(moves); i++ {
		movesCount[string(moves[i])]++
	}

	return movesCount["R"] == movesCount["L"] && movesCount["U"] == movesCount["D"]
}
