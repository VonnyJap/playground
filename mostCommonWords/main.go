package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	test := "hello"

	for _, r := range test {
		fmt.Println(r - 'a')
	}

	fmt.Println(mostCommonWords("It was the best of times, it was the worst of times."))
}

type WordFrequency struct {
	Word  string
	Count int
}

func mostCommonWords(text string) []WordFrequency {

	// clean up the words from punctuation
	// strings to lower
	// split by space
	// create a map

	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ";", "", -1)
	text = strings.ToLower(text)
	fmt.Println(text)

	words := strings.Split(text, " ")
	wordsMap := map[string]int{}

	for _, word := range words {
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
			continue
		}
		wordsMap[word] = 1
	}
	fmt.Println(wordsMap)

	wordFreq := []WordFrequency{}

	for word, count := range wordsMap {
		wordFreq = append(wordFreq, WordFrequency{word, count})
	}

	sort.Slice(wordFreq, func(i, j int) bool {
		if wordFreq[i].Count == wordFreq[j].Count {
			return wordFreq[i].Word < wordFreq[j].Word
		}
		return wordFreq[i].Count > wordFreq[j].Count
	})

	return wordFreq
}

// type Alphabetic []string

// func (list Alphabetic) Len() int { return len(list) }

// func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

// func (list Alphabetic) Less(i, j int) bool {
//     var si string = list[i]
//     var sj string = list[j]
//     var si_lower = strings.ToLower(si)
//     var sj_lower = strings.ToLower(sj)
//     if si_lower == sj_lower {
//         return si < sj
//     }
//     return si_lower < sj_lower
// }
