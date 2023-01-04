package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))

}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {

	w1 := -1
	w2 := -1
	distance := int(math.Inf(1))
	for i, word := range wordsDict {
		if word == word1 {
			w1 = i
		} else if word == word2 {
			w2 = i
		}

		if w1 >= 0 && w2 >= 0 {
			temp := w2 - w1
			if temp < 0 {
				temp *= -1
			}
			if temp < distance {
				distance = temp
			}
		}
	}

	return distance
}
