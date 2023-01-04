package main

import "fmt"

func main() {
	// fmt.Println(wordsTyping([]string{"hello", "world"}, 2, 8))
	// fmt.Println(wordsTyping([]string{"a", "bcd", "e"}, 3, 6))
	// fmt.Println(wordsTyping([]string{"i", "had", "apple", "pie"}, 4, 5))
	fmt.Println(wordsTyping([]string{"f", "p", "a"}, 8, 7))
}

func wordsTyping(sentence []string, rows int, cols int) int {

	ctr := 0
	currentRow := 1
	currentCol := 0
	// result := []string{}

	for {
		// need a stop condition
		if currentRow > rows {
			break
		}
		currentCol += len(sentence[ctr%len(sentence)])
		// no need to ctr the sentence because it is not filled in yet
		// how to update the cells? - maybe no need?
		// cells = row * cols
		if currentCol > cols {
			currentRow++
			currentCol = 0
			continue
		}
		// result = append(result, sentence[ctr%len(sentence)])
		ctr++
		if currentCol < cols {
			currentCol++ // for space
		}
	}

	// fmt.Println(result)
	// fmt.Println(ctr)

	return ctr / len(sentence)
}

// calculate how many cells needed for each sentence?
// total rows and cols = rows * cols
// ["hello","world"], total = 2*8, rows = 2, cols = 8
// row = 1, col = 1 => row = 1, col = 1+5(len of hello)-1 = 5
// row = 1, col = 5 => if (8-5) enough for next sentence then good otherwise row++ and col=0
// => row = 2, col = 1+5(len of world)-1 = 5 // if we are at the end of sentence then sentence++ and then set back the ctr to 0
// repeat until done - like when we break the loop?
