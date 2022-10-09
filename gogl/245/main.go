package main

import "fmt"

func main() {
	fmt.Println(shortestWordDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	var min = 100000000000

	if word1 == word2 {
		var pos = []int{}
		for i, word := range wordsDict {
			if word == word1 {
				pos = append(pos, i)
			}
		}
		for i := range pos[:len(pos)-1] {
			if min > pos[i+1]-pos[i] {
				min = pos[i+1] - pos[i]
			}
		}
		return min
	}

	var prevWords = []string{}
	var prevIdx = -1
	for i, word := range wordsDict {
		if word == word1 || word == word2 {
			prevWords = append(prevWords, word)
			prevIdx = i
			break
		}
	}
	for i, word := range wordsDict {
		if i <= prevIdx {
			continue
		}
		if word == word1 || word == word2 {
			if prevWords[len(prevWords)-1] != word {
				if min > i-prevIdx {
					min = i - prevIdx
				}
			}
			prevWords = append(prevWords, word)
			prevIdx = i
		}
	}

	return min
}
